
### Quick Start

    1.Download Caddy and put it in your PATH.(下载caddy并加入环境变量)

    2.cd to your website's directory.(打开网站根目录)

    3.Run caddy.(运行caddy)


###安装



	mkdir /caddy
	wget https://github.com/mholt/caddy/releases/download/v0.10.6/caddy_v0.10.6_linux_amd64.tar.gz
	tar zxvf caddy_v0.10.6_linux_amd64.tar.gz 

加入环境变量

	vi ~/.bash_profile
	修改Path为：
	PATH=$PATH:$HOME/bin:/caddy
	source ~/.bash_profile


### 配置文件
	mkdir /caddy-www && cd /caddy-www
	echo 'localhost:8888'>>Caddyfile
	#caddy -conf ../path/to/Caddyfile
	caddy


### 浏览器打开

	ip:8888



[https://github.com/Unknwon/wuwen.org/issues/12](https://github.com/Unknwon/wuwen.org/issues/12)
[https://github.com/Unknwon/wuwen.org/issues/11](https://github.com/Unknwon/wuwen.org/issues/11)