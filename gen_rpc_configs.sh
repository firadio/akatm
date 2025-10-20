#!/bin/bash
set -e

echo "=== 为所有 RPC 服务生成多环境配置 ==="

# RPC 服务列表和端口映射
SERVICES=("iam:8080" "admin:8081" "fams:8082" "mail:8083")

# 为每个服务创建多环境配置
for service_port in "${SERVICES[@]}"; do
    service=$(echo $service_port | cut -d: -f1)
    port=$(echo $service_port | cut -d: -f2)
    echo "处理 $service-rpc 服务 (端口: $port)..."
    
    # 创建 etc 目录（如果不存在）
    mkdir -p "rpc/$service/etc"
    
    # 基础配置（默认环境）
    cat > "rpc/$service/etc/service.yaml" << EOF
Name: $service.rpc
ListenOn: 0.0.0.0:$port
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: $service.rpc
Log:
  ServiceName: $service.rpc
  Mode: console
  Level: info
  Encoding: json
EOF

    # 开发环境配置
    cat > "rpc/$service/etc/service-dev.yaml" << EOF
Name: $service.rpc
ListenOn: 0.0.0.0:$port
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: $service.rpc.dev
Log:
  ServiceName: $service.rpc.dev
  Mode: console
  Level: debug
  Encoding: json
EOF

    # 测试环境配置
    cat > "rpc/$service/etc/service-test.yaml" << EOF
Name: $service.rpc
ListenOn: 0.0.0.0:$port
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: $service.rpc.test
Log:
  ServiceName: $service.rpc.test
  Mode: console
  Level: info
  Encoding: json
EOF

    # 生产环境配置
    cat > "rpc/$service/etc/service-pro.yaml" << EOF
Name: $service.rpc
ListenOn: 0.0.0.0:$port
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: $service.rpc.pro
Log:
  ServiceName: $service.rpc.pro
  Mode: file
  Level: warn
  Encoding: json
  Path: /var/log/akatm/$service.rpc.log
EOF

    echo "✓ $service-rpc 多环境配置生成完成"
done

echo "=== 所有 RPC 服务多环境配置生成完成 ==="
echo ""
echo "服务端口分配："
echo "  iam.rpc   - 端口 8080"
echo "  admin.rpc - 端口 8081"
echo "  fams.rpc  - 端口 8082"
echo "  mail.rpc  - 端口 8083"
echo ""
echo "配置文件说明："
echo "  service.yaml      - 默认环境配置"
echo "  service-dev.yaml  - 开发环境配置"
echo "  service-test.yaml - 测试环境配置"
echo "  service-pro.yaml  - 生产环境配置"
echo ""
echo "使用方法："
echo "  # 使用启动脚本（推荐）"
echo "  ./start_rpc.sh iam dev"
echo "  ./start_rpc.sh admin test"
echo "  ./start_rpc.sh fams pro"
echo ""
echo "  # 手动启动"
echo "  ./rpc/iam/service -f etc/service.yaml"
echo "  ./rpc/iam/service -f etc/service-dev.yaml"
