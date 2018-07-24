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
tags: [post,mariadb]
#分类
categories: [post,mariadb]
#作者
author: "邹慧刚"
---
删除文件之前，先备份，不要直接删除了，尤其是数据库操作。

	systemctl start mariadb 启动不了

报错：

	
	Apr 07 13:23:31 k8s-master1 mysqld[5489]: 2018-04-07 13:23:31 0 [Note] InnoDB: Starting shutdown...
	Apr 07 13:23:31 k8s-master1 mysqld[5489]: 2018-04-07 13:23:31 0 [ERROR] Plugin 'InnoDB' init function returned error.
	Apr 07 13:23:31 k8s-master1 mysqld[5489]: 2018-04-07 13:23:31 0 [ERROR] Plugin 'InnoDB' registration as a STORAGE ENGINE failed.
	Apr 07 13:23:31 k8s-master1 mysqld[5489]: 2018-04-07 13:23:31 0 [Note] Plugin 'FEEDBACK' is disabled.
	Apr 07 13:23:31 k8s-master1 mysqld[5489]: 2018-04-07 13:23:31 0 [ERROR] Unknown/unsupported storage engine: INNODB
	Apr 07 13:23:31 k8s-master1 mysqld[5489]: 2018-04-07 13:23:31 0 [ERROR] Aborting


解决：

修改 vi /etc/my.cnf.d/server.cnf 增加一行tmpdir = /var/tmp 
（这个问题的原因是自己定制的my.cnf中没有设置tmpdir信息，在mysqld段增加tmpdir = /var/tmp 即解决问题。）

	[mysqld]
	port=3310
	tmpdir = /var/tmp 



	http://www.cnblogs.com/zhjh256/p/5763631.html


	http://www.netingcn.com/mac-os-mysql.html

	https://www.chriscalender.com/tag/unknownunsupported-storage-engine-innodb/