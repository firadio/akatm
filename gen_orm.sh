#!/bin/bash
set -e

echo "=== 生成 ORM 代码 ==="

# 获取参数
SERVICE=${1:-""}
ENV=${2:-""}

# 如果没有参数，显示使用说明
if [ -z "$SERVICE" ]; then
    echo "使用方法:"
    echo "  ./gen_orm.sh <service> [env]     # 生成指定 ORM 服务"
    echo "  ./gen_orm.sh all [env]           # 生成所有 ORM 服务"
    echo ""
    echo "可用的服务:"
    echo "  iam     - 身份认证 ORM"
    echo "  admin   - 管理端 ORM"
    echo "  fams    - 金融账户 ORM"
    echo "  all     - 所有服务"
    echo ""
    echo "可用的环境:"
    echo "  (空)    - 默认环境 (.env)"
    echo "  dev     - 开发环境 (.env.dev)"
    echo "  test    - 测试环境 (.env.test)"
    echo "  pro     - 生产环境 (.env.pro)"
    echo ""
    echo "示例:"
    echo "  ./gen_orm.sh iam           # 生成 IAM ORM（默认环境）"
    echo "  ./gen_orm.sh admin test     # 生成 Admin ORM（测试环境）"
    echo "  ./gen_orm.sh all dev        # 生成所有 ORM（开发环境）"
    exit 1
fi

# 设置环境变量
if [ -n "$ENV" ]; then
    export ENV="$ENV"
    echo "使用环境: $ENV"
else
    echo "使用默认环境"
fi

# 生成指定服务的函数
generate_service() {
    local service=$1
    local current_dir=$(pwd)
    case $service in
        "iam")
            echo "生成 IAM ORM..."
            cd rpc/iam/orm/gen
            go run .
            cd "$current_dir"
            echo "✓ IAM ORM 生成完成"
            ;;
        "admin")
            echo "生成 Admin ORM..."
            cd rpc/admin/orm/gen
            go run .
            cd "$current_dir"
            echo "✓ Admin ORM 生成完成"
            ;;
        "fams")
            echo "生成 FAMS ORM..."
            cd rpc/fams/orm/gen
            go run .
            cd "$current_dir"
            echo "✓ FAMS ORM 生成完成"
            ;;
        "all")
            echo "生成所有 ORM 服务..."
            generate_service "iam"
            generate_service "admin"
            generate_service "fams"
            ;;
        *)
            echo "错误: 未知的服务 '$service'"
            echo "可用的服务: iam, admin, fams, all"
            exit 1
            ;;
    esac
}

# 执行生成
generate_service "$SERVICE"

echo "=== ORM 代码生成完成 ==="
