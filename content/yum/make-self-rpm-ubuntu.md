---
#标题
title: "make-self-rpm-ubuntu"
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
### 安装ruby

	sudo apt install ruby
	sudo apt-get intall ruby-dev
	sudo apt-get install rpm (#含有rpmbuild包)
	sudo gem install fpm
	


### 打包fpm

	cd /usr/local/software/zrule  && chmod +x ZruleServer

	sudo fpm -s dir -t rpm -n ZruleServer -v 1.0.0 --config-files  /usr/lib/systemd/system/ZruleServer.service  -f /usr/local/software/zrule

### 问题汇总

Q1:

	ERROR: Failed to build gem native extension.

A1:

	在Ubuntu 默认的apt-get安装ruby时，没有安装ruby-dev，需在terminal中输入sudo apt-get intall ruby-dev 即可。亲测，问题解决。

	[https://blog.csdn.net/csvdvg/article/details/62444144](https://blog.csdn.net/csvdvg/article/details/62444144)


Q2:

	Need executable 'rpmbuild' to convert dir to rpm {:level=>:error}

A2:
	制作 rpm 包需要用到 rpmbuild 工具。在 ubuntu 上，该工具包含在 rpm 包中，可以直接从源里安装：
	sudo apt-get install rpm
	