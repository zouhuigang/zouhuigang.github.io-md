---
#标题
title: "logstash-input-jdbc - new"
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
tags: [post,elasticsearch]
#分类
categories: [post,elasticsearch]
#作者
author: "邹慧刚"
---
### 同步mysql中的数据到es中

正确安装java环境


	 $ sudo su -
	 # cd /opt
	 # wget --no-cookies --no-check-certificate --header "Cookie: gpw_e24=http%3A%2F%2Fwww.oracle.com%2F; oraclelicense=accept-securebackup-cookie" "http://download.oracle.com/otn-pub/java/jdk/8u144-b01/090f390dda5b47b9b721c7dfaa008135/jdk-8u144-linux-x64.tar.gz"
	 # tar xzf jdk-8u144-linux-x64.tar.gz
	 # cd /opt/jdk1.8.0_144/
	 # alternatives --install /usr/bin/java java /opt/jdk1.8.0_144/bin/java 2
	 # alternatives --config java
	  --  (choose the new one, #3)
	 # alternatives --install /usr/bin/jar jar /opt/jdk1.8.0_144/bin/jar 2
	 # alternatives --install /usr/bin/javac javac /opt/jdk1.8.0_144/bin/javac 2
	 # alternatives --set jar /opt/jdk1.8.0_144/bin/jar
	 # alternatives --set javac /opt/jdk1.8.0_144/bin/javac
	 # java -version
	  -- it responds with this display
	java version "1.8.0_144"
	Java(TM) SE Runtime Environment (build 1.8.0_144-b01)
	Java HotSpot(TM) 64-Bit Server VM (build 25.144-b01, mixed mode)

Add some more commands to setup up JAVA_HOME and PATH variables:

	 # echo "export JAVA_HOME=/opt/jdk1.8.0_144" >> /etc/profile.d/myprofile.sh
	 # echo "export JRE_HOME=/opt/jdk1.8.0_144/jre" >> /etc/profile.d/myprofile.sh
	 # echo "export PATH=$PATH:/opt/jdk1.8.0_111/bin:/usr/java/jdk1.8.0_144/jre/bin" >> /etc/profile.d/myprofile.sh
	 # chmod 755 /etc/profile.d/myprofile.sh
	 # exit  (to leave the "root" account)
	 $ source /etc/profile.d/myprofile.sh
	 $ echo $JAVA_HOME

QA里面的目录是mkdir /usr/java



创建目录

	mkdir -p /opt/logstash && cd /opt/logstash
	wget https://artifacts.elastic.co/downloads/logstash/logstash-5.2.1.tar.gz
	tar -zxvf logstash-5.2.1.tar.gz

### 安装

	cd /opt/logstash/
	bin/logstash-plugin install logstash-input-jdbc (这句会报错误,cannot install logstash-core-event-java-5.1.1-java)
	bin/logstash-plugin install --local logstash-input-jdbc (安装本地logstash-input-jdbc)



### 安装mysql-connector-java-5.1.43.zip

>官网下载地址 [https://dev.mysql.com/downloads/connector/j/](https://dev.mysql.com/downloads/connector/j/)
>
>下载mysql-connector-java-5.1.33.zip 
>
>curl -o mysql-connector-java-5.1.33.zip -L 'http://dev.mysql.com/get/Downloads/Connector-J/mysql-connector-java-5.1.33.zip/from/http://cdn.mysql.com/' （不能使用mysql-connector-java-5.1.33,会报错）

	cd /opt/logstash 
	wget https://dev.mysql.com/get/Downloads/Connector-J/mysql-connector-java-5.1.43.zip
	unzip mysql-connector-java-5.1.43.zip 

### 使用

>jdbc.config中指定上一部解压的位置,具体看配置文件

	bin/logstash -f jdbc.config
	



参考文档：

[http://blog.csdn.net/fenglailea/article/details/56282414](http://blog.csdn.net/fenglailea/article/details/56282414)

[https://javabirder.wordpress.com/2016/02/21/install-java-8-centos/](https://javabirder.wordpress.com/2016/02/21/install-java-8-centos/)

[https://www.elastic.co/guide/en/logstash/current/offline-plugins.html](https://www.elastic.co/guide/en/logstash/current/offline-plugins.html)

[https://discuss.elastic.co/t/how-to-install-logstash-plugin-in-off-line/68130/2](https://discuss.elastic.co/t/how-to-install-logstash-plugin-in-off-line/68130/2)