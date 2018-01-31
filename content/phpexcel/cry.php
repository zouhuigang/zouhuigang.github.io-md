<?php
header("Content-type: text/html; charset=utf-8"); 
echo phpinfo();
gc_enable();
function convert($size){
$unit=array('b','kb','mb','gb','tb','pb');
return @round($size/pow(1024,($i=floor(log($size,1024)))),2).' '.$unit[$i];
}

echo "释放变量后内存占用: ".convert(memory_get_usage(true))."</br>";
$gcnum=gc_collect_cycles(); // 显式调用GC回收循环引用的对象
echo "已回收对象:".$gcnum."个,垃圾回收后内存占用: ".convert(memory_get_usage(true))."</br>";


?>
