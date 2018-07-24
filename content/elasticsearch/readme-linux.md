---
#标题
title: "readme-linux"
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
tags: [post,elasticsearch]
#分类
categories: [post,elasticsearch]
#作者
author: "邹慧刚"
---
https://www.elastic.co/cn/

环境：

	yum install -y java


1.下载

	cd ~
	wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-5.5.1.tar.gz


解压：

	tar -zxvf  elasticsearch-5.5.1.tar.gz

####或者yum安装(推荐使用此方法安装）

1.添加源：

	$ sudo vi /etc/yum.repos.d/elasticsearch.repo
	[elasticsearch-5.x]
	name=Elasticsearch repository for 5.x packages
	baseurl=https://artifacts.elastic.co/packages/5.x/yum
	gpgcheck=1
	gpgkey=https://artifacts.elastic.co/GPG-KEY-elasticsearch
	enabled=1
	autorefresh=1
	type=rpm-md

2.yum安装

	yum install elasticsearch -y

	systemctl daemon-reload

	systemctl enable elasticsearch.service


修改配置：

vi /etc/elasticsearch/jvm.options

	将
	#-Xms2g
	#-Xmx2g
	改为：
	-Xms1g 
	-Xmx1g

外网访问:

vi  /etc/elasticsearch/elasticsearch.yml

	network.host: 0.0.0.0



3.启动

	systemctl start elasticsearch.service
	
	重启动：
	systemctl restart elasticsearch

4.访问

	http://192.168.122.143:9200/

	curl 127.0.0.1:9200

输出信息：

	{
	  "name" : "jjC8h3r",
	  "cluster_name" : "elasticsearch",
	  "cluster_uuid" : "ROD_vAbIQjioGOZ5OEp_Fw",
	  "version" : {
	    "number" : "5.5.2",
	    "build_hash" : "b2f0c09",
	    "build_date" : "2017-08-14T12:33:14.154Z",
	    "build_snapshot" : false,
	    "lucene_version" : "6.6.0"
	  },
	  "tagline" : "You Know, for Search"
	}

关闭

	curl -XPOST 'http://localhost:9200/_shutdown'


### 现在可以使用es啦~~~~同步mysql，请安装logstash-input-jdbc=========
	

### 配置elasticsearch

	cd  elasticsearch-5.5.1/config/elasticsearch.yml
	vim elasticsearch.yml 
	cluster.name: myes-1   

设置数据目录权限

	sudo mkdir -p /data/es-date/
 	chown -R elasticsearch:elasticsearch /data/es-date/#这个是我们存放数据的目录，手动创建

2.启动：

	cd elasticsearch-5.5.1/bin
   	./elasticsearch

	或
	sudo ./bin/elasticsearch
	./bin/elasticsearch -d -Ecluster.name=my_cluster -Enode.name=node_1

	或

	ln -s /usr/local/elasticsearch-1.7.2 /usr/local/elasticsearch
	systemctl start elasticsearch

### 插件说明

Kibana:实现数据可视化。导览 Elastic Stack。

安装Kiabna：

	vi /etc/yum.repos.d/kibana.repo
	[kibana-5.x]
	name=Kibana repository for 5.x packages
	baseurl=https://artifacts.elastic.co/packages/5.x/yum
	gpgcheck=1
	gpgkey=https://artifacts.elastic.co/GPG-KEY-elasticsearch
	enabled=1
	autorefresh=1
	type=rpm-md

下载：

 	yum -y install kibana

外网访问:/etc/kibana/kibana.yml

	server.host: "0.0.0.0"

启动：

	[root@localhost ~]# /bin/systemctl daemon-reload
	[root@localhost ~]# sudo systemctl enable kibana
	[root@localhost ~]# systemctl start kibana

测试：

	ip:5601 有界面就安装成功了

	http://192.168.122.143:5601



X-Pack:是一个Elastic Stack的扩展，将安全，警报，监视，报告和图形功能包含在一个易于安装的软件包中。

安装，在Elasticsearch 的安装目录( /usr/share/elasticsearch/ )下运行:

	[root@localhost elasticsearch]# bin/elasticsearch-plugin install x-pack

	安装过程中跳出选项选择y即可

如果你在Elasticsearch已禁用自动索引的创建，
在elasticsearch.yml配置action.auto_create_index允许X-pack创造以下指标：

	action.auto_create_index: .security,.monitoring*,.watches,.triggered_watches,.watcher-history*

Kibana安装，在Kibana的安装目录( /usr/share/kibana/ )下运行，Kibana下载X-Pack

	[root@localhost kibana]# bin/kibana-plugin install x-pack

	Attempting to transfer from x-pack
	Attempting to transfer from https://artifacts.elastic.co/downloads/kibana-plugins/x-pack/x-pack-5.5.1.zip
	Transferring 119276972 bytes....................
	Transfer complete
	Retrieving metadata from plugin archive
	Extracting plugin archive
	Extraction complete
	Optimizing and caching browser bundles...
	Plugin installation complete

安装过程有点长，慢慢等待成功安装....注意需要关闭ES服务器，如果已经启动的话，否则安装失败。

验证X-Pack

重新启动Elasticsearch和Kibana服务
然后访问http://localhost:9200/ 和 http://localhost:5601/ 此时需要输入用户名和密码登录，默认的用户名: elastic 密码: changeme

	systemctl restart elasticsearch
	systemctl restart kibana



### ES-head

    git clone git://github.com/mobz/elasticsearch-head.git
    cd elasticsearch-head
    npm install
    npm run start

    open http://192.168.122.143:9100/

或者直接用docker 

 	docker run -p 9100:9100 registry.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/elasticsearch-head:5

或者使用谷歌的扩展：

	chrome-extension://ffmkiejjmecolpfloofpjologoblkegm/elasticsearch-head/index.html

如果没有npm命令，则安装

	git clone https://github.com/cnpm/nvm.git
	cd nvm && source nvm.sh
	nvm  list-remote #列出所有版本的node
	nvm install v8.4.0 



### 健康检查

	curl -u elastic http://127.0.0.1:9200/_cat/health


### 数据同步

 	logstash-input-jdbc

问题：

Q1：OpenJDK 64-Bit Server VM warning: If the number of processors is expected to increase from one, then you should configure the number of parallel GC threads appropriately using -XX:ParallelGCThreads=N

A1：将jvm.options改成
#-Xms2g
#-Xmx2g
-Xms256m
-Xmx256m


参考文档：

[https://es.xiaoleilu.com/010_Intro/10_Installing_ES.html](https://es.xiaoleilu.com/010_Intro/10_Installing_ES.html)

[http://abcdocker.blog.51cto.com/11255059/1907687](http://abcdocker.blog.51cto.com/11255059/1907687)

[http://blog.csdn.net/beitiandijun/article/details/56019960](http://blog.csdn.net/beitiandijun/article/details/56019960)

[http://qiita.com/tjinjin/items/7ea3ebd228748f9d5224](http://qiita.com/tjinjin/items/7ea3ebd228748f9d5224)

[http://www.jianshu.com/p/e49ed6acd7da](http://www.jianshu.com/p/e49ed6acd7da)

[http://www.cnblogs.com/xing901022/p/5962722.html](http://www.cnblogs.com/xing901022/p/5962722.html)

[http://www.jianshu.com/p/eaf54202aa08](http://www.jianshu.com/p/eaf54202aa08)

[https://www.zghhome.cn/?p=350](https://www.zghhome.cn/?p=350)