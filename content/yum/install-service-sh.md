---
#标题
title: "install-service-sh"
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
tags: [post,yum]
#分类
categories: [post,yum]
#作者
author: "邹慧刚"
---
### 自动生成service脚本

	
	#!/bin/bash
	touch /usr/lib/systemd/system/gdash.service
	echo "[Unit]" > /usr/lib/systemd/system/gdash.service
	echo "Description=gdash: GlusterFS Dashboard">> /usr/lib/systemd/system/gdash.service
	echo "After=glusterd.service" >> /usr/lib/systemd/system/gdash.service
	echo >> /usr/lib/systemd/system/gdash.service
	echo "[Service]" >> /usr/lib/systemd/system/gdash.service
	echo "type=simple" >> /usr/lib/systemd/system/gdash.service
	echo "ExecStart=/usr/bin/gdash -p 80" >> /usr/lib/systemd/system/gdash.service
	echo >> /usr/lib/systemd/system/gdash.service
	echo "[Install]" >> /usr/lib/systemd/system/gdash.service
	echo "WantedBy=multi-user.target" >> /usr/lib/systemd/system/gdash.service
	systemctl start gdash
	systemctl enable gdash

或者:

	#!/usr/bin/bash
	/usr/bin/cp *.service /usr/lib/systemd/system/
	/usr/bin/cp *.timer /usr/lib/systemd/system/



### 使用yum 安装rpm包的时候，不需要用sh脚本来生成service，只需要把/usr/lib/systemd/system/要用到的service打包进rpm里面就行。
