package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func fixFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// 正则表达式
	backtick := "`"
	jsonCustomerIdRe := regexp.MustCompile(`UserId\s+(int64\s+` + backtick + `json:"customerId")\s+//\s+客户ID`)
	formCustomerIdRe := regexp.MustCompile(`UserId\s+(int64\s+` + backtick + `form:"customerId,optional")?\s+//\s+客户ID`)
	pathCustomerIdRe := regexp.MustCompile(`UserId\s+(int64\s+` + backtick + `path:"customerId")\s+//\s+客户ID`)

	// 修复缺少反引号的情况
	missingBacktickRe := regexp.MustCompile(`(` + backtick + `(?:json|form|path):"[^` + backtick + `]+")(\s+//)`)

	for scanner.Scan() {
		line := scanner.Text()

		// 修复 json:"customerId" 对应的字段名
		line = jsonCustomerIdRe.ReplaceAllString(line, `CustomerId $1`+backtick+` // 客户ID`)

		// 修复 form:"customerId,optional" 对应的字段名
		line = formCustomerIdRe.ReplaceAllString(line, `CustomerId $1`+backtick+` // 客户ID筛选`)

		// 修复 path:"customerId" 对应的字段名
		line = pathCustomerIdRe.ReplaceAllString(line, `CustomerId $1`+backtick+` // 客户ID`)

		// 修复缺少反引号的情况
		line = missingBacktickRe.ReplaceAllString(line, `$1`+backtick+`$2`)

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// 写回文件
	file.Close()
	output := strings.Join(lines, "\n")
	return os.WriteFile(filePath, []byte(output), 0644)
}

func main() {
	docsDir := "docs"
	files, err := filepath.Glob(filepath.Join(docsDir, "*.api"))
	if err != nil {
		fmt.Printf("❌ 读取文件失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("开始修复 CustomerId 和语法错误...")
	for _, file := range files {
		if err := fixFile(file); err != nil {
			fmt.Printf("⚠️  修复失败 %s: %v\n", filepath.Base(file), err)
		} else {
			fmt.Printf("✅ %s\n", filepath.Base(file))
		}
	}
	fmt.Println("完成！")
}
