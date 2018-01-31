<?php
header("Content-type: text/html; charset=utf-8"); 
require_once("../include/share.php");

use KHerGe\Excel\Workbook;

echo "初始内存占用: ".memory_get_usage()." 字节</br>";

$pathToFile=ZROOT.'/m/pingan/excel/900k.xlsx';
// Open a workbook file.
$workbook = new Workbook($pathToFile);

// Iterate through all worksheets.
$rowCount = 0;
foreach ($workbook->iterateWorksheets() as $index => $worksheet) {
    // Iterate through all rows in a worksheet.
    foreach ($worksheet->iterateRows() as $row => $values) {
        $rowCount++;
        
        // foreach ($values as $column => $value) {
            
        // }
    }
}

unset($workbook);
echo "总行数：" + $rowCount;

echo "结束后内存占用: ".memory_get_usage()." 字节</br>";

?>
