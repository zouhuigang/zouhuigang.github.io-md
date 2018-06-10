### 1.安装

	cd /home/zhg/software/
	git clone https://github.com/rofl0r/proxychains-ng.git
	cd proxychains-ng
	sudo ./configure --prefix=/usr --sysconfdir=/etc
	sudo make && make install
	sudo make install-config
	cd .. && rm -rf proxychains-ng

### 2.编辑proxychains配置

	vi /etc/proxychains.conf
	//将socks4  127.0.0.1 9050改为
	socks5  127.0.0.1 1080  //1080改为你自己的端口


### 3.检查是否成功

	proxychains4 curl ip.cn

显示:

![images](./选区_004.png)



### 问题汇总

Q:
	
	proxychains can't load process....: No such file or directory


A:


	由于curl命令没有安装，找不到curl
	
