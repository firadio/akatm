#!/bin/bash
set -e

echo "=== RPC 服务启动脚本 ==="

# 获取参数
SERVICE=${1:-""}
ENV=${2:-"default"}

# 如果没有参数，显示使用说明
if [ -z "$SERVICE" ]; then
    echo "使用方法:"
    echo "  ./start_rpc.sh <service> [env]     # 启动指定 RPC 服务"
    echo ""
    echo "可用的服务:"
    echo "  iam     - 身份认证服务"
    echo "  admin   - 管理端服务"
    echo "  fams    - 金融账户服务"
    echo "  mail    - 邮件服务"
    echo ""
    echo "可用的环境:"
    echo "  default - 默认环境 (默认)"
    echo "  dev     - 开发环境"
    echo "  test    - 测试环境"
    echo "  pro     - 生产环境"
    echo ""
    echo "示例:"
    echo "  ./start_rpc.sh iam dev         # 启动 IAM 服务 (开发环境)"
    echo "  ./start_rpc.sh fams pro        # 启动 FAMS 服务 (生产环境)"
    echo "  ./start_rpc.sh admin           # 启动 Admin 服务 (默认环境)"
    exit 1
fi

# 检查服务目录是否存在
if [ ! -d "rpc/$SERVICE" ]; then
    echo "错误: 服务 '$SERVICE' 不存在"
    echo "可用的服务: iam, admin, fams, mail"
    exit 1
fi

# 检查服务是否已生成
if [ ! -f "rpc/$SERVICE/service" ]; then
    echo "错误: 服务 '$SERVICE' 尚未生成，请先运行:"
    echo "  ./gen_rpc.sh $SERVICE"
    exit 1
fi

# 根据环境选择配置文件
case $ENV in
    "dev")
        CONFIG_FILE="etc/service-dev.yaml"
        echo "启动 $SERVICE-rpc 服务 (开发环境)..."
        ;;
    "test")
        CONFIG_FILE="etc/service-test.yaml"
        echo "启动 $SERVICE-rpc 服务 (测试环境)..."
        ;;
    "pro")
        CONFIG_FILE="etc/service-pro.yaml"
        echo "启动 $SERVICE-rpc 服务 (生产环境)..."
        ;;
    "default"|"")
        CONFIG_FILE="etc/service.yaml"
        echo "启动 $SERVICE-rpc 服务 (默认环境)..."
        ;;
    *)
        echo "错误: 未知的环境 '$ENV'"
        echo "可用的环境: default, dev, test, pro"
        exit 1
        ;;
esac

# 检查配置文件是否存在
if [ ! -f "rpc/$SERVICE/$CONFIG_FILE" ]; then
    echo "错误: 配置文件 'rpc/$SERVICE/$CONFIG_FILE' 不存在"
    echo "请先运行: ./gen_rpc_configs.sh"
    exit 1
fi

# 切换到服务目录并启动服务
cd "rpc/$SERVICE"
echo "配置文件: $CONFIG_FILE"
echo "服务目录: $(pwd)"
echo "启动命令: ./service -f $CONFIG_FILE"
echo ""

# 启动服务
./service -f "$CONFIG_FILE"
