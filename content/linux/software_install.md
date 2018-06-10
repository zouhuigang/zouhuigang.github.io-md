notepadqq:

Ubuntu下的安装方法:

	sudo add-apt-repository ppa:notepadqq-team/notepadqq

	sudo apt-get update

	sudo apt-get install notepadqq

Ubuntu下的卸载方法:

	sudo apt-get remove notepadqq

s	udo add-apt-repository --remove ppa:notepadqq-team/notepadqq

shadowsocks 

	sudo add-apt-repository ppa:hzwhuang/ss-qt5 
	sudo apt-get update 
	sudo apt-get install shadowsocks-qt5
	
翻墙服务搭配好了，可以配合谷歌的插件:https://www.switchyomega.com/download.html使用，目的是在浏览器端，建立一个pac文件，pac文件指向shadowsocks代理软件，实现翻墙上网。

https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt

### 截屏工具 

	sudo apt-get install shutter
### 迁移数据库

	host = 47.100.76.132
	port = 3308
	user = root
	password = Zouhuigang199201
	dbname = anooc
	
备份:

mysqldump --databases anooc3 -uroot -P3308 -h47.100.76.132 -p | gzip > /mnt/anooc3-20180610.sql.gz
mysql -uanooc_root -p anooc -hrm-uf6m5938b7c1m2835.mysql.rds.aliyuncs.com < /mnt2/anooc-20180513 .sql

### 暂停网站

	kubectl scale golang-anooc-rc-name-v1 --replicas=0
	
### 百度云盘

	sudo apt install python-pip
	pip install bypy