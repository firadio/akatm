package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// APIInfo 存储API文件的信息
type APIInfo struct {
	File   string
	Group  string
	Tag    string
	Prefix string
}

// 从API文件中提取group和tags信息
func extractAPIInfo(filePath string) (*APIInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info := &APIInfo{File: filepath.Base(filePath)}
	scanner := bufio.NewScanner(file)

	// 正则表达式
	groupRe := regexp.MustCompile(`group:\s*(\w+)`)
	tagsRe := regexp.MustCompile(`tags:\s*"([^"]+)"`)
	prefixRe := regexp.MustCompile(`prefix:\s*([^\s]+)`)

	inServerBlock := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 检测@server块
		if strings.Contains(line, "@server(") {
			inServerBlock = true
			continue
		}

		if inServerBlock {
			// 提取group
			if matches := groupRe.FindStringSubmatch(line); matches != nil {
				info.Group = matches[1]
			}

			// 提取tags
			if matches := tagsRe.FindStringSubmatch(line); matches != nil {
				info.Tag = matches[1]
			}

			// 提取prefix
			if matches := prefixRe.FindStringSubmatch(line); matches != nil {
				info.Prefix = matches[1]
			}

			// 检测块结束
			if strings.Contains(line, ")") {
				inServerBlock = false
			}
		}
	}

	if info.Group == "" || info.Tag == "" {
		return nil, fmt.Errorf("未找到group或tags: %s", filePath)
	}

	return info, scanner.Err()
}

// 扫描所有API文件
func scanAPIFiles(dir string) (map[string]string, map[string]string, error) {
	groupToTag := make(map[string]string)
	prefixToTag := make(map[string]string)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理.api文件
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".api") {
			apiInfo, err := extractAPIInfo(path)
			if err != nil {
				fmt.Printf("⚠️  跳过 %s: %v\n", info.Name(), err)
				return nil
			}

			groupToTag[apiInfo.Group] = apiInfo.Tag
			prefixToTag[apiInfo.Prefix] = apiInfo.Tag
			fmt.Printf("✓ %s: %s (%s) -> %s\n", info.Name(), apiInfo.Group, apiInfo.Prefix, apiInfo.Tag)
		}

		return nil
	})

	return groupToTag, prefixToTag, err
}

// 处理swagger.yaml文件
func fixSwaggerTags(swaggerFile string, prefixToTag map[string]string) error {
	fmt.Printf("\n读取 Swagger 文件: %s\n", swaggerFile)

	// 读取文件
	data, err := os.ReadFile(swaggerFile)
	if err != nil {
		return fmt.Errorf("读取文件失败: %v", err)
	}

	// 解析YAML
	var swagger map[string]interface{}
	if err := yaml.Unmarshal(data, &swagger); err != nil {
		return fmt.Errorf("解析YAML失败: %v", err)
	}

	// 收集实际使用的tags
	usedTags := make(map[string]bool)
	addedCount := 0

	// 处理paths中的tags - 根据path前缀匹配tag
	if paths, ok := swagger["paths"].(map[string]interface{}); ok {
		for pathStr, methods := range paths {
			// 查找最长匹配的prefix
			var matchedTag string
			var maxPrefixLen int

			for prefix, tag := range prefixToTag {
				if strings.HasPrefix(pathStr, prefix) && len(prefix) > maxPrefixLen {
					matchedTag = tag
					maxPrefixLen = len(prefix)
				}
			}

			// 如果找到匹配的tag，添加到所有操作中
			if matchedTag != "" {
				if methodMap, ok := methods.(map[string]interface{}); ok {
					for method, operation := range methodMap {
						// 跳过非HTTP方法的键
						if method == "parameters" {
							continue
						}

						if opMap, ok := operation.(map[string]interface{}); ok {
							// 添加tags（goctl不生成tags，所以我们直接添加）
							opMap["tags"] = []interface{}{matchedTag}
							usedTags[matchedTag] = true
							addedCount++
						}
					}
				}
			}
		}
	}

	// 创建tags定义列表
	tagsList := make([]map[string]string, 0)
	sortedTags := make([]string, 0, len(usedTags))
	for tag := range usedTags {
		sortedTags = append(sortedTags, tag)
	}
	sort.Strings(sortedTags)

	for _, tag := range sortedTags {
		tagsList = append(tagsList, map[string]string{
			"name":        tag,
			"description": tag,
		})
	}

	// 添加或更新tags定义
	swagger["tags"] = tagsList

	// 更新info信息
	if info, ok := swagger["info"].(map[string]interface{}); ok {
		info["title"] = "AKATM Admin API"
		info["description"] = "AKATM 后台管理系统 API 文档"
		if _, ok := info["version"]; !ok {
			info["version"] = "v1.0"
		}
	}

	// 备份原文件
	backupFile := swaggerFile + ".backup"
	if err := os.Rename(swaggerFile, backupFile); err == nil {
		fmt.Printf("✓ 已备份原文件到: %s\n", backupFile)
	}

	// 保存修改后的文件
	output, err := yaml.Marshal(&swagger)
	if err != nil {
		return fmt.Errorf("生成YAML失败: %v", err)
	}

	if err := os.WriteFile(swaggerFile, output, 0644); err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	fmt.Printf("\n✅ 完成！\n")
	fmt.Printf("   - 添加了 %d 个接口的tags\n", addedCount)
	fmt.Printf("   - 共 %d 个分类标签\n", len(usedTags))
	fmt.Printf("\n标签列表:\n")

	// 按RPC分组显示
	tagsByPrefix := make(map[string][]string)
	for tag := range usedTags {
		prefix := strings.Split(tag, "-")[0]
		tagsByPrefix[prefix] = append(tagsByPrefix[prefix], tag)
	}

	prefixes := make([]string, 0, len(tagsByPrefix))
	for prefix := range tagsByPrefix {
		prefixes = append(prefixes, prefix)
	}
	sort.Strings(prefixes)

	for _, prefix := range prefixes {
		fmt.Printf("\n【%s】\n", prefix)
		tags := tagsByPrefix[prefix]
		sort.Strings(tags)
		for _, tag := range tags {
			fmt.Printf("  • %s\n", tag)
		}
	}

	return nil
}

func main() {
	fmt.Println("🔧 Swagger Tags 修复工具")
	fmt.Println("========================================")

	// 扫描docs目录下的所有API文件
	fmt.Println("\n📂 扫描 API 文件...")
	groupToTag, prefixToTag, err := scanAPIFiles("docs")
	if err != nil {
		fmt.Printf("❌ 扫描API文件失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n📊 找到 %d 个API定义\n", len(groupToTag))

	// 处理swagger.yaml
	fmt.Println("\n🔄 处理 Swagger 文件...")
	if err := fixSwaggerTags("swagger.yaml", prefixToTag); err != nil {
		fmt.Printf("❌ 处理失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n🎉 现在可以将 swagger.yaml 导入到 Apifox 了！")
}
