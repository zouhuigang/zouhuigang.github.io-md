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
//ini_set('memory_limit', '1G');
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
     




		function test ($reader, &$data) {
			
		$inputFileName=ZROOT.'/m/pingan/excel/5000.xlsx';
    $reader->open($inputFileName);
    $sheets = $reader->getSheetIterator();

    foreach ($sheets as $sheet) {
        $data[] = [
            'title' => $sheet->getName(),
            'data' => iterator_to_array($sheet->getRowIterator(), false),
        ];

		var_dump('Read '.round(memory_get_usage(true)/1048576,2));
		echo "</br>";
    }

    $reader->close();
    $data = null;

	var_dump('After close '.round(memory_get_usage(true)/1048576,2));
	echo "</br>";
}

var_dump('Start '.round(memory_get_usage(true)/1048576,2));
echo "</br>";

$data = null;
$reader = ReaderFactory::create(Type::XLSX);
$reader->setShouldFormatDates(false)->setShouldPreserveEmptyRows(false);

test($reader, $data);

$reader = null;
usleep(2000000);   
$gcnum1=gc_collect_cycles();
echo $gcnum1."</br>";
var_dump('After reader destroy 1:'.round(memory_get_usage(true)/1048576,2));
echo "</br>";

$data = null;
$reader = ReaderFactory::create(Type::XLSX);
$reader->setShouldFormatDates(false)->setShouldPreserveEmptyRows(false);

test($reader, $data);

$reader = null;
usleep(2000000);   
$gcnum2=gc_collect_cycles();
echo $gcnum2."</br>";
var_dump('After reader destroy 2:'.round(memory_get_usage(true)/1048576,2));
echo "</br>";
usleep(2000000);   
die;
?>

