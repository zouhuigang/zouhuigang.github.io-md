---
#标题
title: "mysql重置密码"
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
tags: [post,mysql-lb]
#分类
categories: [post,mysql-lb]
#作者
author: "邹慧刚"
---
### 查看版本
 mysql > status
	mysql  Ver 8.0.11 for Linux on x86_64 (Source distribution)

### 修改密码（老版本）

	$ mysql
	use mysql;
	update user set Password = password (`123456`) where User = `root` ;
	select host,user from user;
	update user set password=('123456') where user='root';
	
mysql 8修改密码：

	mysql -u root -p '原来的密码'   //进入数据库中
	show databases；
	use mysql；

使用下面的语句修改密码：

      ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '你的密码';  

例如：

        ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456';

推出，使用新的密码登陆
	
	

### 问题汇总

Q1:
	update user set password=('123456') where user='root';
	ERROR 1054 (42S22): Unknown column 'password' in 'field list'
	
A1:

	新版本mysql没有password字段了
	update mysql.user set authentication_string=password('123456') where user='root' ;
	
Q2:

	mysql 配置文件目录：/etc/my.cnf

	root 密码为空的时候配置文件中下面这句：

	skip-grant-tables
	GRANT ALL PRIVILEGES ON *.* TO IDENTIFIED BY '123' WITH GRANT OPTION;

	执行这句时候错误：

	ERROR 1290 (HY000): The MySQL server is running with the --skip-grant-tables option so it cannot execute this statement
	mysql> GRANT ALL PRIVILEGES ON *.* TO IDENTIFIED BY '123' WITH GRANT OPTION;
	ERROR 1290 (HY000): The MySQL server is running with the --skip-grant-tables option so it cannot execute this statement

A2:

	这个时候我们只需要

	flush privileges
	一下，在添加用户就OK了，

	mysql> flush privileges;
	Query OK, 0 rows affected (0.01 sec)

	mysql> GRANT ALL PRIVILEGES ON *.* TO IDENTIFIED BY '123' WITH GRANT OPTION;
	Query OK, 0 rows affected (0.00 sec)

	这个时候我们成功搞定了，再登录就可以了。

	如果报错如下信息：

	Error: Cannot retrieve repository metadata (repomd.xml) for repository: InstallMedia. Please verify its path and try again
	 You could try using --skip-broken to work around the problem
	 You could try running: rpm -Va --nofiles --nodigest

	我们只要到/etc/yum.repo.s下面把packetxxxx.repo和redhat.repo两个文件删除掉，再启动就可以了.