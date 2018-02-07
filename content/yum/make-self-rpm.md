---
title: "Make Self Rpm"
date: 2018-02-07T22:44:45+08:00
draft: true
---


### 自己制作rpm软件包，方便使用yum安装软件


### 安装fpm打包工具

1.安装ruby，因为fpm是ruby编写的。

	[root@k8s-master1 ~]# ruby -v
	ruby 2.5.0p0 (2017-12-25 revision 61468) [x86_64-linux]

我的机子上，前几天已经安装好了ruby,安装方法可查看本库ruby模块。


2.安装fpm

	gem install fpm


版本:

	[root@k8s-master1 ~]# fpm --version
	1.9.3

常用参数:

	-s          指定源类型
	-t          指定目标类型，即想要制作为什么包
	-n          指定包的名字
	-v          指定包的版本号
	-C          指定打包的相对路径  Change directory to here before searching forfiles
	-d          指定依赖于哪些包
	-f          第二次打包时目录下如果有同名安装包存在，则覆盖它
	-p          输出的安装包的目录，不想放在当前目录下就需要指定
	--post-install      软件包安装完成之后所要运行的脚本；同--after-install
	--pre-install       软件包安装完成之前所要运行的脚本；同--before-install
	--post-uninstall    软件包卸载完成之后所要运行的脚本；同--after-remove
	--pre-uninstall     软件包卸载完成之前所要运行的脚本；同--before-remove


3.制作软件包

	mkdir -p /nfs_file/software/rpm && cd /nfs_file/software/rpm
	fpm -s dir -t rpm -n ZphoneServer -v 1.0.0 -f /nfs_file/software/zphone 



查看：

	[root@k8s-master1 rpm]# ls
	ZphoneServer-1.0.0-1.x86_64.rpm

居然能生成一个rpm包了。


之前还对上面的命令一知半解，不知道啥情况的，看了很多帖子，都搞的莫名其妙，原来fpm包只是将软件的安装目录下的所有文件，打包放进了rpm里面了。所以什么也在这之前，什么也不需要做，只需要有个等待打包进rpm的目录即可。


	将rpm包下载下来，解压再解压之后，发现就是那个目录下所有的文件。




问题汇总:

Q:

	Need executable 'rpmbuild' to convert dir to rpm {:level=>:error}

A:

	出现这种情况就是需要安装rpm-build
	yum install -y rpm-build