## 安装java环境

	yum update
	mkdir /usr/java && cd /usr/java
	
	wget --no-cookies --no-check-certificate --header "Cookie: gpw_e24=http%3A%2F%2Fwww.oracle.com%2F; oraclelicense=accept-securebackup-cookie" "http://download.oracle.com/otn-pub/java/jdk/8u152-b16/aa0333dd3019491ca4f6ddbe78cdb6d0/jdk-8u152-linux-x64.tar.gz"

	#注，上面的地址如果找不到，可以在http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html中找最新的版本
	
	tar -xvf jdk-8u152-linux-x64.tar.gz
	
	cd /usr/java/jdk1.8.0_152/
	alternatives --install /usr/bin/java java /usr/java/jdk1.8.0_152/bin/java 2
	alternatives --config java
	
	--  (选择最新的一个1)
	
	alternatives --install /usr/bin/jar jar /usr/java/jdk1.8.0_152/bin/jar 2
	alternatives --install /usr/bin/javac javac /usr/java/jdk1.8.0_152/bin/javac 2
	alternatives --set jar /usr/java/jdk1.8.0_152/bin/jar
	alternatives --set javac /usr/java/jdk1.8.0_152/bin/javac
	
查看版本:

	[root@k8s-master1 jdk1.8.0_152]# java -version
	java version "1.8.0_152"
	Java(TM) SE Runtime Environment (build 1.8.0_152-b16)
	Java HotSpot(TM) 64-Bit Server VM (build 25.152-b16, mixed mode)
	[root@k8s-master1 jdk1.8.0_152]# 

添加环境变量:

	echo "export JAVA_HOME=/usr/java/jdk1.8.0_152" >> /etc/profile.d/myprofile.sh
  	echo "export JRE_HOME=/usr/java/jdk1.8.0_152/jre" >> /etc/profile.d/myprofile.sh
  	echo "export PATH=$PATH:/usr/java/jdk1.8.0_152/bin:/usr/java/jdk1.8.0_152/jre/bin" >> /etc/profile.d/myprofile.sh
 	chmod 755 /etc/profile.d/myprofile.sh
  	退出  (to leave the "root" account)
  	source /etc/profile.d/myprofile.sh

查看环境变量:

	echo $JAVA_HOME 如果返回为空，则需要配置环境变量



## 安装es

添加yum源

sudo vi /etc/yum.repos.d/elasticsearch.repo

	[elasticsearch-5.x]
	name=Elasticsearch repository for 5.x packages
	baseurl=https://artifacts.elastic.co/packages/5.x/yum
	gpgcheck=1
	gpgkey=https://artifacts.elastic.co/GPG-KEY-elasticsearch
	enabled=1
	autorefresh=1
	type=rpm-md

注意前面不能有空格或tab,enter键

安装:

	yum install elasticsearch -y

	systemctl daemon-reload

	systemctl enable elasticsearch.service

修改配置:

vi /etc/elasticsearch/jvm.options

	将
	#-Xms2g
	#-Xmx2g
	改为：
	-Xms1g 
	-Xmx1g

使外网可以访问:

vi /etc/elasticsearch/elasticsearch.yml

	network.host: 0.0.0.0

启动:

	systemctl start elasticsearch.service


访问:

	http://ip:9200/

	curl ip:9200


## 同步mysql数据库中的数据至es

安装logstash:

	mkdir -p /opt/logstash && cd /opt/logstash
	wget https://artifacts.elastic.co/downloads/logstash/logstash-5.2.1.tar.gz
	tar -zxvf logstash-5.2.1.tar.gz

	cd /opt/logstash/logstash-5.2.1
	#bin/logstash-plugin install logstash-input-jdbc (这句会报错误,cannot install logstash-core-event-java-5.1.1-java)
	bin/logstash-plugin install --local logstash-input-jdbc (安装本地logstash-input-jdbc)


安装mysql-connector-java-5.1.43.zip:

	cd /opt/logstash 
	wget https://dev.mysql.com/get/Downloads/Connector-J/mysql-connector-java-5.1.43.zip
	unzip mysql-connector-java-5.1.43.zip 

使用:

将jdbc.config上传到/opt/logstash/logstash-5.2.1 然后:

	cd /opt/logstash/logstash-5.2.1
	bin/logstash -f jdbc.config



google浏览器可以使用elasticsearch-head插件，连接该es集群，方便查看数据。


jdbc.config例子:

	
	input {
	  	jdbc {
	    jdbc_driver_library => "/opt/logstash/mysql-connector-java-5.1.43/mysql-connector-java-5.1.43-bin.jar"
	    jdbc_driver_class => "com.mysql.jdbc.Driver"
	    jdbc_connection_string => "jdbc:mysql://{数据库ip地址}:3306/{数据库名称}?useUnicode=true&characterEncoding=utf-8"
	    jdbc_user => "root"
	    jdbc_password => "{数据库登录密码}"
	    schedule => "* * * * *"
	
		clean_run => false 
		record_last_run => true
		last_run_metadata_path => "/var/tmp/last_run_value.last"
		use_column_value => true 
		tracking_column => "id" 
		tracking_column_type => "numeric"
	
		jdbc_paging_enabled => "true"
		jdbc_page_size => 500 
		jdbc_fetch_size => 500 
	    statement => "select * from anote where  cateid=31 AND  id>= :sql_last_value" 
	
	    type => "table_anote"
	  }
	}
	output {
		stdout {
	        codec => json_lines
	    }
	
		if[type] == "table_anote"{
	        elasticsearch {
	        hosts  => "localhost:9200"
	        index => "db_whateat"
	        document_type => "%{type}" # <- use the type from each input
	        document_id => "%{id}" #防止数据重复
	        }
	    }
	
	  
	}

修改替换掉{}中的参数，例如:

	{数据库ip地址}=>127.0.0.1
	{数据库名称}=>testDB