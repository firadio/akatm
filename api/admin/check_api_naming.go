package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type APIFileInfo struct {
	FileName string
	Prefix   string
	Group    string
	Tags     string
	FirstID  string
}

func extractAPIFileInfo(filePath string) (*APIFileInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info := &APIFileInfo{FileName: filepath.Base(filePath)}
	scanner := bufio.NewScanner(file)

	prefixRe := regexp.MustCompile(`prefix:\s*([^\s]+)`)
	groupRe := regexp.MustCompile(`group:\s*(\w+)`)
	tagsRe := regexp.MustCompile(`tags:\s*"([^"]+)"`)
	idRe := regexp.MustCompile(`id:\s*"([^"]+)"`)

	inServerBlock := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, "@server(") {
			inServerBlock = true
			continue
		}

		if inServerBlock {
			if matches := prefixRe.FindStringSubmatch(line); matches != nil {
				info.Prefix = matches[1]
			}
			if matches := groupRe.FindStringSubmatch(line); matches != nil {
				info.Group = matches[1]
			}
			if matches := tagsRe.FindStringSubmatch(line); matches != nil {
				info.Tags = matches[1]
			}
			if strings.Contains(line, ")") {
				inServerBlock = false
			}
		}

		// 提取第一个ID
		if info.FirstID == "" {
			if matches := idRe.FindStringSubmatch(line); matches != nil {
				info.FirstID = matches[1]
			}
		}
	}

	return info, scanner.Err()
}

func checkNamingConsistency(info *APIFileInfo) []string {
	var issues []string

	// 从文件名推断期望的命名
	fileName := strings.TrimSuffix(info.FileName, ".api")

	// 从prefix推断期望的group名称
	// /api/admin/system/user -> adminSystemUser
	// /api/fams/bank/account -> famsBankAccount
	expectedGroup := prefixToGroup(info.Prefix)

	// 检查group命名
	if expectedGroup != "" && info.Group != expectedGroup {
		issues = append(issues, fmt.Sprintf("  ⚠️  group不匹配: 当前=%s, 期望=%s", info.Group, expectedGroup))
	}

	// 检查prefix和文件名的一致性
	expectedPrefix := fileNameToPrefix(fileName)
	if expectedPrefix != "" && !strings.Contains(info.Prefix, expectedPrefix) {
		issues = append(issues, fmt.Sprintf("  ⚠️  prefix可能不匹配文件名: prefix=%s, 文件名=%s", info.Prefix, fileName))
	}

	// 检查ID前缀
	if info.FirstID != "" {
		idPrefix := strings.Split(info.FirstID, ".")[0]
		groupPrefix := getGroupPrefix(info.Group)
		if idPrefix != groupPrefix && idPrefix != "admin" {
			issues = append(issues, fmt.Sprintf("  ⚠️  ID前缀可能不正确: ID=%s, group=%s", info.FirstID, info.Group))
		}
	}

	return issues
}

func prefixToGroup(prefix string) string {
	// /api/admin/system/user -> adminSystemUser
	// /api/fams/bank/account -> famsBankAccount
	parts := strings.Split(strings.Trim(prefix, "/"), "/")
	if len(parts) < 3 {
		return ""
	}

	// 跳过 "api"
	parts = parts[1:]

	// 转换为驼峰命名
	var result strings.Builder
	for i, part := range parts {
		if i == 0 {
			result.WriteString(part) // 第一个部分保持小写
		} else {
			// 首字母大写
			result.WriteString(strings.ToUpper(part[:1]))
			result.WriteString(part[1:])
		}
	}

	return result.String()
}

func fileNameToPrefix(fileName string) string {
	// admin_system_user -> system/user
	// fams_bank_account -> bank/account
	parts := strings.Split(fileName, "_")
	if len(parts) < 2 {
		return ""
	}

	return strings.Join(parts[1:], "/")
}

func getGroupPrefix(group string) string {
	// adminSystemUser -> admin
	// famsBankAccount -> fams
	// iamUser -> iam
	if strings.HasPrefix(group, "admin") {
		return "admin"
	}
	if strings.HasPrefix(group, "fams") {
		return "fams"
	}
	if strings.HasPrefix(group, "iam") {
		return "iam"
	}
	if strings.HasPrefix(group, "public") {
		return "public"
	}
	return ""
}

func main() {
	fmt.Println("========================================")
	fmt.Println("API文件命名一致性检查")
	fmt.Println("========================================")
	fmt.Println()

	docsDir := "docs"
	files, err := filepath.Glob(filepath.Join(docsDir, "*.api"))
	if err != nil {
		fmt.Printf("❌ 读取文件失败: %v\n", err)
		os.Exit(1)
	}

	totalIssues := 0

	for _, file := range files {
		info, err := extractAPIFileInfo(file)
		if err != nil {
			fmt.Printf("⚠️  跳过 %s: %v\n", filepath.Base(file), err)
			continue
		}

		fmt.Printf("📄 %s\n", info.FileName)
		fmt.Printf("   prefix: %s\n", info.Prefix)
		fmt.Printf("   group:  %s\n", info.Group)
		fmt.Printf("   tags:   %s\n", info.Tags)
		if info.FirstID != "" {
			fmt.Printf("   ID示例: %s\n", info.FirstID)
		}

		issues := checkNamingConsistency(info)
		if len(issues) > 0 {
			for _, issue := range issues {
				fmt.Println(issue)
				totalIssues++
			}
		} else {
			fmt.Println("   ✅ 命名一致")
		}
		fmt.Println()
	}

	fmt.Println("========================================")
	if totalIssues == 0 {
		fmt.Println("✅ 所有文件命名一致性检查通过！")
	} else {
		fmt.Printf("⚠️  发现 %d 个潜在问题\n", totalIssues)
	}
}
