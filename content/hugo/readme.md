---
title: "README"
date: 2018-01-30T17:13:58+08:00
tags: ["post", "tag2", "tag3"]
categories: ["index"]
draft: true
---

### hugo

Hugo是由Go语言实现的静态网站生成器。简单、易用、高效、易扩展、快速部署。


其他方案:

	github page+ hexo 




### 下载

https://github.com/gohugoio/hugo/releases

得到二进制文件hugo.exe,将hugo.exe复制到D:\software\hugo目录下，然后设置环境变量（	D:\software\hugo；）。




### 使用

	hugo new site /path/to/site

1.cmd打开命令行

	cd D:\mnt\zouhuigang.github.io

	hugo new site D:\mnt\zouhuigang.github.io\pp
	cd D:\mnt\zouhuigang.github.io\pp
	hugo new about.md


about.md 自动生成到了 content/about.md ，打开 about.md 看下:

	---
	title: "About"
	date: 2018-01-30T15:53:11+08:00
	draft: true
	---

2.创建文章

创建第一篇文章，放到 post 目录，方便之后生成聚合页面。

	hugo new post/first.md


3.安装皮肤

	# 创建 themes 目录
	$ cd themes
	$ git clone https://github.com/balaramadurai/hugo-travelify-theme.git


4.运行Hugo

在你的站点根目录执行 Hugo 命令进行调试：

	$ hugo server --theme=hugo-travelify-theme --buildDrafts

（注明：v0.15 版本之后，不再需要使用 --watch 参数了）

浏览器里打开： http://localhost:1313


5.部署

假设你需要部署在 GitHub Pages 上，首先在GitHub上创建一个Repository，命名为：coderzh.github.io （coderzh替换为你的github用户名）。

在站点根目录执行 Hugo 命令生成最终页面：

	$ hugo --theme=hugo-travelify-theme --baseUrl="http://zouhuigang.github.io/"

（注意，以上命令并不会生成草稿页面，如果未生成任何文章，请去掉文章头部的 draft=true 再重新生成。）

如果一切顺利，所有静态页面都会生成到 public 目录，将pubilc目录里所有文件 push 到刚创建的Repository的 master 分支。
	






	