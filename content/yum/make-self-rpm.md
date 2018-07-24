---
#标题
title: "make-self-rpm"
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
### 自己制作rpm软件包，方便使用yum安装软件


### 安装fpm打包工具

### 1.安装ruby，因为fpm是ruby编写的。

	[root@k8s-master1 ~]# ruby -v
	ruby 2.5.0p0 (2017-12-25 revision 61468) [x86_64-linux]

我的机子上，前几天已经安装好了ruby,安装方法可查看本库ruby模块。


### 2.安装fpm

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



所有配置:

	-f :强制覆盖[覆盖同名rpm包]
	-n :指定的rpm包名
	-p :指定的rpm包文件放置位置
	-v :指定的rpm包版本
	-d :指定依赖的软件   ( [-d 'name'] or [-d 'name > version'] 例子: -d 'libstdc++ >= 4.4.3')
	-a :指定系统架构,如果是noarch则为'-a all' 或者 '-a native' [x86_64] 当软件不区分64位或32位的时候可以 noarch
	-s :指定INPUT的数据类型 (["-s dir"] 省略数据类型)
	-m :指定打包人员[Packager]  ([ -m 'user'])
	-C :指定打包的相对路径,类似于buildroot. 譬如-C /tmp/apr/ 而打包机器的数据包路径是/tmp/apr/{opt,usr,etc} 那安装这个rpm包后,在本地的数据就是/opt/,/usr/,/etc/
	-t :指定需要制作成什么包,可选项有(deb,rpm,solaris,etc)
	    支持的源类型:：
	        "dir" "rpm" "gem" "python" "empty" "tar" "deb" "cpan" "npm" "osxpkg" "pear" "pkgin" "virtualenv" "zip"
	    支持的目标类型:
	        "rpm" "deb" "solaris" "puppet" "dir" "osxpkg" "p5p" "puppet" "sh" "solaris" "tar" "zip"
	--description         :软件包描述
	--conflicts         :指定冲突软件
	--url                 :指定站点[惯例都是添加软件的官网 例如: --url "http://www.cnblog.com/roach57" ]
	--verbose             :安装过程详细打印
	--after-install     :包安装之后执行的脚本 也可写作 --post-install FILE
	--before-install     :包安装之前执行的脚本 
	--after-remove         :包卸载之后执行的脚本
	--before-remove     :包卸载之前执行的脚本
	--after-upgrade     :包更新之后执行的脚本[仅支持 deb 和 rpm 这两种包]
	--before-upgrade     :包更新之前执行的脚本
	--iteration         :发布序号[就是rpm包里面的release]
	--epoch             :纪元  [不知道干嘛用的]
	--no-rpm-sign        :不使用rpm签名   Signature
	--license             :证书许可 [可选项有 'BSD(开源软件)' 'GPLv2(自由软件)' 'MIT' 'Public Domain(公共域)' 'Distributable(贡献)' 'commercial(商业)' 'Share(共享)等',一般的开发都写'BSD'或'GPL']
	--vendor             :供应商名称 [ --vendor 'roach57@163.com']
	--no-depends         :代表没有任何依赖包,和-d是对立的,不能共用
	--config-files         :指定配置文件,可以指定目录[递归]
	--directories         :指定包目录
	--category             :软件所属的类别[这是个什么软件]下面有个对应的表格:
	    [参考这个文件 /usr/share/doc/rpm-x.x.x/GROUPS ]
	    Amusements/Games [娱乐/游戏]
	    Amusements/Graphics [娱乐/图形]
	    Applications/Archiving [应用/文档]
	    Applications/Communications [应用/通讯]
	    Applications/Databases [应用/数据库]
	    Applications/Editors [应用/编辑器]
	    Applications/Emulators [应用/仿真器]
	    Applications/Engineering [应用/工程]
	    Applications/File [应用/文件]
	    Applications/Internet [应用/因特网]
	    Applications/Multimedia [应用/多媒体]
	    Applications/Productivity [应用/产品]
	    Applications/Publishing [应用/印刷]
	    Applications/System [应用/系统]
	    Applications/Text [应用/文本]
	    Development/Debuggers [开发/调试器]
	    Development/Languages [开发/语言]
	    Development/Libraries [开发/函数库]
	    Development/System [开发/系统]
	    Development/Tools [开发/工具]
	    Documentation [文档]
	    System Environment/Base [系统环境/基础]
	    System Environment/Daemons [系统环境/守护]
	    System Environment/Kernel [系统环境/内核]
	    System Environment/Libraries [系统环境/函数库]
	    System Environment/Shells [系统环境/接口]
	    User Interface/Desktops [用户界面/桌面]
	    User Interface/X [用户界面/X窗口]
	    User Interface/X Hardware Support [用户界面/X硬件支持]


组成格式:

	roach-1.0.1-57.el6.x86_64.rpm
	  |    |     |       |     |
	软件名称|     |       |     |
	     版本号   |       |　　  |
	           发布号     |      |
	                   硬件平台  |
	                            扩展名

例子备注:
    roach  :软件名称
    1.0.1  :软件版本号
    57.el6 :发布号主要是对软件存在的bug或漏洞进行修补,在软件功能上并没有变化,el6指的是rhel6系统中发布
    x86_64 :指64位的PC架构,另外还有'i386' 'i686' 等32位的PC架构,noarch是指不区分硬件架构
    rpm    :扩展名


### 3.制作软件包

	mkdir -p /nfs_file/software/rpm && cd /nfs_file/software/rpm
	fpm -s dir -t rpm -n ZphoneServer -v 1.0.0 -f /nfs_file/software/zphone 



查看：

	[root@k8s-master1 rpm]# ls
	ZphoneServer-1.0.0-1.x86_64.rpm

居然能生成一个rpm包了。


之前还对上面的命令一知半解，不知道啥情况的，看了很多帖子，都搞的莫名其妙，原来fpm包只是将软件的安装目录下的所有文件，打包放进了rpm里面了。所以什么也在这之前，什么也不需要做，只需要有个等待打包进rpm的目录即可。


	将rpm包下载下来，解压再解压之后，发现就是那个目录下所有的文件。


### 4.安装rpm软件


	yum localinstall ZphoneServer-1.0.0-1.x86_64.rpm

发现只有个二进制文件在这里，因为没有把services文件打包进rpm，所以目前不能用systemctl start/stop ZphoneServer 来管理软件。


### 5.卸载rpm软件

	yum remove -y ZphoneServer

### 6.启动rpm软件

这次重新将启动文件(ZphoneServer.service)也打包进rpm包里面

	fpm -s dir -t rpm -n ZphoneServer -v 1.0.0 --config-files  /usr/lib/systemd/system/ZphoneServer.service  -f /nfs_file/software/zphone   

查看rpm包文件

	rpm -qpl ZphoneServer-1.0.0-1.x86_64.rpm

显示:

	[root@k8s-master1 rpm]# rpm -qpl ZphoneServer-1.0.0-1.x86_64.rpm 
	/nfs_file/software/zphone/ZphoneServer
	/nfs_file/software/zphone/phone.dat
	/usr/lib/systemd/system/ZphoneServer.service


将rpm包上传到另外一台服务器，安装rpm包:

	
	yum localinstall -y ZphoneServer-1.0.0-1.x86_64.rpm

启动:
	
	systemctl start ZphoneServer
	systemctl status ZphoneServer
	

注:此处不再需要 systemctl daemon-reload 重新加载配置





问题汇总:

Q:

	Need executable 'rpmbuild' to convert dir to rpm {:level=>:error}

A:

	出现这种情况就是需要安装rpm-build
	yum install -y rpm-build



参考文档:

[https://darknode.in/linux/static-pack-apps/](https://darknode.in/linux/static-pack-apps/)
[http://www.iersai.com/archives/63](http://www.iersai.com/archives/63)
[http://www.cnblogs.com/Roach57/p/5130283.html](http://www.cnblogs.com/Roach57/p/5130283.html)
[https://github.com/jordansissel/fpm/issues/463](https://github.com/jordansissel/fpm/issues/463)
