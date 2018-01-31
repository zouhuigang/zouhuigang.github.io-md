### 查看已添加规则

	iptables -L -n

	iptables -nL --line-number

	iptables -nv --line-number

	iptables -t nat -nL



### 例子
	
	#允许所有本机向外的访问
	iptables -A OUTPUT -j ACCEPT

	#允许外部访问22端口
	iptables -A INPUT -p tcp --dport 22 -j ACCEPT

	#屏蔽ip访问本服务器
	iptables -l INPUT -s 123.45.6.7 -j DROP

	#减少不安全的端口连接
	iptables -A OUTPUT -p tcp --sport 31337 -j DROP

	#关闭端口 34163
	iptables -A INPUT -p tcp --dport 34163  -j DROP

	#如果我们允许某个网段下的所有ip都可以访问的话比如1.2.3.[0-255]，我们可以
	iptables -A INPUT -s 1.2.3.0/24 -p tcp --dport -j ACCEPT

	#关闭8080端口,禁止所有的ip访问
	iptables -A INPUT -p tcp --dport 8080 -j DROP
	#只让183.193.113.185访问8080端口,ip在前，禁止端口在后
	iptables -A INPUT -s 183.193.113.185 -p tcp --dport 8080 -j ACCEPT
	iptables -A INPUT -s 139.196.48.36 -p tcp --dport 8080 -j ACCEPT
	iptables -I  INPUT 5 -s 10.254.74.13 -p tcp --dport 8080 -j ACCEPT #kube-da
	iptables -I  INPUT 5 -s 127.0.0.1 -p tcp --dport 8080 -j ACCEP
	iptables -A INPUT -p tcp --dport 8080 -j DROP
	例如：
	ACCEPT     tcp  --  183.193.113.185      0.0.0.0/0            tcp dpt:8080
	DROP       tcp  --  0.0.0.0/0            0.0.0.0/0            tcp dpt:8080


	#删除规则
	iptables -D INPUT 2 #删除第二条规则

	#插入第4条

	iptables -I  INPUT 4 -s 139.196.48.36 -p tcp --dport 8080 -j ACCEPT
	iptables -I  INPUT 4 -s 139.196.16.67 -p tcp --dport 8080 -j ACCEPT

	#添加局域网的ip段，拨号上网ip会变，网段不会变
	iptables -I  INPUT 4  -s 101.81.0.0/16 -p tcp --dport 8080 -j ACCEPT
	
	单个IP的命令是
	iptables -I INPUT -s 211.1.0.0 -j DROP	
	
	封IP段的命令是
	iptables -I INPUT -s 211.1.0.0/16 -j DROP
	iptables -I INPUT -s 211.2.0.0/16 -j DROP
	iptables -I INPUT -s 211.3.0.0/16 -j DROP
	
	封整个段的命令是
	iptables -I INPUT -s 211.0.0.0/8 -j DROP
	
	封几个段的命令是
	iptables -I INPUT -s 61.37.80.0/24 -j DROP
	iptables -I INPUT -s 61.37.81.0/24 -j DROP
	
	

### 测试端口的可用性

	#https://nmap.org/man/zh/
	#nmap-ncat 
	yum install -y nc

	下面的命令会检查远程主机 192.168.5.10 上是否打开了端口 80、22 和 21(我们也可以使用主机名)：

	nc  -v 139.196.16.67 8080
	端口未打开返回状态为非0


方法2.

	yum install nmap -y
	nmap ip -p port 测试端口
	 
	nmap ip 显示全部打开的端口
	 
	根据显示close/open确定端口是否打开。

测试:

	[root@k8s-master ~]# nmap 139.196.16.67 -p 8080

	Starting Nmap 6.40 ( http://nmap.org ) at 2017-07-08 08:23 CST
	Nmap scan report for 139.196.16.67
	Host is up (0.00037s latency).
	PORT     STATE SERVICE
	8080/tcp open  http-proxy
	
	Nmap done: 1 IP address (1 host up) scanned in 0.07 seconds
	[root@k8s-master ~]# 


state状态说明：

	open意味着目标主机的应用程序在监听(listen)
	closed意味着端口没有监听，但随时可能打开。
	filtered意味着firewall,阻碍了端口访问。
	unfiltered表示无法确定开放与否。


###端口问题

查看端口被关闭：

	PORT     STATE  SERVICE
	9090/tcp closed zeus-admin

	netstat -tln | grep 9090
	netstat -apn|grep 9090

	配置正确
	
	kubernetes apiserver-host主机地址

	iptables -I  INPUT 5 -s 10.254.74.13 -p tcp --dport 8080 -j ACCEPT #kube-da


### 保存问题

	首先iptables -L -n看一下配置是否正确。
	没问题后，先不要急着保存，因为没保存只是当前有效，重启后就不生效，这样万一有什么问题，可以后台强制重启服务器恢复设置。

	#保存
	[root@woxplife ~]# service iptables save
 
	#开机启动
	[root@woxplife ~]# systemctl enable iptables.service

参考文档：

[http://www.cnblogs.com/bethal/p/5806525.html](http://www.cnblogs.com/bethal/p/5806525.html)
[https://stackoverflow.com/questions/11222222/unkown-service-running-on-my-server](https://stackoverflow.com/questions/11222222/unkown-service-running-on-my-server)

[http://dockone.io/question/1266#!answer_form](http://dockone.io/question/1266#!answer_form)