#!/bin/bash
set -e

echo "=== 生成 RPC 代码 ==="

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
    echo "  ./gen_rpc.sh <service>     # 生成指定 RPC 服务"
    echo "  ./gen_rpc.sh all           # 生成所有 RPC 服务"
    echo ""
    echo "可用的服务:"
    echo "  iam     - 身份认证服务"
    echo "  admin   - 管理端服务"
    echo "  fams    - 金融账户服务"
    echo "  mail    - 邮件服务"
    echo "  all     - 所有服务"
    exit 1
fi

# 生成指定服务的函数
generate_service() {
    local service=$1
    case $service in
        "iam")
            echo "生成 IAM RPC..."
            goctl rpc protoc rpc/iam/proto/service.proto --go_out=rpc/iam/pb --go-grpc_out=rpc/iam/pb --zrpc_out=rpc/iam --style goZero
            echo "✓ IAM RPC 生成完成"
            ;;
        "admin")
            echo "生成 Admin RPC..."
            goctl rpc protoc rpc/admin/proto/service.proto --go_out=rpc/admin/pb --go-grpc_out=rpc/admin/pb --zrpc_out=rpc/admin --style goZero
            echo "✓ Admin RPC 生成完成"
            ;;
        "fams")
            echo "生成 FAMS RPC..."
            goctl rpc protoc rpc/fams/proto/service.proto --go_out=rpc/fams/pb --go-grpc_out=rpc/fams/pb --zrpc_out=rpc/fams --style goZero
            echo "✓ FAMS RPC 生成完成"
            ;;
        "mail")
            echo "生成 Mail RPC..."
            goctl rpc protoc rpc/mail/proto/service.proto --go_out=rpc/mail/pb --go-grpc_out=rpc/mail/pb --zrpc_out=rpc/mail --style goZero
            echo "✓ Mail RPC 生成完成"
            ;;
        "all")
            echo "生成所有 RPC 服务..."
            generate_service "iam"
            generate_service "admin"
            generate_service "fams"
            generate_service "mail"
            ;;
        *)
            echo "错误: 未知的服务 '$service'"
            echo "可用的服务: iam, admin, fams, mail, all"
            exit 1
            ;;
    esac
}

# 执行生成
generate_service "$SERVICE"

echo "=== RPC 代码生成完成 ==="
