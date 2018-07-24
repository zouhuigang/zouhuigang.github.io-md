---
#标题
title: "linux_python3"
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
tags: [post,python]
#分类
categories: [post,python]
#作者
author: "邹慧刚"
---
### centos安装python3

	cd ~
	下载
	https://www.python.org/ftp/python/3.7.0/Python-3.7.0.tgz


安装:

	tar xfz Python-3.7.0.tgz 
	cd  Python-3.7.0
	./configure   --prefix=/usr/bin/python3.7 
	sudo make
	sudo make install

更改软链接,修改默认Python版本:

	###sudo rm -rf /usr/bin/python(不要删除了)
	sudo mv /usr/bin/python /usr/bin/python2.7.5

	ln -s /usr/bin/python3.7/bin/python3.7  /usr/bin/python

	python -V
	


### 解决yum不能使用的问题

	因为yum包使用python2*等开发，修该为环境修改python3之后有问题，修改文件/usr/bin/yum、/usr/libexec/urlgrabber-ext-down头中相应python为#!/usr/bin/python2.7
解决：


	vi /usr/bin/yum
	#!/usr/bin/python2.7.5
	
	vi /usr/libexec/urlgrabber-ext-down
	#!/usr/bin/python2.7.5


### 升降pip

	 pip install --upgrade pip （升级）

	 python -m pip install --user --upgrade pip==9.0.3  (降级)


### 问题汇总

Q：

	ModuleNotFoundError: No module named '_ctypes'

A:

	yum install libffi-devel -y
