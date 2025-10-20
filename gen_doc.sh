#!/bin/bash
set -e

echo "=== 生成 Swagger API 文档 ==="

# 检查 goctl 是否安装
if ! command -v goctl &> /dev/null; then
    echo "错误: goctl 未安装，请先安装 goctl"
    echo "安装命令: go install github.com/zeromicro/go-zero/tools/goctl@latest"
    exit 1
fi

# 获取参数
SERVICE=${1:-""}

# 如果没有参数，显示使用说明
if [ -z "$SERVICE" ]; then
    echo "使用方法:"
    echo "  ./gen_doc.sh <service>     # 生成指定 Swagger 文档"
    echo "  ./gen_doc.sh all           # 生成所有 Swagger 文档"
    echo ""
    echo "可用的服务:"
    echo "  admin   - 管理端 Swagger 文档"
    echo "  manager - 经理端 Swagger 文档"
    echo "  all     - 所有 Swagger 文档"
    echo ""
    echo "示例:"
    echo "  ./gen_doc.sh admin         # 生成管理端 Swagger 文档"
    echo "  ./gen_doc.sh manager       # 生成经理端 Swagger 文档"
    echo "  ./gen_doc.sh all           # 生成所有 Swagger 文档"
    exit 1
fi

# 创建文档输出目录
mkdir -p docs/swagger

# 生成指定服务的函数
generate_service() {
    local service=$1
    case $service in
        "admin")
            echo "生成管理端 Swagger 文档..."
            goctl api swagger --api api/admin/admin.api --dir docs/swagger --filename admin-swagger --yaml
            echo "✓ 管理端 Swagger 文档生成完成"
            ;;
        "manager")
            echo "生成经理端 Swagger 文档..."
            goctl api swagger --api api/manager/manager.api --dir docs/swagger --filename manager-swagger --yaml
            echo "✓ 经理端 Swagger 文档生成完成"
            ;;
        "all")
            echo "生成所有 Swagger 文档..."
            generate_service "admin"
            generate_service "manager"
            ;;
        *)
            echo "错误: 未知的服务 '$service'"
            echo "可用的服务: admin, manager, all"
            exit 1
            ;;
    esac
}

# 执行生成
generate_service "$SERVICE"

echo "=== Swagger API 文档生成完成 ==="
echo "文档位置:"
if [ "$SERVICE" = "admin" ] || [ "$SERVICE" = "all" ]; then
    echo "  管理端: docs/swagger/admin-swagger.yaml"
fi
if [ "$SERVICE" = "manager" ] || [ "$SERVICE" = "all" ]; then
    echo "  经理端: docs/swagger/manager-swagger.yaml"
fi
echo ""
echo "可以使用 Swagger UI 查看文档:"
echo "  - 在线工具: https://editor.swagger.io/"
echo "  - 本地工具: swagger-ui-serve 或 swagger-codegen"
