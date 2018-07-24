---
#标题
title: "java"
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
tags: [post,java]
#分类
categories: [post,java]
#作者
author: "邹慧刚"
---
### 下载


	http://www.oracle.com/technetwork/cn/java/javase/downloads/jdk8-downloads-2133151-zhs.html
	https://www.cnblogs.com/liangml/p/5969007.html


以root用户登录将下载的jdk-8u92-linux-x64.tar.gz文件放到/temp目录下，使用如下命令解压

	wget http://download.oracle.com/otn-pub/java/jdk/8u171-b11/512cd62ec5174c3487ac17c61aaa89e8/jdk-8u171-linux-x64.tar.gz

	tar zxvf jdk-8u171-linux-x64.tar.gz

	mv /home/zhg/下载/jdk1.8.0_171  /opt/


创建软连接

	cd /bin

	ln -s /opt/jdk1.8.0_171/bin/java java


验证:

	zhg@zhg-ThinkPad-E450c:/bin$ java -version
	java version "1.8.0_171"
	Java(TM) SE Runtime Environment (build 1.8.0_171-b11)
	Java HotSpot(TM) 64-Bit Server VM (build 25.171-b11, mixed mode)
	zhg@zhg-ThinkPad-E450c:/bin$ 


设置Java环境(sudo vi /etc/profile):

	

	 #java环境变量
	 export JAVA_HOME=/opt/jdk1.8.0_171
	 export JRE_HOME=${JAVA_HOME}/jre
	 export CLASSPATH=.:${JAVA_HOME}/lib:${JRE_HOME}/lib
	 export PATH={JAVA_HOME}/bin:$PATH


生效:

	source /etc/profile

查看:

	echo $JAVA_HOME

安装tomcat

	https://tomcat.apache.org/download-90.cgi

	sudo tar zvxf apache-tomcat-9.0.10.tar.gz

