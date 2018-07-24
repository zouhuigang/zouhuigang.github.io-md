---
#标题
title: "array"
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
tags: [post,语言对比]
#分类
categories: [post,语言对比]
#作者
author: "邹慧刚"
---
###介绍

 	array 也可写作 []


php:

	array(
	   'k1'=>'v1',
	   'k2'=>'v2'
	);

python:
>在python中array，关联数组被称字典(Dictionary),可设置k,v值。索引数组被称为[]为列表(List),不能自定义k值。

关联array:

	#!/usr/bin/python
 
	dict = {'Name': 'Zara', 'Age': 7, 'Class': 'First'};
	 
	print "dict['Name']: ", dict['Name'];
	print "dict['Age']: ", dict['Age'];

type:

	print(type(arrDict)) 
	<class 'dict'>

索引array:

	#!/usr/bin/python
	 
	list1 = ['physics', 'chemistry', 1997, 2000]
	list2 = [1, 2, 3, 4, 5, 6, 7 ]
	 
	print "list1[0]: ", list1[0]
	print "list2[1:5]: ", list2[1:5]
	
以上实例输出结果：

	list1[0]:  physics
	list2[1:5]:  [2, 3, 4, 5]

	
