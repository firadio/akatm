#!/bin/bash
set -e

echo "=== 生成 API 网关代码 ==="

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
    echo "  ./gen_api.sh <service>     # 生成指定 API 网关"
    echo "  ./gen_api.sh all           # 生成所有 API 网关"
    echo ""
    echo "可用的服务:"
    echo "  admin   - 管理端网关"
    echo "  manager - 经理端网关"
    echo "  all     - 所有网关"
    echo ""
    echo "示例:"
    echo "  ./gen_api.sh admin         # 生成管理端网关"
    echo "  ./gen_api.sh manager       # 生成经理端网关"
    echo "  ./gen_api.sh all           # 生成所有网关"
    exit 1
fi

# 生成指定服务的函数
generate_service() {
    local service=$1
    case $service in
        "admin")
            echo "生成管理端网关..."
            goctl api go -api api/admin/admin.api -dir api/adminGateway -style goZero
            echo "✓ 管理端网关生成完成"
            ;;
        "manager")
            echo "生成经理端网关..."
            goctl api go -api api/manager/manager.api -dir api/managerGateway -style goZero
            echo "✓ 经理端网关生成完成"
            ;;
        "all")
            echo "生成所有 API 网关..."
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

echo "=== API 网关生成完成 ==="
