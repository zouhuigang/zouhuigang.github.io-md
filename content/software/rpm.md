---
#标题
title: "rpm"
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
tags: [post,software]
#分类
categories: [post,software]
#作者
author: "邹慧刚"
---
### 初始化环境

构建软件环境

	yum install -y go    #编译golang程序使用
	yum install -y git
	go get github.com/hprose/hprose-golang
	go get github.com/zouhuigang/package/zphone/read
	go get github.com/zouhuigang/mahonia
	go get github.com/tealeg/xlsx


打包工具


[ruby安装,参考此链接](http://zouhuigang.anooc.com/ruby/install/)

	gem install fpm
	$ fpm --version
	1.9.3

	  
	  



### 开始构建程序并打包成rpm软件

1.创建软件工作路径

	mkdir -p /usr/local/software/ && cd  /usr/local/software/

2.将source目录下的所有项目上传到/usr/local/software/目录下


3.构建并编译软件，这里用其中一个为例子：

	cd /usr/local/software/zphone
	$ ls
	phone.dat  ZphoneServer.go ZphoneServer.service
	go build ZphoneServer.go
	
在目录下生成ZphoneServer文件。

软件生成完成，删除源代码，防止打包进了rpm包里面了：

	rm -rf ZphoneServer.go && mv ZphoneServer.service /usr/lib/systemd/system/ZphoneServer.service && chmod +x ZphoneServer



4.将ZphoneServer.service上传到/usr/lib/systemd/system/目录下，如需要更改软件路径，则需要修改ZphoneServer.service文件WorkingDirectory和ExecStart路径。



5.打包成rpm，制作软件包

	cd /usr/local/software/zphone 
	fpm -s dir -t rpm -n ZphoneServer -v 1.0.0 --config-files  /usr/lib/systemd/system/ZphoneServer.service  -f /usr/local/software/zphone

之后会在/usr/local/software/zphone目录下，生成ZphoneServer-1.0.0-1.x86_64.rpm

6.将ZphoneServer-1.0.0-1.x86_64.rpm下载到github.com/zouhuigang/software下，更新到github即可。



ZipServer-1.0.0-1.x86_64.rpm :

	cd /usr/local/software/zip && go build ZipServer.go 

	rm -rf ZipServer.go && mv ZipServer.service /usr/lib/systemd/system/ZipServer.service && chmod +x ZipServer

	fpm -s dir -t rpm -n ZipServer -v 1.0.0 --config-files  /usr/lib/systemd/system/ZipServer.service  -f /usr/local/software/zip




zsurname:

	cd /usr/local/software/zsurname && go build ZsurnameServer.go
	rm -rf ZsurnameServer.go && mv ZsurnameServer.service /usr/lib/systemd/system/ZsurnameServer.service && chmod +x ZsurnameServer
	fpm -s dir -t rpm -n ZsurnameServer -v 1.0.0 --config-files  /usr/lib/systemd/system/ZsurnameServer.service  -f /usr/local/software/zsurname


zexcel:
	
	cd /usr/local/software/zexcel && go build ZexcelServer.go
	rm -rf ZexcelServer.go && mv ZexcelServer.service /usr/lib/systemd/system/ZexcelServer.service && chmod +x ZexcelServer
	fpm -s dir -t rpm -n ZexcelServer -v 1.0.0 --config-files  /usr/lib/systemd/system/ZexcelServer.service  -f /usr/local/software/zexcel
	



问题汇总：

Q1：

fpm -s dir -t rpm -n ZphoneServer -v 1.0.0 --config-files  /usr/lib/systemd/system/ZphoneServer.service  -f /usr/local/software/zphone 执行报错:

	Need executable 'rpmbuild' to convert dir to rpm {:level=>:error}

A1:

	出现这种情况就是需要安装rpm-build,安装yum install -y rpm-build.x86_64



	
	
	

	