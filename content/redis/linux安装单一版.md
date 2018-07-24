---
#标题
title: "linux安装单一版"
#描述
description: ""
#创建日期
date: 2018-07-24
#修改日期
lastmod: 2018-07-24
#草稿
draft: false
#关键字
keywords: []
#标签
tags: [post,redis]
#分类
categories: [post,redis]
#作者
author: "邹慧刚"
---
### 安装

	yum install -y redis


	vi /etc/redis.conf

配置密码：

	requirepass youpwd

查看配置密码:

	 cat /etc/redis.conf |grep requirepass

启动redis:

	systemctl start redis