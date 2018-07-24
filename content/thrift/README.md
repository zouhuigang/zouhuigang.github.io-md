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
tags: [post,thrift]
#分类
categories: [post,thrift]
#作者
author: "邹慧刚"
---
### golang 

1、安装golang的Thrift包：

	go get git.apache.org/thrift.git/lib/go/thrift

	#go install git.apache.org/thrift.git/lib/go/thrift

安装 Thrift 的 IDL 编译工具	

	win:http://www.apache.org/dyn/closer.cgi?path=/thrift/0.10.0/thrift-0.10.0.exe
	linux:https://github.com/apache/thrift/archive/0.10.0.zip

下载完成，加入环境变量，查看是否安装成功

	thrift -version

2.说明

	client目录下的 client.go 实现了客户端用于发送数据并打印接收到 server 端处理后的数据
	server 目录下的 server.go 实现了服务端用于接收客户端发送的数据，并对数据进行大写处理后返回给客户端
	thrift_file 用于存放 thrift 的 IDL 文件： *.thrift


3.编译thrift文件

	进入thrift_file目录,运行:

	thrift -out .. --gen go example.thrift

生成各个版本的库

	thrift -r --gen go batu.thrift
	thrift -r --gen php batu.thrift  
	thrift -r --gen php:server batu.thrift #生成PHP服务端接口代码有所不一样





### 问题:

Q1:

	[root@localhost zexcel]# go run server.go 
	# command-line-arguments
	./server.go:40: cannot use handler (type *RpcServiceImpl) as type rpc.RpcService in argument to rpc.NewRpcServiceProcessor:
	        *RpcServiceImpl does not implement rpc.RpcService (wrong type for FunCall method)
	                have FunCall(int64, string, map[string]string) ([]string, error)
	                want FunCall(context.Context, int64, string, map[string]string) ([]string, error)


A1:

	the reason is your thrift version is different ， maybe you compile the thrift protocol using 0.9.2 version， but include the thrift library file is 1.0.0 version

thrift编译工具的版本:

	[root@localhost zexcel]# thrift -version
	Thrift version 1.0.0-dev
	[root@localhost zexcel]# 

golang引用的代码的版本:

	[root@localhost thrift.git]# git branch -a 
	* master
	  remotes/origin/0.1.x
	  remotes/origin/0.10.0
	  remotes/origin/0.11.0
	  remotes/origin/0.2.x
	  remotes/origin/0.3.x
	  remotes/origin/0.4.x
	  remotes/origin/0.5.x
	  remotes/origin/0.6.x
	  remotes/origin/0.7.x
	  remotes/origin/0.8.x
	  remotes/origin/0.9.1
	  remotes/origin/0.9.2
	  remotes/origin/0.9.3
	  remotes/origin/0.9.x
	  remotes/origin/HEAD -> origin/master
	  remotes/origin/master
	  remotes/origin/py-compiler
	[root@localhost thrift.git]# 

切换到分支:

	git checkout -b 0.10.0 origin/0.10.0



### 参考文档

[http://www.jianshu.com/p/a58665a38022](http://www.jianshu.com/p/a58665a38022)

[https://www.kancloud.cn/digest/batu-go/153528](https://www.kancloud.cn/digest/batu-go/153528)

[https://www.36nu.com/post/248.html](https://www.36nu.com/post/248.html)



