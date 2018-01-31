### 搜索引擎elasticsearch的安装

### windows下载地址

	https://www.elastic.co/downloads/elasticsearch
	https://www.elastic.co/cn/downloads/elasticsearch#ga-release

### 解压缩在安装目录

环境：java环境，（https://www.java.com/zh_CN/）

在cmd命令行进入安装目录，再进入 bin目录，运行elasticsearch.bat命令：
启动成功后。在浏览器中输入:http://localhost:9200/

返回：

	{
	  "name" : "7uzRCoL",
	  "cluster_name" : "elasticsearch",
	  "cluster_uuid" : "HhLG_cBgRG2DwaBzi9uwMA",
	  "version" : {
	    "number" : "5.5.1",
	    "build_hash" : "19c13d0",
	    "build_date" : "2017-07-18T20:44:24.823Z",
	    "build_snapshot" : false,
	    "lucene_version" : "6.6.0"
	  },
	  "tagline" : "You Know, for Search"
	}
页面上json里的 name 是节点的名字，集群名称是 elasticsearch，还有其他版本信息。
想停止的话，ctrl+c

单个 节点 可以作为一个运行中的 Elasticsearch 的实例。 而一个 集群 是一组拥有相同 cluster.name 的节点， 他们能一起工作并共享数据，还提供容错与可伸缩性。(当然，一个单独的节点也可以组成一个集群) 你可以在 elasticsearch.yml 配置文件中 修改 cluster.name ，该文件会在节点启动时加载 (这个重启服务后才会生效)。 



安装插件：

	D:\software\elasticsearch-5.5.1\bin>elasticsearch-plugin install mobz/elasticsearch-head


	elasticsearch-plugin  install elastic/sense


问题：

Q1:Could not reserve enough space for 2097152KB object heap

A1:

因为jvm有内存使用配置。es默认的配置是2g。所以要修改为自己需要的。比如512m
搜索一下，找到修改的位置。

修改的位置，是在D:\software\elasticsearch-5.5.1\config\jvm.options

将：

	-Xms2g
	
	-Xmx2g
	
	修改为
	
	
	-Xms256m
	-Xmx256m

注意，配置文件中说明了initial size和 maximum size必须一致。



参考文档：

[http://www.cnblogs.com/CrazyAnts/p/5849726.html](http://www.cnblogs.com/CrazyAnts/p/5849726.html)

[http://blog.csdn.net/ebw123/article/details/46707559](http://blog.csdn.net/ebw123/article/details/46707559)