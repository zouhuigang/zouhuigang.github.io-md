---
#标题
title: "install"
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
tags: [post,ruby]
#分类
categories: [post,ruby]
#作者
author: "邹慧刚"
---
### 安装最新版本的ruby

下载地址:

	https://www.ruby-lang.org/en/downloads/


安装:

	mkdir -p /home/ruby && cd /home/ruby
	wget https://cache.ruby-lang.org/pub/ruby/2.5/ruby-2.5.0.tar.gz

	tar xzvf ruby-2.5.0.tar.gz
	cd ruby-2.5.0
	./configure  --prefix=/usr/local/ruby
	make 
	sudo make install

./configure  --prefix=/usr/local/ruby
是将ruby安装到指定目录，安装的时候开始没有使用root用户安装，出现问题，于是切换到root用户执行 make && make install



### 添加环境变量
vi /etc/profile

添加在最后添加

	export RUBY_HOME=/usr/local/ruby
	export PATH=$PATH:$RUBY_HOME/bin

保存,然后刷新环境变量

	source /etc/profile
	ruby -v




### 安装rubygems

	下载rubygems: http://www.oschina.net/news/19237/rubygems-v-185

	cd /home/ruby && wget http://production.cf.rubygems.org/rubygems/rubygems-1.8.5.tgz
	
	tar xzvf rubygems-1.8.5.tgz
	cd rubygems-1.8.5/
	ruby setup.rb
	gem --version


### 问题汇总


Q1:

	[root@ceph-admin rubygems-1.8.5]# gem install fpm
	ERROR:  Loading command: install (LoadError)
	    cannot load such file -- zlib
	ERROR:  While executing gem ... (NameError)
	    uninitialized constant Gem::Commands::InstallCommand
	Did you mean?  Gem::InstallUpdateOptions


A1:

	#yum install -y gcc-c++ patch readline readline-devel zlib zlib-devel libyaml-devel libffi-devel openssl-devel
	#yum install -y make bzip2
	#yum install -y iconv-devel 
	yum install zlib-devel 
	如果安装不成功，可以更改yum源为阿里的源。
	
	发现还是不行，查询发现需要，集成zlib库到ruby环境
	cd /home/ruby/ruby-2.5.0
	cd ext/zlib
	ruby extconf.rb
	//在操作下一步之前需要修改Makefile文件中的zlib.o: $(top_srcdir)/include/ruby.h,将$(top_srcdir)修改为../..如下
	//zlib.o: ../../include/ruby.h
	//这一步如果不修改，make时会爆出另外一个错误
	//make:*** No rule to make target `/include/ruby.h', needed by `zlib.o'.  Stop
	make && make install

调试命令：

	gem install fpm  --debug



Q2:
	
gem install fpm安装时，报错

	ERROR:  While executing gem ... (FrozenError)
    can't modify frozen String

A2：

	gem update --system

查看版本:

	[root@mon1 zlib]# gem --version
	2.7.6


Q3:

	ERROR:  While executing gem ... (Gem::Exception)
    Unable to require openssl, install OpenSSL and rebuild Ruby (preferred) or use non-HTTPS sources

A3:

	yum install openssl-devel  
	cd /home/ruby/ruby-2.5.0/ext/openssl
	ruby extconf.rb
	//同样修改Makefile中的$(top_srcdir)为../..
	make && make install


报错：
	make: *** No rule to make target `/include/ruby.h', needed by `ossl.o'.  Stop.
	make: *** No rule to make target `/include/ruby.h', needed by `ossl_asn1.o'.  Stop.

vi Makefile

报错：


![images](../images/20180228162902.png)
![images](../images/20180228163835.png)

这个Makefile中有很多.o文件用上了$(top_srcdir) ，因此可在文件的第一行定义一下(在Makefile第一行添加)：

	top_srcdir = /home/ruby/ruby-2.5.0/include
	 


参考文档：

[https://qiita.com/aTaroAsari/items/9be9f905772637fecef6](https://qiita.com/aTaroAsari/items/9be9f905772637fecef6)
[https://www.cnblogs.com/xjh713/p/7458437.html](https://www.cnblogs.com/xjh713/p/7458437.html)
[http://blog.csdn.net/feinifi/article/details/78251486](http://blog.csdn.net/feinifi/article/details/78251486)
[http://blog.csdn.net/kenera/article/details/6524557](http://blog.csdn.net/kenera/article/details/6524557)



