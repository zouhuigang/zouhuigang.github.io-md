---
#标题
title: "ssh"
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
tags: [post,linux]
#分类
categories: [post,linux]
#作者
author: "邹慧刚"
---
重新生成ssh key:

ssh-keygen 

回车，在/home/zhg/.ssh下生成id_rsa,id_rsa.pub

将id_rsa.pub里面的内容复制进bitbucket.org的ssh key

---

解决navicat乱码问题,然后打开navicat找到选项，将字体设置为文泉驿-微米黑

sudo apt-get install ttf-wqy-microhei  #文泉驿-微米黑

sudo apt-get install ttf-wqy-zenhei  #文泉驿-正黑

sudo apt-get install xfonts-wqy #文泉驿-点阵宋体

---
解决sogou输入法安装之后，没法用的问题，安装完成sogou输入法后，打开fcitx配置，将搜狗输入法置顶，即可使用。

如果没有fcitx配置,可以安装一下:

    sudo add-apt-repository ppa:fcitx-team/nightly // 添加FCITX仓库.

    sudo apt-get update // 更新仓库.

    sudo apt-get install fcitx // 安装fcitx输入法框架.

---

复制windows字体到linux下面


以窗口的形式打开当前目录：

nautilus $PWD

cd /usr/share/fonts
sudo cp -r /media/zhg/mobile-store/字体/** /usr/share/fonts/chinese

cd /usr/share/fonts/chinese

### 刷新字体


mkfontscale
mkfontdir
fc-cache -fv




### golang

包管理

    go get -u github.com/golang/dep/cmd/dep

    sudo apt install go-dep

    $ sudo go get github.com/Masterminds/glide
    $ sudo go install github.com/Masterminds/glide