---
title: "让git push命令不再需要密码"
description: "让git push命令不再需要密码-zouhuigang.anooc.com"
author: 邹慧刚
date: 2018-02-24T16:45:55+08:00
categories:
  - "hugo"
tags: ["post", "Use"]
draft: false
---


hugo的一些md参数

	{{.LogicalName}} {{.Path}} {{.Dir}}



ssh提交git，避免每次bat脚本运行都要输入用户名和密码


### 1.生成ssh

	 ssh-keygen -t rsa -C "952750120@qq.com"


![images](../images/20180224150632.png)

一直按回车即可。

可以看出，ssh被保存在C:\Users\mdshi\.ssh目录下。

![images](../images/20180224153047.png)


### 2.把在本地生成ssh公玥复制到github上

id_rsa.pub


把出现的内容包括ssh开头和已邮箱结尾的全部内容复制下来，打开github，登录进入自己的账户，点击自己账号的头像，点击setting，在打开的网页中点击屏幕左边的SSH and GPG keys，在出现的界面下，点击New SSH key,然后把之前赋值的内容复制进key下面的框内即可，title随便你取。


![images](../images/20180224153927.png)



参考文档：

[http://blog.csdn.net/lonyw/article/details/75392410](http://blog.csdn.net/lonyw/article/details/75392410)

