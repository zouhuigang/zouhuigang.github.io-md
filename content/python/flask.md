---
#标题
title: "flask"
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
### 安装flask

	pip install flask

pip install -r requirements.txt

运行:

	python manage.py runserver
	
	python manage.py -c development  # 开发环境运行
	python manage.py -c testing      # 测试环境运行


问题汇总:

Q:
	安装gevent
	  error: command 'x86_64-linux-gnu-gcc' failed with exit status 1

A:
	(py3) zhg@zhg-ThinkPad-E450c:~/workspaces/py-test$ python -V
	Python 3.6.5


	For Python 2.x use:

	   $ sudo apt-get install python-dev
	For Python 2.7 use:

	   $ sudo apt-get install libffi-dev
	For Python 3.x use:

	   $ sudo apt-get install python3-dev
	For Python 3.4 use:

	   $ sudo apt-get install python3.4-dev
	For Python 3.5 use:

	   $ sudo apt-get install python3.5-dev
	For Python 3.6 use:

	   $ sudo apt-get install python3.6-dev
	
