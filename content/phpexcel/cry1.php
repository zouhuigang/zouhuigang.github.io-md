<?php
//https://github.com/PHPOffice/PHPExcel/issues/546
//https://github.com/PHPOffice/PHPExcel/blob/develop/Documentation/markdown/Overview/05-Deleting-a-Workbook.md
header("Content-type: text/html; charset=utf-8"); 
require_once("../include/share.php");
use PhpOffice\PhpSpreadsheet\IOFactory;
use PhpOffice\PhpSpreadsheet\Reader\IReadFilter;
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
        /**  Define a Read Filter class implementing IReadFilter  */
        class chunkReadFilter implements IReadFilter
        {
            private $_startRow = 0;
            private $_endRow = 0;
            /**
             * Set the list of rows that we want to read.
             *
             * @param mixed $startRow
             * @param mixed $chunkSize
             */
            public function setRows($startRow, $chunkSize)
            {
                $this->_startRow = $startRow;
                $this->_endRow = $startRow + $chunkSize;
            }
            public function readCell($column, $row, $worksheetName = '')
            {
                //  Only read the heading row, and the rows that are configured in $this->_startRow and $this->_endRow
                if (($row == 1) || ($row >= $this->_startRow && $row < $this->_endRow)) {
                    return true;
                }
                return false;
            }
        }
        echo 'Loading file ', pathinfo($inputFileName, PATHINFO_BASENAME), ' using IOFactory with a defined reader type of ', $inputFileType, '<br />';
        /*  Create a new Reader of the type defined in $inputFileType  * */
		$reader = IOFactory::createReader($inputFileType);
        /*  Define how many rows we want to read for each "chunk"  * */
		$chunkSize = 1000;
		$totalRows = 10000;
        /*  Create a new Instance of our Read Filter  * */
        $chunkFilter = new chunkReadFilter();
        /*  Tell the Reader that we want to use the Read Filter that we've Instantiated  * */
		$reader->setReadFilter($chunkFilter);
		$reader->setReadDataOnly(true);
        /*  Loop to read our worksheet in "chunk size" blocks  * */
		for ($startRow = 2; $startRow <= $totalRows; $startRow += $chunkSize) {
            echo 'Loading WorkSheet using configurable filter for headings row 1 and for rows '.$startRow.'---'.($startRow + $chunkSize - 1).'<br />';
			$_reader=$reader;
			$_chunkFilter=$chunkFilter;
			$_chunkFilter->setRows($startRow, $chunkSize);
			$_reader->setReadDataOnly(true);
            $objPHPExcel = $_reader->load($inputFileName);
			$sheetData = $objPHPExcel->getSheet(0)->toArray(null,true,null);
   			 if (!empty($sheetData) && $startRow < $totalRows) {
				 //   call_user_func($rowCallback, $sheetData);
    		}
			echo convert(memory_get_peak_usage(true)).'<br />';

			//$objPHPExcel->getActiveSheet()->disconnectCells();
			$objPHPExcel->disconnectWorksheets();
			unset($objPHPExcel,$sheetData,$objectReader,$_chunkFilter);
			
		}

		var_dump($chunkFilter);
		
		unset($reader,$chunkFilter);

		echo "释放变量后内存占用: ".memory_get_usage()." 字节</br>";
		//var_export($reader);
		for($i=0;$i<10;$i++){
			gc_collect_cycles(); // 显式调用GC回收循环引用的对象
			echo "垃圾回收后内存占用: ".memory_get_usage()." 字节</br>";
		}

        ?>

