---
#标题
title: "README"
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
tags: [post,mysql]
#分类
categories: [post,mysql]
#作者
author: "邹慧刚"
---
### mysql单独安装，修改端口

	#yum install -y mysql

	yum install mariadb-server mariadb 


日志：

	Updated:
	  mariadb.x86_64 1:5.5.52-1.el7                                                                                                                                                                                   
	
	Dependency Updated:
	  mariadb-libs.x86_64 1:5.5.52-1.el7                                                                                                                                                                              
	
	Complete!


由于MySQL被Oracle收购，所以从Centos7开始不再默认安装mysql而用Mariadb代替，需要使用一些方法才能正确安装Mysql

CentOS 7 版本将MySQL数据库软件从默认的程序列表中移除，用mariadb代替了

	



### 相关命令

	systemctl start mariadb  #启动MariaDB

	systemctl stop mariadb  #停止MariaDB
	
	systemctl restart mariadb  #重启MariaDB
	
	systemctl enable mariadb  #设置开机启动

所以先启动数据库

[root@yl-web yl]# systemctl start mariadb

#### 然后就可以正常使用mysql了



### 方式二 由于用mariadb不习惯，还是安装mysql吧

	yum -y remove mariadb-libs

1. 下载mysql的repo源
	
		wget http://repo.mysql.com/mysql-community-release-el7-5.noarch.rpm

2. 安装mysql-community-release-el7-5.noarch.rpm包

	 	sudo rpm -ivh mysql-community-release-el7-5.noarch.rpm

安装这个包后，会获得两个mysql的yum repo源：/etc/yum.repos.d/mysql-community.repo，/etc/yum.repos.d/mysql-community-source.repo


3.安装mysql

	sudo yum install mysql-server

根据步骤安装就可以了，不过安装完成后，密码为随机密码，需要重置密码。

4.修改端口：

	[mysqld]
	
	port = 3308 #修改为你想要改的端口 vi /etc/my.cnf
	
	datadir=/var/lib/mysql
	
	socket=/var/lib/mysql/mysql.sock
	
	[mysql.server]
	
	user=mysql
	
	basedir=/var/lib
	
	[safe_mysqld]
	
	err-log=/var/log/mysqld.log
	
	pid-file=/var/run/mysqld/mysqld.pid

5. 启动

		systemctl start mysql

5. 重置密码

	重置密码前，首先要登录

	mysql -u root

	mysql > use mysql;
	mysql > update user set password=password('yy2017622') where user='root';
	mysql > flush privileges; 
	mysql > exit;


	#./mysqladmin -u root -p password

	#grep "password" /var/log/mysqld.log 

6.开放端口

	iptables -A INPUT -p tcp -m state --state NEW -m tcp --dport 3308 -j ACCEPT

	service iptables restart


问题：

解决远程连接mysql错误1130

远程连接Mysql服务器的数据库，错误代码是1130，ERROR 1130: Host xxx.xxx.xxx.xxx  is not allowed to connect to this MySQL server  
猜想是无法给远程连接的用户权限问题。 
这样子操作mysql库，即可解决。 
 
在本机登入mysql后，更改 “mysql” 数据库里的 “user” 表里的 “host” 项，从”localhost”改称'%'即可 

	mysql -u root -p  
	use mysql;  
	select 'host' from user where user='root';  
	update user set host = '%' where user ='root';  
	flush privileges;  
	select 'host'   from user where user='root'; 

或者：

	GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'yy2017622' WITH GRANT OPTION; 

	flush privileges; 


ERROR 1062 (23000): Duplicate entry '%-root' for key 'PRIMARY'：

	然后查看了下数据库的host信息如下：
	MySQL> select host from user where user = 'root';
	+-----------------------+
	| host |
	+-----------------------+
	| % |
	| 127.0.0.1 |
	| localhost.localdomain |
	+-----------------------+
	3 rows in set (0.00 sec)
	host已经有了%这个值，所以直接运行命令：
	复制代码 代码如下:
	
	
	MySQL>flush privileges;
	
	
	再用MySQL administrator连接...成功！！




参考：
	
	
	In CentOS 7 - Maria DB is already installed. due to this you are facing the issue.{MariaDB is a compatible implementation of MySQL}. If you want to install Mysql - first you should remove the current existing Mari DB.
	
	Step 1 : Remove Maria DB completely i.e. yum -y remove mariadb-libs
	
	Step 2 : Install MySQL
	
	     i. wget http://dev.mysql.com/get/mysql57-community-release-el7-7.noarch.rpm
	     ii. yum localinstall mysql57-community-release-el7-7.noarch.rpm
	     iii. yum repolist enabled | grep "mysql.*-community.*"
	     iv. yum install mysql-community-server
	     v. yum-config-manager --disable mysql57-community
	     vi. yum-config-manager --enable mysql56-community
	     viii. service mysqld start
	     ix. service mysqld status
	     x. mysql --version
	     xi. mysql -u root -p







[https://www.cnblogs.com/starof/p/4680083.html](https://www.cnblogs.com/starof/p/4680083.html)

[http://www.mamicode.com/info-detail-503994.html](http://www.mamicode.com/info-detail-503994.html)

[https://stackoverflow.com/questions/30696902/fails-installing-mysql-on-centos-7](https://stackoverflow.com/questions/30696902/fails-installing-mysql-on-centos-7)