---
#标题
title: "practice"
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
### 1.环境

	服务器:47.100.76.132
	端口:8081


### 2.服务端搭建

安装软件:

	mkdir -p /home/yum/centos7/x86_64/
	yum -y install createrepo
	createrepo -pdo /home/yum/centos7/x86_64/ /home/yum/centos7/x86_64/
	cd /home/yum/centos7/x86_64/ 
	python -m SimpleHTTPServer 8081 &>/dev/null &  

[http://blog.csdn.net/why19940926/article/details/62054757](http://blog.csdn.net/why19940926/article/details/62054757)

将rpm软件上传到/home/yum/centos7/x86_64目录下，然后更新：

	createrepo --update /home/yum/centos7/x86_64/


浏览器浏览：

	http://47.100.76.132:8081/


![images](../images/20180224120204.png)


### 3.客户端访问

	cd /etc/yum.repos.d

vi zhg.repo

	[zhg]
	name=Server
	baseurl=http://47.100.76.132:8081
	enable=1
	gpgcheck=0


指定使用zhg库

	yum --enablerepo=zhg --disablerepo=base,extras,updates,epel list



### 4.安装软件

	yum install -y ZphoneServer

	systemctl start ZphoneServer #启动软件
	systemctl status ZphoneServer #查看状态




