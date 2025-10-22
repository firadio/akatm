#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
为swagger.yaml文件添加tags，使Apifox能够正确分类显示
"""
import re
import yaml

# 定义路径前缀到标签的映射关系
PATH_TAG_MAPPING = {
    # Admin RPC - 认证
    '/api/admin/auth': {'tag': '管理端-认证', 'order': 1},

    # Admin RPC - 系统管理
    '/api/admin/system/user': {'tag': '管理端-系统管理-用户', 'order': 2},
    '/api/admin/system/role': {'tag': '管理端-系统管理-角色', 'order': 3},
    '/api/admin/system/menu': {'tag': '管理端-系统管理-菜单', 'order': 4},
    '/api/admin/system/config': {'tag': '管理端-系统管理-配置', 'order': 5},
    '/api/admin/system/audit': {'tag': '管理端-系统管理-审计日志', 'order': 6},

    # Admin RPC - 数据管理
    '/api/admin/data/country': {'tag': '管理端-数据管理-国家', 'order': 7},

    # Admin RPC - 仪表板
    '/api/admin/dashboard': {'tag': '管理端-仪表板', 'order': 8},

    # IAM RPC - 用户管理
    '/api/iam/user': {'tag': 'IAM-用户管理', 'order': 10},
    '/api/iam/user/profile': {'tag': 'IAM-用户档案', 'order': 11},
    '/api/iam/user/email': {'tag': 'IAM-邮箱管理', 'order': 12},
    '/api/iam/user/credential': {'tag': 'IAM-凭证管理', 'order': 13},
    '/api/iam/user/session': {'tag': 'IAM-会话管理', 'order': 14},
    '/api/iam/user/invite': {'tag': 'IAM-邀请管理', 'order': 15},
    '/api/iam/user/country': {'tag': 'IAM-国家数据', 'order': 16},

    # FAMS RPC - 用户钱包
    '/api/fams/user/wallet': {'tag': 'FAMS-用户钱包', 'order': 20},

    # FAMS RPC - 银行管理
    '/api/fams/bank/customer': {'tag': 'FAMS-银行-客户', 'order': 21},
    '/api/fams/bank/account': {'tag': 'FAMS-银行-账户', 'order': 22},
    '/api/fams/bank/account/application': {'tag': 'FAMS-银行-开户申请', 'order': 23},
    '/api/fams/bank/deposit': {'tag': 'FAMS-银行-存款', 'order': 24},
    '/api/fams/bank/withdrawal': {'tag': 'FAMS-银行-提现', 'order': 25},

    # FAMS RPC - 代理收益
    '/api/fams/agent/earnings': {'tag': 'FAMS-代理收益', 'order': 26},

    # FAMS RPC - 报表
    '/api/fams/user-report': {'tag': 'FAMS-报表-用户报表', 'order': 27},
    '/api/fams/agent-report': {'tag': 'FAMS-报表-代理报表', 'order': 28},

    # 公开接口
    '/api/public': {'tag': '公开接口', 'order': 100},
}

def get_tag_for_path(path):
    """根据路径获取对应的tag"""
    # 按照前缀长度从长到短排序，确保匹配最具体的前缀
    sorted_prefixes = sorted(PATH_TAG_MAPPING.keys(), key=len, reverse=True)

    for prefix in sorted_prefixes:
        if path.startswith(prefix):
            return PATH_TAG_MAPPING[prefix]['tag']

    # 默认返回路径的第二和第三段作为tag
    parts = path.strip('/').split('/')
    if len(parts) >= 3:
        return f"{parts[1]}-{parts[2]}"
    elif len(parts) >= 2:
        return parts[1]
    else:
        return "其他"

def add_tags_to_swagger():
    """为swagger.yaml添加tags"""
    print("正在读取 swagger.yaml...")

    with open('swagger.yaml', 'r', encoding='utf-8') as f:
        content = f.read()

    # 使用yaml库解析
    try:
        swagger_data = yaml.safe_load(content)
    except Exception as e:
        print(f"YAML解析失败: {e}")
        return False

    # 收集所有使用的tags
    used_tags = set()

    # 为每个路径的每个操作添加tags
    if 'paths' in swagger_data:
        for path, methods in swagger_data['paths'].items():
            tag = get_tag_for_path(path)
            used_tags.add(tag)

            for method, operation in methods.items():
                if isinstance(operation, dict):
                    # 添加tags字段
                    if 'tags' not in operation:
                        operation['tags'] = [tag]

    # 在顶层添加tags定义（用于在Apifox中显示分组描述）
    if used_tags:
        tags_list = []
        # 创建tag到order的映射
        tag_order_map = {}
        for prefix, info in PATH_TAG_MAPPING.items():
            tag_order_map[info['tag']] = info['order']

        # 按order排序
        sorted_tags = sorted(used_tags, key=lambda x: tag_order_map.get(x, 999))

        for tag in sorted_tags:
            tags_list.append({
                'name': tag,
                'description': f'{tag}相关接口'
            })

        swagger_data['tags'] = tags_list

    # 写回文件
    print("正在写入更新后的 swagger.yaml...")
    with open('swagger.yaml', 'w', encoding='utf-8') as f:
        yaml.dump(swagger_data, f, allow_unicode=True, sort_keys=False, default_flow_style=False)

    print(f"✅ 完成！已为 {len(used_tags)} 个分类添加tags")
    print(f"分类列表: {', '.join(sorted_tags)}")

    return True

if __name__ == '__main__':
    add_tags_to_swagger()
