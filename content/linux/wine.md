---
#标题
title: "wine"
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
### wine的使用

换国内源(http://mirrors.163.com/.help/ubuntu.html)：

sudo vi /etc/apt/sources.list 在文件最前面添加(trusty(14.04))

	deb http://mirrors.163.com/ubuntu/ trusty main restricted universe multiverse
	deb http://mirrors.163.com/ubuntu/ trusty-security main restricted universe multiverse
	deb http://mirrors.163.com/ubuntu/ trusty-updates main restricted universe multiverse
	deb http://mirrors.163.com/ubuntu/ trusty-proposed main restricted universe multiverse
	deb http://mirrors.163.com/ubuntu/ trusty-backports main restricted universe multiverse
	deb-src http://mirrors.163.com/ubuntu/ trusty main restricted universe multiverse
	deb-src http://mirrors.163.com/ubuntu/ trusty-security main restricted universe multiverse
	deb-src http://mirrors.163.com/ubuntu/ trusty-updates main restricted universe multiverse
	deb-src http://mirrors.163.com/ubuntu/ trusty-proposed main restricted universe multiverse
	deb-src http://mirrors.163.com/ubuntu/ trusty-backports main restricted universe multiverse

然后保存！

	sudo apt-get update

### 安装

	sudo dpkg --add-architecture i386 
	wget -nc https://dl.winehq.org/wine-builds/Release.key
	sudo apt-key add Release.key
	sudo apt-add-repository https://dl.winehq.org/wine-builds/ubuntu/
	sudo apt-add-repository 'deb https://dl.winehq.org/wine-builds/ubuntu/ trusty main'
	sudo apt-get update

##### winehq-staging 

	sudo apt-get install --install-recommends winehq-staging

运行:

	winecfg 检查


	sudo apt install wine-mono
	sudo apt install wine-gecko


重启wine

	wineboot

复制字体：

	sudo cp -r /usr/share/fonts/chinese/** /home/zhg/.wine/drive_c/windows/Fonts/


登录:
	952750120



### 参考文档

	[https://wiki.winehq.org/Ubuntu](https://wiki.winehq.org/Ubuntu)

	[https://www.hiczp.com/post-142.html](https://www.hiczp.com/post-142.html)
	[https://jingyan.baidu.com/article/72ee561a7dc16fe16138df92.html](https://jingyan.baidu.com/article/72ee561a7dc16fe16138df92.html)
