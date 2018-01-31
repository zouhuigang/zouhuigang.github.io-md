### 单实例的局限

目前，我所遇见的问题是在单台mysql中，接口源源不断的进来数据，导致一台mysql的写入并发在高峰时，支持每秒N多数据的写入。


参考：

	https://github.com/jumpjumpbean/k8smysqlcluster.git
	http://www.jumpbeandev.com/2017/06/01/mysqlcluster/#more
	https://blog.yaodataking.com/2016/02/25/mycat-performance-test/
	https://blog.yaodataking.com/2016/04/03/jconsole-remote-mycat/
	https://blog.yaodataking.com/2016/01/17/mycat-mysql-docker-sample1/#2-1-%E4%B8%8B%E6%8B%89Image


### k8s上部署
---

##### 构建docker镜像

master:

	cd /d/mnt/zouhuigang.github.io/mysql集群方案/mycat-master-slave/Dockerfiles/master
	docker build -t mysql-master -f ./Dockerfile .
	docker tag mysql-master registry.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/mysql:master-1.0.0
	docker push registry.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/mysql:master-1.0.0


slaver:

	cd /d/mnt/zouhuigang.github.io/mysql集群方案/mycat-master-slave/Dockerfiles/slave
	docker build -t mysql-slave -f ./Dockerfile .
	docker tag mysql-slave registry.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/mysql:slave-1.0.0
	docker push registry.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/mysql:slave-1.0.0



### 测试主从复制

	kubectl create -f ./rc-mysql-master
	kubectl create -f ./svc-mysql-master
	kubectl create -f ./rc-mysql-slave
	kubectl create -f ./svc-mysql-slave


参考文档:

	https://www.jianshu.com/p/509b65e9a4f5


mycat:

基于gaven/mycat镜像，修改schema.xml进行一主一从读写分离自动切换配置，修改server.xml添加db用户名密码等

	cd /d/mnt/zouhuigang.github.io/mysql集群方案/mycat-master-slave/Dockerfiles/mycat
	docker build -t mycat-v1.0.2 -f ./Dockerfile .
	docker tag mycat-v1.0.2 registry.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/mycat:v1.0.2
	docker push registry.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/mycat:v1.0.2

创建:

	kubectl create -f ./rc-mycat
	kubectl create -f ./svc-mycat



docker:

	https://www.jianshu.com/p/b73cf127a4b9

启动:

	docker run --name mysql-mycat  -it   \
	-v /mnt/a/mycat/conf:/usr/local/mycat/conf  \
	-v /mnt/a/mycat/logs:/usr/local/mycat/logs \
	-p 8066:8066 -p 9066:9066  \
	registry.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/mycat:v1.0.2


登录:

	mysql -uroot -p -h127.0.0.1 -P8066  #密码: root

	mysql -uroot -p -h127.0.0.1 -P9066  #密码: root
	show @@heartbeat;
	show @@database;   #查看逻辑数据库
	show @@datanode;   #查看分片节点
	show @@server;     #查看服务器状态
	show @@version;    #查看版本



日志debug开启:

	将Mycat的日志模式改为debug模式，执行sql后，到日志查看执行的节点ip就知道是不是自己设置的读节点 
	在logs文件夹下执行 tail -f mycat.log,你用navicat增删改成会在控制台刷出日志。注意，conf文件夹下log4j。xml里的info改成debug


将mycat的日志输出级别改完debug，在conf/log4j2.xml里配置

查询语句不要加事务，否则读操作会被分发到写服务器上。

日志:

	将info改为debug,即可查看日志


 


	