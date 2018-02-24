---
title: "让git push命令不再需要密码"
author: 邹慧刚
date: 2018-02-24 15:01:27
categories:
  - "git"
tags: ["post", "Push"]
draft: false
---


hugo的一些md参数

	{{.LogicalName}} {{.Path}} {{.Dir}}



ssh提交git，避免每次bat脚本运行都要输入用户名和密码

	两个方式的url地址不同，认证方式也不同。使用ssh时保存密钥对以后可以不再输入帐号密码，而https却不能。所以如果想要不再输入帐号密码，一种方式就是在git clone的时候使用ssh方式，另一种方式就是去修改已有项目.git目录下的config文件中的url，如下：