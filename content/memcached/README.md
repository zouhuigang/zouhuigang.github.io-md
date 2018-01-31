### 安装memcached

	yum install -y memcached

启动

	systemctl start memcached

状态

	systemctl status memcached



### telnet登录

	yum install -y telnet # windows下使用cmd->optionalfeatures->打开telnet客户端即可使用

	telnet 127.0.0.1 11211  #输入quit退出telnet

### 列出所有keys

	stats items // 这条是命令,注意不要打成status，真tmd fuck

得到结果:

	STAT items:2:number 542
	STAT items:2:age 1807
	STAT items:2:evicted 0
	STAT items:2:evicted_nonzero 0
	STAT items:2:evicted_time 0
	STAT items:2:outofmemory 0
	STAT items:2:tailrepairs 0
	STAT items:2:reclaimed 144
	STAT items:2:expired_unfetched 137
	STAT items:2:evicted_unfetched 0
	STAT items:2:crawler_reclaimed 0
	STAT items:2:crawler_items_checked 0
	STAT items:2:lrutail_reflocked 0
	END


### 通过itemid获取key
>通过命令stats cachedump id 0获得key的值。0表示全部列出

	stats cachedump 2 0 #2是上面的items:2

得到:

	ITEM 57ee9db7e791826e7b5e0b005506702f [1 b; 1515727329 s]
	ITEM 419932cec5bf37cc297d777102f94a6d [1 b; 1515727325 s]
	ITEM 4420e6b3e6198018ad08cdd8b5a466c0 [1 b; 1515727323 s]
	ITEM e72fffbccd4e8c197b91fc44d3fcdc9b [1 b; 1515727323 s]
	ITEM 058a734bfb704f8703e15d9d97c13657 [1 b; 1515727319 s]
	ITEM 1adc3a90d0298bd4f288e5ee479166e3 [1 b; 1515727315 s]
	ITEM 9e1949731eb52626bde8e9acfc1c6798 [1 b; 1515727309 s]
	ITEM 4b2c7056c4ce9f67861b99d3f97a982a [1 b; 1515727306 s]
	ITEM 876a05b3695433fbe88f97bd2043c054 [1 b; 1515727302 s]
	ITEM 898610cc60d532be24629f82dc29c95c [1 b; 1515727294 s]
	ITEM d0e298fcffa2aab8dc327bc3a5e09251 [1 b; 1515727283 s]
	END

### 通过get获取key值

	get d0e298fcffa2aab8dc327bc3a5e09251




### 问题汇总

Q1：输入命令一直得到ERROR

A1：

	stats不要打成status
	https://groups.google.com/forum/#!topic/memcached/b_7Brn8WNOE