### 同步mysql数据到es中

### 环境Java 8
	
	yum update
	mkdir /usr/java && cd /usr/java

	wget --no-cookies --no-check-certificate --header "Cookie: gpw_e24=http%3A%2F%2Fwww.oracle.com%2F; oraclelicense=accept-securebackup-cookie" "http://download.oracle.com/otn-pub/java/jdk/8u144-b01/090f390dda5b47b9b721c7dfaa008135/jdk-8u144-linux-x64.tar.gz"

	tar -xvf jdk-8u144-linux-x64.tar.gz

	cd /usr/java/jdk1.8.0_144/
	alternatives --install /usr/bin/java java /usr/java/jdk1.8.0_144/bin/java 2
	alternatives --config java

	--  (选择最新的一个3)

	alternatives --install /usr/bin/jar jar /usr/java/jdk1.8.0_144/bin/jar 2
	alternatives --install /usr/bin/javac javac /usr/java/jdk1.8.0_144/bin/javac 2
	alternatives --set jar /usr/java/jdk1.8.0_144/bin/jar
	alternatives --set javac /usr/java/jdk1.8.0_144/bin/javac

	java -version

环境变量:

	  echo "export JAVA_HOME=/usr/java/jdk1.8.0_144" >> /etc/profile.d/myprofile.sh
	  echo "export JRE_HOME=/usr/java/jdk1.8.0_144/jre" >> /etc/profile.d/myprofile.sh
	  echo "export PATH=$PATH:/usr/java/jdk1.8.0_111/bin:/usr/java/jdk1.8.0_144/jre/bin" >> /etc/profile.d/myprofile.sh
	  chmod 755 /etc/profile.d/myprofile.sh
	  退出  (to leave the "root" account)
	  source /etc/profile.d/myprofile.sh
	 echo $JAVA_HOME


或者

	#jdk
	export JAVA_HOME=/usr/local/jdk1.8.0_111
	export PATH=$JAVA_HOME/bin:$PATH
	export CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar

立刻生效

	source /etc/profile

	echo $JAVA_HOME 如果返回为空，则需要配置环境变量


 logstash-input-jdbc插件是logstash 的一个个插件。

使用ruby语言开发

 下载插件过程中最大的坑是下载插件相关的依赖的时候下不动，因为国内网络的原因，访问不到亚马逊的服务器。

解决办法，改成国内的ruby仓库镜像。此镜像托管于淘宝的阿里云服务器上 ,[https://ruby.taobao.org/](https://ruby.taobao.org/)

	yum install gem

	gem sources --add https://ruby.taobao.org/ --remove https://rubygems.org/

	gem sources -l


说明：

	# 请确保只有 ruby.taobao.org
	如果 还是显示 https://rubygems.org/ 进入 home的 .gemrc 文件
	sudo vim ~/.gemrc 
	手动删除 https://rubygems.org/


### 修改Gemfile的数据源地址

	gem install bundler
	
	bundle config mirror.https://rubygems.org https://ruby.taobao.org





### ==其他软件，安装logstash,这个不是logstash-input-jdbc插件

	cd /opt/logstash/

	wget https://artifacts.elastic.co/downloads/logstash/logstash-5.2.1.tar.gz
	tar -zxvf logstash-5.2.1.tar.gz

测试是否安装成功
	
	cd logstash-5.2.1
	bin/logstash -e 'input { stdin { } } output { stdout {}}'



### 安装logstash-input-jdbc

	cd logstash-5.2.1
	#bin/logstash-plugin install logstash-input-jdbc
	bin/logstash-plugin install --local logstash-input-jdbc


### 列出已安装插件

	#bin/logstash-plugin list --verbose | grep elasticsearch

	bin/logstash-plugin list --verbose


### 使用,#这个包要在logstash服务器上存在包

	 首先为es安装mysql-connector插件,https://dev.mysql.com/downloads/connector/j/

	cd /opt/logstash/
	#curl -o mysql-connector-java-5.1.33.zip -L 'http://dev.mysql.com/get/Downloads/Connector-J/mysql-connector-java-5.1.33.zip/from/http://cdn.mysql.com/'
	#unzip mysql-connector-java-5.1.33.zip 
	wget https://dev.mysql.com/get/Downloads/Connector-J/mysql-connector-java-5.1.43.zip
	unzip mysql-connector-java-5.1.43.zip 

	bin/logstash -f jdbc.config

测试mysql-connector-java-5.1.33会提示报错，不能加载包，改成mysql-connector-java-5.1.43就好了


使用前先将jdbc.config上传到/opt/logstash/logstash-5.2.1目录下，在运行bin/logstash -f jdbc.config


### 问题汇总

Q1：An error occurred while installing logstash-core-event-java (5.2.1), and Bundler cannot continue.
Make sure that `gem install logstash-core-event-java -v '5.2.1'` succeeds before bundling.

A1：使用 bin/logstash-plugin install --local logstash-input-jdbc

参考文档：

[http://blog.csdn.net/yeyuma/article/details/50240595#quote](http://blog.csdn.net/yeyuma/article/details/50240595#quote)

[http://blog.csdn.net/fenglailea/article/details/56282414](http://blog.csdn.net/fenglailea/article/details/56282414)

[https://www.elastic.co/guide/en/logstash/current/plugins-inputs-jdbc.html](https://www.elastic.co/guide/en/logstash/current/plugins-inputs-jdbc.html)

[http://www.jianshu.com/p/bcd3a2210177](http://www.jianshu.com/p/bcd3a2210177)

[https://www.digitalocean.com/community/tutorials/how-to-install-java-on-centos-and-fedora](https://www.digitalocean.com/community/tutorials/how-to-install-java-on-centos-and-fedora)

[http://www.ruanyifeng.com/blog/2017/08/elasticsearch.html](http://www.ruanyifeng.com/blog/2017/08/elasticsearch.html)

[https://discuss.elastic.co/t/i-can-not-install-any-logstash-plug-in-in-es-5-0/64313/26](https://discuss.elastic.co/t/i-can-not-install-any-logstash-plug-in-in-es-5-0/64313/26)

[https://javabirder.wordpress.com/2016/02/21/install-java-8-centos/](https://javabirder.wordpress.com/2016/02/21/install-java-8-centos/)

[https://www.elastic.co/guide/en/logstash/current/offline-plugins.html](https://www.elastic.co/guide/en/logstash/current/offline-plugins.html)

[http://blog.csdn.net/laoyang360/article/details/51747266](http://blog.csdn.net/laoyang360/article/details/51747266)

[http://itblog.rrslj.com/logstash-input-jdbc-tong-bu-yuan-li-ji-xiang-guan-wen-ti-jie-du/](http://itblog.rrslj.com/logstash-input-jdbc-tong-bu-yuan-li-ji-xiang-guan-wen-ti-jie-du/)

[http://www.cnblogs.com/licongyu/p/5383334.html](http://www.cnblogs.com/licongyu/p/5383334.html)