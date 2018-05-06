---
title: "Grun"
description: "这是一段描述-zouhuigang.anooc.com"
author: 邹慧刚
date: 2018-04-10T09:51:15+08:00
categories:
  - "front"
tags: ["post", "Grun"]
draft: false
---




### 打开项目目录：

 	cd D:\www\yiyang\static


### 安装 CLI
	npm install -g grunt-cli  
	npm install


### 准备一份新的 Grunt 项目

一般需要在你的项目中添加两份文件：

	package.json 和 Gruntfile。


生成:

	npm init
	npm install -g grunt-init
	npm install grunt --save-dev


然后安装模板，目前有三种模板是由grunt官方做维护的，还有别的可在github上找到，或者你自己实现一个。 官方模板的安装命令如下：

	git clone git://github.com/gruntjs/grunt-init-gruntfile.git $HOME/.grunt-init/
	git clone git://github.com/gruntjs/grunt-init-jquery.git %HOME%/.grunt-init/
	git clone git://github.com/gruntjs/grunt-init-node.git $HOME/.grunt-init/

上面的$HOME是系统目录,例如C:\Users\mdshi\,也可以将它加入到环境变量

	HOME
	C:\Users\mdshi\

查看

	cmd-> echo %HOME%


windows:

	git clone git://github.com/gruntjs/grunt-init-gruntfile.git %USERPROFILE%/.grunt-init/gruntfile 



或者:

	git clone git://github.com/gruntjs/grunt-init-gruntfile.git %HOME%/.grunt-init/gruntfile

	grunt-init  %HOME%/.grunt-init/gruntfile




三种分别对应默认grunt模板，jquery插件的grunt模板，node包的grunt模板。

然后就可以适用grunt-init命令来初始化你的Gruntfile.js文件了，例如你要安装默认模板：

	grunt-init grunt-init-gruntfile #最后一个参数也可以是模板所在的文件夹
	grunt-init  %HOME%/.grunt-init/

它会问你一些问题，然后根据你的答案创建当前项目的Gruntfile.js文件。




### 安装iview-admin

	cd D:\www\yiyang\static2 && git clone https://github.com/iview/iview-admin.git
	cd D:\www\yiyang\static2\iview-admin
	安装依赖
	npm install --save-dev

	运行
	npm run dev
	
	打包
	npm run build
	

https://www.cnblogs.com/chaojidan/p/4239562.html?utm_source=tuicool&utm_medium=referral


http://blog.sae.sina.com.cn/archives/4106

https://www.cnblogs.com/vajoy/p/3983831.html



修复mysql:

	myisamchk

http://blog.itpub.net/29500582/viewspace-1301666/