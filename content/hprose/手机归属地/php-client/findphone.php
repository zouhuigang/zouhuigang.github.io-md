<?php
//http://m.xbaod.com/m/ztest/findphone.php
header("Content-type: text/html; charset=utf-8"); 
require_once("init.php");
error_reporting(-1);
ini_set('display_errors', 1);
ini_set('memory_limit','1000M');//给php-fpm设置可用的内存大小
set_time_limit(0);//执行不限时


use Hprose\Client;

try {
		$client = Client::create('http://139.196.48.36:3333/', false);    //qa
		$arrObject=$client->findPhone("18516573852");//0代表，第一个excel,1代表第二行数据开始读取到文件结束,含第二行


}catch (Exception $e) {
		print $e->getMessage();
		exit();
}



//得到数据
$arr=json_decode(json_encode($arrObject),true);
	print_r($arr);

die;
?>
