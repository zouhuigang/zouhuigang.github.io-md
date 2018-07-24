---
#标题
title: "mysql安装"
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
tags: [post,mysql]
#分类
categories: [post,mysql]
#作者
author: "邹慧刚"
---
### 安装

	mkdir -p /home/mysql-install && cd /home/mysql-install
	wget http://repo.mysql.com/mysql-community-release-el7-5.noarch.rpm
	sudo rpm -ivh mysql-community-release-el7-5.noarch.rpm



安装这个包后，会获得两个mysql的yum repo源：

	/etc/yum.repos.d/mysql-community.repo
	/etc/yum.repos.d/mysql-community-source.repo


安装mysql

	sudo yum install mysql-server


修改端口:

vi /etc/my.cnf
	