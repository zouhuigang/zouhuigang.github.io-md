---
#标题
title: "golang-install"
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
tags: [post,linux]
#分类
categories: [post,linux]
#作者
author: "邹慧刚"
---
### golang官网

	https://golang.org/doc/install
	https://golang.org/dl/
	https://www.golangtc.com/download
	
### 安装

下载:

	wget  https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
	
解压二进制文件到/usr/local目录

	sudo tar zxvf  go1.10.3.linux-amd64.tar.gz   -C  /usr/local
	
设置环境变量,在/etc/profile/后面添加 

	export PATH=$PATH:/usr/bin
	export GOROOT=/usr/local/go
	export GOBIN=$GOROOT/bin
	export GOPATH=$HOME/workspacego
	export PATH=.:$PATH:$GOBIN 
	
使得环境变量生效

	source /etc/profile

	
	
参考文档:

[https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/02.3.md](https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/02.3.md)

[https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/01.1.md](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/01.1.md)