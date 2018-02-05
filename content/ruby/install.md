### 安装最新版本的ruby

下载地址:

	https://www.ruby-lang.org/en/downloads/


安装:
	mkdir -p /home/ruby && cd /home/ruby
	wget https://cache.ruby-lang.org/pub/ruby/2.5/ruby-2.5.0.tar.gz

	tar xzvf ruby-2.5.0.tar.gz
	cd ruby-2.5.0
	./configure  --prefix=/usr/local/ruby
	make 
	sudo make install

./configure  --prefix=/usr/local/ruby
是将ruby安装到指定目录，安装的时候开始没有使用root用户安装，出现问题，于是切换到root用户执行 make && make install



### 添加环境变量
vi /etc/profile

添加在最后添加

	export RUBY_HOME=/usr/local/ruby
	export PATH=$PATH:$RUBY_HOME/bin

保存,然后刷新环境变量

	source /etc/profile
	ruby -v




### 安装rubygems

	下载rubygems: http://www.oschina.net/news/19237/rubygems-v-185
	
	tar xzvf rubygems-1.8.5.tgz
	cd rubygems-1.8.5/
	ruby setup.rb
	gem --version

