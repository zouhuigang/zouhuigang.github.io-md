---
#标题
title: "README"
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
tags: [post,docker-compose]
#分类
categories: [post,docker-compose]
#作者
author: "邹慧刚"
---
Q1:

使用docker-compose报错:


	pkg_resources.DistributionNotFound: The 'docker-compose==1.7.1' distribution was not found and is required by the application


A1:

	先启动docker mysql

	docker run  -d  -e MYSQL_ROOT_PASSWORD=TYwy2016720 -v /mnt/gitcoding/docker-lnmp-redis/mysql/mysql.cnf:/etc/mysql/conf.d/mysql.cnf:ro -v  /mnt/gitcoding/docker-lnmp-redis/site/mysqldata:/var/lib/mysql -p 3306:3306  registry.aliyuncs.com/zhg_docker_ali_r/mysql 