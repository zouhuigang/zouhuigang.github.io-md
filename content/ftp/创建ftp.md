---
#标题
title: "创建ftp"
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
tags: [post,ftp]
#分类
categories: [post,ftp]
#作者
author: "邹慧刚"
---
生成一个fpt账户:

	阿里云主机
	用户名：1171
	密码：用户密码


### 目录

	cd /home/ftproot
	mkdir 1171
	lnmp ftp add

生成:

	Enter ftp account name: 1171 (#用户名)
	Enter password for ftp account 1171: (#用户密码)
	Enter directory for ftp account 1171: /home/ftproot/1171 (#ftp根目录)

然后赋予权限组:

	chown www:www -R ./1171


登录截图:

![images](./ftp.jpg)
	

文档:

[https://lnmp.org/faq/ftpserver.html](https://lnmp.org/faq/ftpserver.html)
