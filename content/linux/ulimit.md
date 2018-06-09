＃＃＃　更改文件打开数量

可以看到默认打开的文件数限制为1024个:

	$ ulimit -n  
	1024 
	
编辑/etc/profile配置文件，在最后添加一行

	ulimit -SHn 65535  
	
	
### 问题汇总

Q1:	bash: ulimit: open files: 无法修改 limit 值: 不允许的操作

A1：

	为啥root用户是可以的？普通用户又会遇到这样的问题？

	看一下/etc/security/limits.conf大概就会明白。
	linux对用户有默认的ulimit限制，而这个文件可以配置用户的硬配置和软配置，硬配置是个上限。超出上限的修改就会出“不允许的操作”这样的错误。在limits.conf加上

	*        soft    noproc  10240
	*        hard    noproc  10240
	*        soft    nofile  10240
	*        hard    nofile  10240

	就是限制了任意用户的最大线程数和文件数为10240。
