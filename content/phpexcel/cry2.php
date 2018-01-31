<?php
//https://github.com/PHPOffice/PHPExcel/issues/546
//https://github.com/PHPOffice/PHPExcel/blob/develop/Documentation/markdown/Overview/05-Deleting-a-Workbook.md
////https://github.com/ViktorKITP/schedule/blob/d497859a21acf7fb409150befa74f193f0a01585/server/downLoadXLS.php
//https://github.com/VanyaProger/AICup/blob/a21f63d88d09e8e4823122b0c7955f44cb6efda7/Cms/Modules/Excel/Easy.php
//https://github.com/spotman/kohana-excel-import/blob/a8299e80fb8d69f9ac276b72f543e902272b5a81/classes/Import/Excel/Doc.php
//https://github.com/icatholic/iwebsite2/blob/f185f2eab49c6e321a63cd183f8052e955680b4f/application/modules/admin/controllers/OrderController.php
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

		//$cacheMethod = Settings::cache_in_memory;//缓存
		//Settings::setCache($cacheMethod);

        $inputFileType = 'Xlsx';
		$inputFileName=ZROOT.'/m/pingan/excel/5000.xlsx';
        echo 'Loading file ', pathinfo($inputFileName, PATHINFO_BASENAME), ' using IOFactory with a defined reader type of ', $inputFileType, '<br />';
        /*  Create a new Reader of the type defined in $inputFileType  * */
		$reader = IOFactory::createReader($inputFileType);

		$reader->setReadDataOnly(true);

		$objPHPExcel = $reader->load($inputFileName);
		unset($reader);
		Calculation::getInstance($objPHPExcel)->cyclicFormulaCount = 1;

		$objWorksheet = $objPHPExcel->getActiveSheet();
		$totalRows = $objWorksheet->getHighestRow();

       echo $totalRows.'<br />';;

		for ($row = 2; $row <= $totalRows; $row++) {
				$rowObj = new Row($objWorksheet, $row);
				$cellIterator = $rowObj->getCellIterator();
		  		$cellIterator->setIterateOnlyExistingCells(false);
		  		$arrayValues = array(); 
				foreach ($cellIterator as $cell) {
				  	$arrayValues[] = str_replace('_x000D_', '', $cell->getValue());
				}
				$result[$row] = $arrayValues;
				
		}

		echo convert(memory_get_peak_usage(true)).'<br />';

		//清空缓存
        $objPHPExcel->disconnectWorksheets();
        //删除变量
		unset($reader, $objPHPExcel, $objWorksheet, $totalRows);
		gc_collect_cycles(); 


	

		echo "释放变量后内存占用: ".memory_get_usage()." 字节</br>";
		//var_export($reader);
	
?>

