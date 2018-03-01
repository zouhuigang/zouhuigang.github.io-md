<?php
//http://m.xbaod.com/m/ztest/findphone.php
header("Content-type: text/html; charset=utf-8"); 
//error_reporting(-1);ini_set('display_errors', 1);
require_once __DIR__ . '/ComposerAutoload.php';
error_reporting(-1); //error_reporting(0);
ini_set('display_errors', 1);
ini_set('memory_limit','1000M');//给php-fpm设置可用的内存大小
set_time_limit(0);//执行不限时


use Hprose\Client;

$client = Client::create('http://192.168.122.150:9999/', false);    //test服务器
$inputFileName='/1.xlsx';//这个文件需要跟ZexcelServer同在一台服务器上，不然找不到文件会报错,使用的是绝对路径。
//$arr=$client->readExcel($inputFileName,0,1);
$arr=$client->readExcel($inputFileName,0,1,'73b0f08a945bb480d0096c70142bbe84');


//excel中的日期，使用TEXT文本格式或yyyy-mm-dd hh:mm:ss，防止乱掉
foreach($arr as $k=>$v){ 
	//日期会返回2017\-09\-21\ 14:03:46,因此，我们需要去掉\
	//
	//echo date('Y-m-d H:i:s',strtotime(stripslashes($v[0])))."-".stripslashes($v[0]).'</br>';
	print_r($v);
}

echo "sss";die;

?>
