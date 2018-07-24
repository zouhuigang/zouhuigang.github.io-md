---
#标题
title: "php"
#描述
description: ""
#创建日期
date: 2018-07-24
#修改日期
lastmod: 2018-07-24
#草稿
draft: false
#关键字
keywords: []
#标签
tags: [post,redis]
#分类
categories: [post,redis]
#作者
author: "邹慧刚"
---
### php连接集群

	$obj_cluster = new RedisCluster(NULL, ['10.81.128.152:7000', '10.81.128.152:7001' ,'10.81.128.152:7002','10.174.113.12:7003', '10.174.113.12:7004','10.174.113.12:7005', 1.5, 1.5]);

	var_dump($obj_cluster);

	$obj_cluster->set('name1', '1111');
	$obj_cluster->set('name2', '2222');
	$obj_cluster->set('name3', '333');
	$name1 = $obj_cluster->get('name1');
	$name2 = $obj_cluster->get('name2');
	$name3 = $obj_cluster->get('name3');
	var_dump($name1, $name2, $name3);


zincrby:

 	$this->redis->zincrby("comment:like",$num,$comment_id); 

加减$num值,如果值不存在，则会创建。


zscore：
	
 	if($this->redis->zscore("comment:like",$comment_id))

判断值是否存在



### list 

		$insertArr=array();
    	$insertArr["siteid"]="1062_1135";
		$insertArr["cname"]="邹慧刚";
		$insertArr["sex"]="男";
		$insertArr["telephone"]="18516573852";
		$insertArr["birthday"]=date("Y-m-d",time());
		$insertArr["adddate"]=date("Y-m-d H:i:s",time());
		$json = json_encode($insertArr);
		$len1=$obj_cluster->rpush("pingan:insert",$json);

		$insertArr=array();
    	$insertArr["siteid"]="1062_11352";
		$insertArr["cname"]="邹慧刚2";
		$insertArr["sex"]="男";
		$insertArr["telephone"]="18516573852";
		$insertArr["birthday"]=date("Y-m-d",time());
		$insertArr["adddate"]=date("Y-m-d H:i:s",time());
		$json = json_encode($insertArr);
		$len2=$obj_cluster->rpush("pingan:insert",$json);

		echo $len1.",".$len2;
		//出队列
		//$list=$obj_cluster->blPop("pingan:insert",10);//在以上实例中，操作会被阻塞，如果指定的列表 key list1 存在数据则会返回第一个元素，否则在等待10秒后会返回 nil 。 
		
		/*brpop和blpop区别就是一个从一个从list末尾去数据.list开始出取数据,blpop和brpop取出数据之后会删除list中的数据,
		如果这时候redis客户端挂了.那个数据就永远丢失了.redis提供了一种备用机制. 
		BRPOPLPUSH 这个命令会先把pop出来的数据.存到指定的list.然后再把传递给redis客户端,多了一层保障
		$data = $obj_cluster->bRPopLPush("pingan:insert","pingan:tmp",30);  //这个一直返回false，用不了。

   		 var_dump($data); //data1  
   		 //some process  
		//$obj_cluster->lrem("pingan:tmp",$data);  */
		//

		//出队列
		//$list=$obj_cluster->blPop("pingan:insert",10);


参考:

[http://www.cnblogs.com/sunshine-H/p/7922285.html](http://www.cnblogs.com/sunshine-H/p/7922285.html)


