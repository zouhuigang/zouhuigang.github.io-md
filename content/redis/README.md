
[windows版下载地址](https://github.com/MicrosoftArchive/redis/releases)

下载完成，安装即可！

### 启动Redis客户端

	cd D:\software\redis
	redis-cli

连接远程docker redis服务端:

	redis-cli -h 服务器地址 -p 端口 -a 密码

	redis-cli -h 192.168.99.100 -p 6379 -a "pass123" 


连接远程服务redis如图：

![image](./images/20171230223043.png)


### 使用说明

##### 设置:

	set 键 值

	set realname zouhuigang


##### 得到:

	get 键
	get realname


获取所有的键值:

	keys *

查看键值的类型:

	type 键

##### 删除:

	del 键
	del realname



### hash(type==hash)

> 如果type是hash,包含哈希表中所有字段的列表。 当 key 不存在时，返回一个空列

列出所有值:

	HKEYS  键
	hkeys myhash
	1) "field1"
	2) "field2"

设置hash值:

	HSET myhash field1 "foo"



库：

>Redis默认有16个库通过select命令切换，库之间隔离默认登录0号库

选择1号库:

	select 1




### Redis Sadd 命令 (type==set)

	redis 127.0.0.1:6379> SADD myset "hello"
	(integer) 1
	redis 127.0.0.1:6379> SADD myset "foo"
	(integer) 1
	redis 127.0.0.1:6379> SADD myset "hello"
	(integer) 0
	redis 127.0.0.1:6379> SMEMBERS myset
	1) "hello"
	2) "foo"



问题汇总:

Q1:

	(error) WRONGTYPE Operation against a key holding the wrong kind of value

A1:

	数据库中有一个key是usrInfo的数据存储的是Hash类型,执行数据

	操作的时候却使用了非Hash的操作方法，比如Sorted Sets里的方法。此时就会报

	ERR Operation against a key holding the wrong kind of value这个错误！


### 参考文档

[http://www.yiibai.com/redis/redis_quick_guide.html](http://www.yiibai.com/redis/redis_quick_guide.html)



