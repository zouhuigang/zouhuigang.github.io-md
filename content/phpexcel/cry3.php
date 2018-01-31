<?php
//https://github.com/mk-j/PHP_XLSXWriter
//https://stackoverflow.com/questions/25033303/cant-find-a-way-to-fix-phpexcel-memory-leak
//phpexcel插件的内存回收问题，是个很大的问题，所以选择放弃了
header("Content-type: text/html; charset=utf-8"); 

require_once("../include/share.php");
use PhpOffice\PhpSpreadsheet\IOFactory;
use PhpOffice\PhpSpreadsheet\Reader\IReadFilter;
use PhpOffice\PhpSpreadsheet\Worksheet\Row;
use PhpOffice\PhpSpreadsheet\Settings;
use PhpOffice\PhpSpreadsheet\Calculation;
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

$inputFileType = IOFactory::identify($inputFileName);
$objReader = IOFactory::createReader($inputFileType);
$objReader->setReadDataOnly(true);

$excelReader = $objReader->load($inputFileName);

echo "Reader Initialised/File Loaded", PHP_EOL;
echo memory_get_usage(true), PHP_EOL;

foreach ($excelReader->setActiveSheetIndex(0)->getRowIterator() as $row) {
    $cellIterator = $row->getCellIterator();
    $cellIterator->setIterateOnlyExistingCells(false);

    foreach ($cellIterator as $cell) {
        $columnIndex = $cell->getColumn();
        $cellValue = $cell->getCalculatedValue();
    }
    if (($row->getRowIndex() % 256) == 0) {
        echo "Row ".$row->getRowIndex(), ' memory usage: ', memory_get_usage(true), PHP_EOL;
    }
}

echo "Went through each row", PHP_EOL;
echo "Final memory usage: ", memory_get_usage(true), PHP_EOL;
	

?>

