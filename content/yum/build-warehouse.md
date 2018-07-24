---
#标题
title: "build-warehouse"
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
tags: [post,yum]
#分类
categories: [post,yum]
#作者
author: "邹慧刚"
---
# yum仓库搭建




服务端配置
---


### 1.创建 yum 仓库目录，rpm 包都上至此目录

	mkdir -p /home/yum/centos7/x86_64/


### 2. 安装 createrepo 软件

	yum -y install createrepo 


### 3. 初始化 repodata 索引文件

	createrepo -pdo /home/yum/centos7/x86_64/ /home/yum/centos7/x86_64/     


### 4. 进入需要提供仓库的目录

	cd /home/yum/centos7/x86_64/  


### 5. 可以用 Apache 或 Nginx 提供 Web 服务
>但用 Python 的 http 模块更简单，适用于内网环境，可以通过浏览器输入本机 IP 查看。

	python -m SimpleHTTPServer 80 &>/dev/null &  

浏览器打开显示：

![images](../images/20180222153847.png)


### 6. 每加入一个 rpm 包就要更新一下

>将之前制作好的rpm包上传到/home/yum/centos7/x86_64/ 目录下，然后更新索引。
>此时的文件有：
>[root@ceph-admin x86_64]# ls
>repodata  ZphoneServer-1.0.0-1.x86_64.rpm


	createrepo --update /home/yum/centos7/x86_64/

报错：

	Error: Only one directory allowed per run.
	Usage: genpkgmetadata.py [options]

解决：

	检查命令是否正确，--update不要写成-update


### 修改 yum 配置文件 keepacache=0 改为 1，保存下载过的软件

	sed -i 's#keepcache=0#keepcache=1#g' /etc/yum.conf 


之前的文件:

	[root@ceph-admin x86_64]# cat /etc/yum.conf 
	[main]
	cachedir=/var/cache/yum/$basearch/$releasever
	keepcache=0
	debuglevel=2
	logfile=/var/log/yum.log
	exactarch=1
	obsoletes=1
	gpgcheck=1
	plugins=1
	installonly_limit=5
	bugtracker_url=http://bugs.centos.org/set_project.php?project_id=23&ref=http://bugs.centos.org/bug_report_page.php?category=yum
	distroverpkg=centos-release
	
	
	#  This is the default, if you make this bigger yum won't see if the metadata
	# is newer on the remote and so you'll "gain" the bandwidth of not having to
	# download the new metadata and "pay" for it by yum not having correct
	# information.
	#  It is esp. important, to have correct metadata, for distributions like
	# Fedora which don't keep old packages around. If you don't like this checking
	# interupting your command line usage, it's much better to have something
	# manually check the metadata once an hour (yum-updatesd will do this).
	# metadata_expire=90m
	
	# PUT YOUR REPOS HERE OR IN separate files named file.repo
	# in /etc/yum.repos.d 


修改之后的文件:

	
	[root@ceph-admin x86_64]# cat /etc/yum.conf 
	[main]
	cachedir=/var/cache/yum/$basearch/$releasever
	keepcache=1
	debuglevel=2
	logfile=/var/log/yum.log
	exactarch=1
	obsoletes=1
	gpgcheck=1
	plugins=1
	installonly_limit=5
	bugtracker_url=http://bugs.centos.org/set_project.php?project_id=23&ref=http://bugs.centos.org/bug_report_page.php?category=yum
	distroverpkg=centos-release
	
	
	#  This is the default, if you make this bigger yum won't see if the metadata
	# is newer on the remote and so you'll "gain" the bandwidth of not having to
	# download the new metadata and "pay" for it by yum not having correct
	# information.
	#  It is esp. important, to have correct metadata, for distributions like
	# Fedora which don't keep old packages around. If you don't like this checking
	# interupting your command line usage, it's much better to have something
	# manually check the metadata once an hour (yum-updatesd will do this).
	# metadata_expire=90m
	
	# PUT YOUR REPOS HERE OR IN separate files named file.repo
	# in /etc/yum.repos.d



客户端配置
---

在另外一台服务器上，例如：192.168.122.148

### 自定义zhg.repo文件

	cd /etc/yum.repos.d

vi zhg.repo

	[zhg] #指定使用zhg库
	name=Server
	baseurl=http://192.168.122.151 #yum仓库ip
	enable=1 #开启什么模块
	gpgcheck=0 #是否检查gpg



指定使用zhg库


	yum --enablerepo=zhg --disablerepo=base,extras,updates,epel list


安装库中的软件

	yum -y install ZphoneServer   # 安装软件
	systemctl start ZphoneServer  # 启动软件
	systemctl status ZphoneServer # 查看状态
	yum remove ZphoneServer #卸载软件




### 问题汇总：

Q1:

	http://192.168.122.151/repodata/repomd.xml: [Errno 12] Timeout on http://192.168.122.151/repodata/repomd.xml: (28, 'Operation too slow. Less than 1000 bytes/sec transferred the last 30 seconds')


A1:

	检查repodata目录下是否有repomd.xml文件，这个文件是自动生成的。



### 参考文档

[http://jcenter.idcos.com/?/article/28](http://jcenter.idcos.com/?/article/28)
	
	




