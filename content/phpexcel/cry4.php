<?php
//https://github.com/mk-j/PHP_XLSXWriter
//https://stackoverflow.com/questions/25033303/cant-find-a-way-to-fix-phpexcel-memory-leak
//phpexcel插件的内存回收问题，是个很大的问题，所以选择放弃了
header("Content-type: text/html; charset=utf-8"); 

require_once("../include/share.php");
use Box\Spout\Reader\ReaderFactory;
use Box\Spout\Common\Type;
//error_reporting(E_ALL);
error_reporting(-1);
ini_set('display_errors', 1);
ini_set('memory_limit','1500M');//excel的解析，非常的耗费内存，所以这个值要很大分配内存
set_time_limit(0);//执行不限时
ini_set('pcre.backtrack_limit', 100000000);
ini_set('pcre.jit', 0);
gc_enable();
//if (ob_get_length() > 0) { ob_end_clean(); }
echo "初始内存占用: ".memory_get_usage()." 字节</br>";

function convert($size){
    $unit=array('b','kb','mb','gb','tb','pb');
    return @round($size/pow(1024,($i=floor(log($size,1024)))),2).' '.$unit[$i];
}



        $inputFileType = 'Xlsx';
		$inputFileName=ZROOT.'/m/pingan/excel/5000.xlsx';
     

echo "Base Memory: ", memory_get_usage(true), PHP_EOL;


$reader = ReaderFactory::create(Type::XLSX); // for XLSX files
//$reader = ReaderFactory::create(Type::CSV); // for CSV files
//$reader = ReaderFactory::create(Type::ODS); // for ODS files

$reader->open($inputFileName);

foreach ($reader->getSheetIterator() as $sheet) {
    foreach ($sheet->getRowIterator() as $row) {
        // do stuff with the row
    }
}

$reader->close();
unset($reader);
$reader=null;

echo "Went through each row", PHP_EOL,"</br/>";
echo "Final memory usage: ", memory_get_usage(true), PHP_EOL,"</br/>";
gc_collect_cycles(); 	

echo " gc : ", memory_get_usage(true), PHP_EOL,"</br/>";

?>

