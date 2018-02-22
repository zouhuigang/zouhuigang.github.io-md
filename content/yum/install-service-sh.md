---
title: "Install Service Sh"
date: 2018-02-08T13:35:30+08:00
draft: false
tags: ["yum", "service"]
categories: ["linux"]
---


### 自动生成service脚本

	
	#!/bin/bash
	touch /usr/lib/systemd/system/gdash.service
	echo "[Unit]" > /usr/lib/systemd/system/gdash.service
	echo "Description=gdash: GlusterFS Dashboard">> /usr/lib/systemd/system/gdash.service
	echo "After=glusterd.service" >> /usr/lib/systemd/system/gdash.service
	echo >> /usr/lib/systemd/system/gdash.service
	echo "[Service]" >> /usr/lib/systemd/system/gdash.service
	echo "type=simple" >> /usr/lib/systemd/system/gdash.service
	echo "ExecStart=/usr/bin/gdash -p 80" >> /usr/lib/systemd/system/gdash.service
	echo >> /usr/lib/systemd/system/gdash.service
	echo "[Install]" >> /usr/lib/systemd/system/gdash.service
	echo "WantedBy=multi-user.target" >> /usr/lib/systemd/system/gdash.service
	systemctl start gdash
	systemctl enable gdash

或者:

	#!/usr/bin/bash
	/usr/bin/cp *.service /usr/lib/systemd/system/
	/usr/bin/cp *.timer /usr/lib/systemd/system/



### 使用yum 安装rpm包的时候，不需要用sh脚本来生成service，只需要把/usr/lib/systemd/system/要用到的service打包进rpm里面就行。
