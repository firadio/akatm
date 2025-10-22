package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// PathTagMapping 定义路径前缀和对应的 3 级 tags 映射
// 格式：RPC名称/分类/对象
var PathTagMapping = []struct {
	Prefix string
	Tag    string
}{
	// Admin RPC
	{"/api/admin/auth", "Admin/会话管理/认证"},
	{"/api/admin/system/user", "Admin/系统管理/用户"},
	{"/api/admin/system/role", "Admin/系统管理/角色"},
	{"/api/admin/system/menu", "Admin/系统管理/菜单"},
	{"/api/admin/system/config", "Admin/系统管理/配置"},
	{"/api/admin/system/audit", "Admin/系统管理/审计"},
	{"/api/admin/data/country", "Admin/数据管理/国家"},
	{"/api/admin/dashboard", "Admin/仪表盘/统计"},

	// IAM RPC - 用户管理
	{"/api/iam/user/user-invite", "IAM/用户管理/邀请"},
	{"/api/iam/user/{userId}/profile", "IAM/用户管理/资料"},
	{"/api/iam/user/{userId}/email", "IAM/用户管理/邮箱"},
	{"/api/iam/user/{userId}/credential", "IAM/用户管理/凭证"},
	{"/api/iam/user/{userId}/session", "IAM/用户管理/会话"},
	{"/api/iam/user/{userId}/country", "IAM/用户管理/国家授权"},
	{"/api/iam/user/user", "IAM/用户管理/代理"},

	// FAMS RPC - 用户钱包
	{"/api/fams/user/wallet", "FAMS/用户钱包/钱包"},

	// FAMS RPC - 银行管理
	{"/api/fams/bank/customer", "FAMS/银行管理/客户"},
	{"/api/fams/bank/account-application", "FAMS/银行管理/开户申请"},
	{"/api/fams/bank/account", "FAMS/银行管理/账户"},
	{"/api/fams/bank/deposit", "FAMS/银行管理/存款"},
	{"/api/fams/bank/withdrawal", "FAMS/银行管理/提现"},

	// FAMS RPC - 代理收益
	{"/api/fams/agent/earnings", "FAMS/代理管理/收益"},

	// FAMS RPC - 报表
	{"/api/fams/report", "FAMS/报表管理/报表"},

	// Public API
	{"/api/public/admin/captcha", "Public/公开接口/管理员验证码"},
	{"/api/public/admin/login", "Public/公开接口/管理员登录"},
	{"/api/public/iam/captcha", "Public/公开接口/用户验证码"},
	{"/api/public/iam/login", "Public/公开接口/用户登录"},
	{"/api/public/iam/register", "Public/公开接口/用户注册"},
	{"/api/public/iam/email", "Public/公开接口/邮箱验证"},
}

func getTagForPath(path string) string {
	// 按最长匹配优先
	maxLen := 0
	result := "其他/未分类/未知"

	for _, mapping := range PathTagMapping {
		if strings.HasPrefix(path, mapping.Prefix) && len(mapping.Prefix) > maxLen {
			maxLen = len(mapping.Prefix)
			result = mapping.Tag
		}
	}

	return result
}

func main() {
	swaggerFile := "swagger.yaml"

	fmt.Printf("读取文件: %s\n", swaggerFile)

	// 读取 YAML 文件
	data, err := ioutil.ReadFile(swaggerFile)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		os.Exit(1)
	}

	// 解析 YAML
	var swagger map[string]interface{}
	err = yaml.Unmarshal(data, &swagger)
	if err != nil {
		fmt.Printf("解析 YAML 失败: %v\n", err)
		os.Exit(1)
	}

	// 收集所有使用的 tags
	usedTags := make(map[string]bool)

	// 给每个 path 添加 tags
	if paths, ok := swagger["paths"].(map[string]interface{}); ok {
		for path, methods := range paths {
			tag := getTagForPath(path)
			usedTags[tag] = true

			if methodMap, ok := methods.(map[string]interface{}); ok {
				for _, details := range methodMap {
					if detailMap, ok := details.(map[string]interface{}); ok {
						// 添加 tags 字段
						detailMap["tags"] = []string{tag}
					}
				}
			}
		}
	}

	// 在文件头部添加全局 tags 定义
	tagsList := make([]map[string]string, 0)
	sortedTags := make([]string, 0, len(usedTags))
	for tag := range usedTags {
		sortedTags = append(sortedTags, tag)
	}
	sort.Strings(sortedTags)

	for _, tag := range sortedTags {
		tagsList = append(tagsList, map[string]string{
			"name":        tag,
			"description": strings.ReplaceAll(tag, "/", " > "),
		})
	}

	swagger["tags"] = tagsList

	// 添加更多元信息
	if _, ok := swagger["info"]; !ok {
		swagger["info"] = make(map[string]interface{})
	}

	info := swagger["info"].(map[string]interface{})
	info["title"] = "AKATM Admin API"
	info["description"] = "AKATM 后台管理系统 API 文档\n\n目录结构：RPC名称/分类/对象"
	info["version"] = "v1.0"

	// 备份原文件
	backupFile := swaggerFile + ".backup"
	fmt.Printf("备份原文件到: %s\n", backupFile)

	// 如果已存在备份，先删除
	os.Remove(backupFile)

	err = os.Rename(swaggerFile, backupFile)
	if err != nil {
		fmt.Printf("备份文件失败: %v\n", err)
		os.Exit(1)
	}

	// 保存修改后的文件
	fmt.Printf("保存修改后的文件: %s\n", swaggerFile)

	output, err := yaml.Marshal(&swagger)
	if err != nil {
		fmt.Printf("生成 YAML 失败: %v\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(swaggerFile, output, 0644)
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n完成! 共添加 %d 个标签:\n", len(usedTags))

	// 按层级分组显示
	tagsByRPC := make(map[string][]string)
	for tag := range usedTags {
		parts := strings.Split(tag, "/")
		rpc := "未知"
		if len(parts) > 0 {
			rpc = parts[0]
		}
		tagsByRPC[rpc] = append(tagsByRPC[rpc], tag)
	}

	rpcs := make([]string, 0, len(tagsByRPC))
	for rpc := range tagsByRPC {
		rpcs = append(rpcs, rpc)
	}
	sort.Strings(rpcs)

	for _, rpc := range rpcs {
		fmt.Printf("\n%s:\n", rpc)
		tags := tagsByRPC[rpc]
		sort.Strings(tags)
		for _, tag := range tags {
			fmt.Printf("  %s\n", tag)
		}
	}
}
