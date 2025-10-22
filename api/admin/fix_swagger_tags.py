#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
修复 swagger.yaml 文件，添加 3 级 tags 以支持 Apifox 目录分组
格式：RPC名称/分类/对象
"""

import yaml
import re
from pathlib import Path

# 定义路径前缀和对应的 3 级 tags 映射
# 格式：RPC名称/分类/对象
PATH_TAG_MAPPING = [
    # Admin RPC
    ("/api/admin/auth", "Admin/会话管理/认证"),
    ("/api/admin/system/user", "Admin/系统管理/用户"),
    ("/api/admin/system/role", "Admin/系统管理/角色"),
    ("/api/admin/system/menu", "Admin/系统管理/菜单"),
    ("/api/admin/system/config", "Admin/系统管理/配置"),
    ("/api/admin/system/audit", "Admin/系统管理/审计"),
    ("/api/admin/data/country", "Admin/数据管理/国家"),
    ("/api/admin/dashboard", "Admin/仪表盘/统计"),

    # IAM RPC - 用户管理
    ("/api/iam/user/user-invite", "IAM/用户管理/邀请"),
    ("/api/iam/user/{userId}/profile", "IAM/用户管理/资料"),
    ("/api/iam/user/{userId}/email", "IAM/用户管理/邮箱"),
    ("/api/iam/user/{userId}/credential", "IAM/用户管理/凭证"),
    ("/api/iam/user/{userId}/session", "IAM/用户管理/会话"),
    ("/api/iam/user/{userId}/country", "IAM/用户管理/国家授权"),
    ("/api/iam/user/user", "IAM/用户管理/代理"),

    # FAMS RPC - 用户钱包
    ("/api/fams/user/wallet", "FAMS/用户钱包/钱包"),

    # FAMS RPC - 银行管理
    ("/api/fams/bank/customer", "FAMS/银行管理/客户"),
    ("/api/fams/bank/account-application", "FAMS/银行管理/开户申请"),
    ("/api/fams/bank/account", "FAMS/银行管理/账户"),
    ("/api/fams/bank/deposit", "FAMS/银行管理/存款"),
    ("/api/fams/bank/withdrawal", "FAMS/银行管理/提现"),

    # FAMS RPC - 代理收益
    ("/api/fams/agent/earnings", "FAMS/代理管理/收益"),

    # FAMS RPC - 报表
    ("/api/fams/report", "FAMS/报表管理/报表"),

    # Public API
    ("/api/public/admin/captcha", "Public/公开接口/管理员验证码"),
    ("/api/public/admin/login", "Public/公开接口/管理员登录"),
    ("/api/public/iam/captcha", "Public/公开接口/用户验证码"),
    ("/api/public/iam/login", "Public/公开接口/用户登录"),
    ("/api/public/iam/register", "Public/公开接口/用户注册"),
    ("/api/public/iam/email", "Public/公开接口/邮箱验证"),
]

def get_tag_for_path(path):
    """根据路径获取对应的 tag"""
    # 按最长匹配优先
    for prefix, tag in sorted(PATH_TAG_MAPPING, key=lambda x: len(x[0]), reverse=True):
        if path.startswith(prefix):
            return tag
    return "其他/未分类/未知"

def add_tags_to_swagger(swagger_file):
    """给 swagger.yaml 添加 tags"""
    print(f"读取文件: {swagger_file}")

    with open(swagger_file, 'r', encoding='utf-8') as f:
        swagger = yaml.safe_load(f)

    # 收集所有使用的 tags
    used_tags = set()

    # 给每个 path 添加 tags
    if 'paths' in swagger:
        for path, methods in swagger['paths'].items():
            tag = get_tag_for_path(path)
            used_tags.add(tag)

            for method, details in methods.items():
                if isinstance(details, dict):
                    # 添加 tags 字段
                    details['tags'] = [tag]

                    # 从 summary 或 description 获取接口说明
                    summary = details.get('summary', '')
                    if summary:
                        details['summary'] = summary

    # 在文件头部添加全局 tags 定义
    tags_list = []
    for tag in sorted(used_tags):
        tags_list.append({
            'name': tag,
            'description': tag.replace('/', ' > ')  # 在描述中用 > 表示层级
        })

    swagger['tags'] = tags_list

    # 添加更多元信息
    if 'info' not in swagger:
        swagger['info'] = {}

    swagger['info']['title'] = 'AKATM Admin API'
    swagger['info']['description'] = 'AKATM 后台管理系统 API 文档\n\n目录结构：RPC名称/分类/对象'
    swagger['info']['version'] = 'v1.0'

    # 保存修改后的文件
    backup_file = str(swagger_file) + '.backup'
    print(f"备份原文件到: {backup_file}")

    # 如果已存在备份，先删除
    if Path(backup_file).exists():
        Path(backup_file).unlink()

    Path(swagger_file).rename(backup_file)

    print(f"保存修改后的文件: {swagger_file}")
    with open(swagger_file, 'w', encoding='utf-8') as f:
        yaml.dump(swagger, f, allow_unicode=True, sort_keys=False, default_flow_style=False)

    print(f"\n完成! 共添加 {len(used_tags)} 个标签:")

    # 按层级分组显示
    tags_by_rpc = {}
    for tag in used_tags:
        parts = tag.split('/')
        rpc = parts[0] if len(parts) > 0 else '未知'
        if rpc not in tags_by_rpc:
            tags_by_rpc[rpc] = []
        tags_by_rpc[rpc].append(tag)

    for rpc in sorted(tags_by_rpc.keys()):
        print(f"\n{rpc}:")
        for tag in sorted(tags_by_rpc[rpc]):
            print(f"  {tag}")

if __name__ == '__main__':
    swagger_file = Path(__file__).parent / 'swagger.yaml'
    add_tags_to_swagger(swagger_file)
