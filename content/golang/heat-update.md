---
title: "golang实现文件的热更新"
description: "golang实现文件的热更新-zouhuigang.anooc.com"
author: 邹慧刚
date: 2018-02-26T15:35:23+08:00
categories:
  - "golang"
tags: ["post", "Heat Update"]
draft: false
---

### golang实现文件的热更新


	https://segmentfault.com/a/1190000008487440
	http://dmdgeeker.com/post/golang-signal/
	http://colobu.com/2015/10/09/Linux-Signals/


### 信号说明

在POSIX.1-1990标准中定义的信号列表

	信号	值	动作	说明
	SIGHUP 	1 	Term 	终端控制进程结束(终端连接断开)
	SIGINT 	2 	Term 	用户发送INTR字符(Ctrl+C)触发
	SIGQUIT 	3 	Core 	用户发送QUIT字符(Ctrl+/)触发
	SIGILL 	4 	Core 	非法指令(程序错误、试图执行数据段、栈溢出等)
	SIGABRT 	6 	Core 	调用abort函数触发
	SIGFPE 	8 	Core 	算术运行错误(浮点运算错误、除数为零等)
	SIGKILL 	9 	Term 	无条件结束程序(不能被捕获、阻塞或忽略)
	SIGSEGV 	11 	Core 	无效内存引用(试图访问不属于自己的内存空间、对只读内存空间进行写操作)
	SIGPIPE 	13 	Term 	消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)
	SIGALRM 	14 	Term 	时钟定时信号
	SIGTERM 	15 	Term 	结束程序(可以被捕获、阻塞或忽略)
	SIGUSR1 	30,10,16 	Term 	用户保留
	SIGUSR2 	31,12,17 	Term 	用户保留
	SIGCHLD 	20,17,18 	Ign 	子进程结束(由父进程接收)
	SIGCONT 	19,18,25 	Cont 	继续执行已经停止的进程(不能被阻塞)
	SIGSTOP 	17,19,23 	Stop 	停止进程(不能被捕获、阻塞或忽略)
	SIGTSTP 	18,20,24 	Stop 	停止进程(可以被捕获、阻塞或忽略)
	SIGTTIN 	21,21,26 	Stop 	后台程序从终端中读取数据时触发
	SIGTTOU 	22,22,27 	Stop 	后台程序向终端中写数据时触发



windows下支持的信号时有限的, 在signal.h中定义, 分别是

             SIGINT      Ctrl+C中断

             SIGILL       非法指令

             SIGFPE      浮点异常

             SIGSEGV   段错误, 非法指针访问

             SIGTERM   kill发出的软件终止

             SIGBREAK Ctrl+Break中断

             SIGABRT   调用abort导致


### 信号用途示例

我们可以利用信号，实现一些动态功能（如不重启刷新加载配置等）。 在使用Nginx时，可以使用nginx -s reload来进行配置文件重载，而不中断服务。 这里我们使用Go语言信号量来模拟一下这个功能。




### 发送信号 

	sudo kill -USR1 24036   # 24036是进程id

	cd D:\mnt\zouhuigang.github.io\content\golang
  

https://github.com/araddon/gou/pull/7
https://github.com/araddon/gou/pull/7/files



### 系统任务

vi /usr/lib/systemd/system/ZsurnameServer.service

	[Unit]
	Description=zouhuigang zsurname server
	Documentation=https://github.com/zouhuigang
	After=network.target
	
	[Service]
	Type=simple
	WorkingDirectory=/nfs_file/software/zsurname
	ExecStart=/nfs_file/software/zsurname/ZsurnameServer
	Restart=on-failure
	PrivateTmp=true
	
	[Install]
	WantedBy=multi-user.target



启动ZsurnameServer：

	systemctl daemon-reload
	systemctl enable ZsurnameServer
	systemctl start ZsurnameServer
	journalctl -u ZsurnameServer.service

更新配置

	sudo kill -USR1 进程id

查看进程

	ps aux | grep ZsurnameServer

显示：

	[root@k8s-master1 system]# ps aux | grep ZsurnameServer
	root      1413  0.0  0.0 187032  3208 ?        Ssl  11:31   0:00 /nfs_file/software/zsurname/ZsurnameServer
	root      1423  0.0  0.0 112660   972 pts/1    S+   11:31   0:00 grep --color=auto ZsurnameServer

可以看到进程id就是1413




日志另存：

![images](../images/20180226113858.jpg)

https://stackoverflow.com/questions/37585758/how-to-redirect-output-of-systemd-service-to-a-file


问题汇总：

Q1:systemd日志出现异常

当使用systemctl status dnsmasq 查看某个服务的状态时候，会提示以下信息

	Warning: Journal has been rotated since unit was started. Log output is incomplete or unavailable.并且服务输出的日志也都不可见。

总结：主要是journalctl满了，才会出现这种情况，可以

	journalctl   --vacuum-size=1G 

来释放占用的磁盘空间，日志文件保存在/var/log/journal/目录下。


默认配置：

	#限制全部日志文件加在一起最多可以占用多少空间，默认值是10%空间与4G空间两者中的较小者
	SystemMaxUse=64G 
	#默认值是15%空间与4G空间两者中的较大者
	SystemKeepFree=1G 
	
	#单个日志文件的大小限制，超过此限制将触发滚动保存
	SystemMaxFileSize=128M 


A1:

这是由于btrfs 挂载的/var/log路径比systemd-journald.service启动时间晚导致。所以我们需要修改/usr/lib/systemd/system/systemd-journald.service ，在After=local-fs.target添加local-fs.target。这样systemd-journald.service就比文件系统挂载晚。该问题就可以得到解决。

	https://bbs.deepin.org/forum.php?mod=viewthread&tid=43821


A1:

	了解现有磁盘使用量

	大家可以利用–disk-usage标记查看journal的当前磁盘使用量：

	[root@k8s-master1 log]# journalctl --disk-usage
	Archived and active journals take up 4.0G on disk.

	删除日志,例如，去年之后的条目才能保留：

	sudo journalctl --vacuum-time=1years
	
或者

	journalctl --vacuum-size=10M
	journalctl   --vacuum-size=1G

日志文件保存在/var/log/journal/目录下。


	http://blog.csdn.net/zstack_org/article/details/56274966
	https://www.linuxquestions.org/questions/linux-newbie-8/systemd-fails-to-write-4175527467/
	http://ju.outofmemory.cn/entry/343962
	https://www.lulinux.com/archives/3135


A2:


### 问题描述

	原先/var/log/journal/的journald日志一直占用4G的存储空间，因为个人电脑不需要保留这么长久的日志信息，通过设置/etc/systemd/journald.conf SystemMaxUse=500M 将日志信息限制到500M
	但是问题来了，通过systemctl status xx.service查看服务运行状态的时候发现所有的服务日志的最后一行全都是Warning: Journal has been rotated since unit was started. Log output is incomplete or unavailable.

### 问题解决

	查阅官方资料发现SystemMaxUse确实可以限制日志容量，但是到达容量上限以后是不会删除原先日志文件
	我是通过SystemMaxFiles=20限制最多存在20个文件，占容量500M~600M，这下才解决问题


	https://github.com/MatcherAny/MatcherAny/issues/3
	https://www.cnblogs.com/hadex/p/6837688.html



