<?php
/********************************************/
/*		时间屏蔽规则操作 timerule.php					*/
/********************************************/
header("Content-type: text/html; charset=utf-8"); 
require_once("../include/share.php");

use PhpOffice\PhpSpreadsheet\Spreadsheet;
use PhpOffice\PhpSpreadsheet\Writer\Xlsx;
use PhpOffice\PhpSpreadsheet\IOFactory;
use PhpOffice\PhpSpreadsheet\Style\NumberFormat;

function uc2html($str) { 
  
  $ret= iconv("utf-8","gbk",$str); 

  return $ret; 
} 
  
function import_pa_del($DB){

}

function import_pa_execl($DB){

	error_reporting(-1);
	ini_set('display_errors', 1);
	ini_set('memory_limit','500M');//分配内存
	
	//echo getcwd(); 
	$siteid=$_POST['siteid'];
	if($_FILES ['execl'] ['tmp_name']){

		$tmp_file = $_FILES ['execl'] ['tmp_name'];
		$file_types = explode ( ".", $_FILES ['execl'] ['name'] );
	    $file_type = $file_types [count ( $file_types ) - 1];
		 /*设置上传路径*/
		 $savePath = ZROOT.'/m/pingan/excel/';
		
	    /*以时间来命名上传的文件*/
	     $str = date("Ymdhis",time());
	     $file_name = $str . "." . $file_type;
		
		 /*是否上传成功*/
		 if(!copy($tmp_file,$savePath.$file_name)){
			 $this->error( '上传失败');
		 }
		
		 set_time_limit(0);//执行不限时
		 $pathToFile=$savePath.$file_name;

		
		//  $identify = IOFactory::identify($pathToFile);
		//  $reader = IOFactory::createReader($identify);
		//  $spreadsheet = $reader->load($pathToFile);

		 $reader = \PhpOffice\PhpSpreadsheet\IOFactory::createReader("Xlsx");
		 $spreadsheet = $reader->load($pathToFile);

		 $m=memory_get_usage(); //获取当前占用内存	
		 echo "垃圾回收前内存占用: ".$m." 字节</br>";
		 $worksheet = $spreadsheet->getSheet(0);//读取第一个sheet
		 $maxRowNumber = $worksheet->getHighestDataRow();// 取得总行数

		
		if($maxRowNumber>=2){
			
			$failArr=array();//失败记录
			$x=0;
			for($j=2;$j<=$maxRowNumber;$j++)
			{
				//读取值
				$a1 = $spreadsheet->getActiveSheet()->getCell("A".$j)->getValue();//获取a列的值：注册时间
				$b1 = $spreadsheet->getActiveSheet()->getCell("B".$j)->getValue();//获取b列的值：姓名
				$b1=strFilter($b1);
				$c1 = $spreadsheet->getActiveSheet()->getCell("C".$j)->getValue();//获取c列的值：性别
				$d1 = $spreadsheet->getActiveSheet()->getCell("D".$j)->getValue();//获取d列的值：生日
				$e1 = $spreadsheet->getActiveSheet()->getCell("E".$j)->getValue();//获取e列的值：电话
				$f1 = $spreadsheet->getActiveSheet()->getCell("F".$j)->getValue();//获取f列的值：省份
				$g1 = $spreadsheet->getActiveSheet()->getCell("G".$j)->getValue();//获取g列的值：城市
				$h1 = $spreadsheet->getActiveSheet()->getCell("H".$j)->getValue();//获取h列的值：城市cityName
				$i1 = $spreadsheet->getActiveSheet()->getCell("I".$j)->getValue();//获取i列的值：mediasource
				$j1 = $spreadsheet->getActiveSheet()->getCell("J".$j)->getValue();//获取j列的值：画像personas
				$k1 = $spreadsheet->getActiveSheet()->getCell("K".$j)->getValue();//获取k列的值：车主Y或空
				$l1 = $spreadsheet->getActiveSheet()->getCell("L".$j)->getValue();//获取l列的值：保险意愿度insuranceAwareness
				$m1 = $spreadsheet->getActiveSheet()->getCell("M".$j)->getValue();//获取m列的值：问题1，高净值、信用卡
				$n1 = $spreadsheet->getActiveSheet()->getCell("N".$j)->getValue();//获取n列的值：答案1
				$o1 = $spreadsheet->getActiveSheet()->getCell("O".$j)->getValue();//获取o列的值：问题2
				$p1 = $spreadsheet->getActiveSheet()->getCell("P".$j)->getValue();//获取P列的值：答案2，贷款
				$q1 = $spreadsheet->getActiveSheet()->getCell("Q".$j)->getValue();//获取Q列的值：comefrom，高净值、信用卡
				$r1 = $spreadsheet->getActiveSheet()->getCell("R".$j)->getValue();//获取R列的值：comefrom2,贷款

				$d1 = NumberFormat::toFormattedString($d1,"yyyy-m-d");
				$a1 = NumberFormat::toFormattedString($a1,"yyyy-m-d H:i:s"); 

			}
		}
		unset($reader);//释放内存,防止每次用完，服务器内存就增加，导致服务器挂掉
		unset($identify);
		unset($spreadsheet);
		$spreadsheet=null;
	 	$reader= null; 
		$identify=null;
		
		$mm=memory_get_usage(); //unset()后再查看当前占用内存
		$f=$m-$mm;
		gc_collect_cycles(); // 显式调用GC回收循环引用的对象
		echo "垃圾回收后内存占用: ".memory_get_usage()." 字节</br>";
		//echo "<script>alert('导入完成,内存是否释放".$f."');location.href='list_pa.php?status=0';</script>";
		die;
	}
	else{

		echo $siteid."没有读取到execl";die;
	}
	
}




function showForm($DB)
{

?>
<link href="../calendar/JSCalendar.css" rel=stylesheet type=text/css>
<script src="../calendar/JSCalendar.js"></SCRIPT>
<script>
function import_pa(){

	
	document.checkoutform.submit();
}


</script>

<table width="98%" border="0" align="center" cellpadding="50" cellspacing="1" bgcolor="#CCCCCC">
  <tr>
    <td bgcolor="#FFFFFF"><table width="450" border="0" align="center" cellpadding="0" cellspacing="1" bgcolor="#CCCCCC">
      <form action="" method="post" name="checkoutform" onsubmit="return import_pa(this)" enctype="multipart/form-data">
	  <input type="hidden" name="func" value="S">
      <tr>
        <td height="40" colspan="2" align="center" bgcolor="#F0F0F0"><b>平安寿名单上载</b></td>
      </tr>
	
<tr>
<td width="100" height="25"  bgcolor="#F0F0F0" align="right">广告位名称：</td>
  <td bgcolor="#FFFFFF">
<select name="siteid">	 
<?php
$_sql="select * from view_adlist where proid=17 and media_ifuse=1 and siteinfo_ifuse=0 and mediaid in(58,59) order by mediaid asc";
/*
$_sql = "SELECT a.*,b.media_name,c.project FROM ".$Tables['siteinfo']." as a ";
	$_sql.=" left join ".$Tables['media']." as b on a.mediaid=b.mediaid ";
	$_sql.=" left join ".$Tables['project']." as c on a.proid=c.proid ";
	if(!empty($_media)){
		$_sql.=" where a.mediaid=".$_mediaid."";
	}
	$_sql.=" order by a.mediaid asc,a.id desc";
	*/
	$_res = $DB->sql_query($_sql); 
	$mediaid=0;
		while($_row = $DB->sql_fetchrow($_res))
		{
			$mid=$_row['mediaid'];
			if(($mediaid!=$mid)&&($mediaid>0)){
				echo "<option></option>";
			}
			$mediaid=$mid;
			echo "<option value='{$_row['siteid']}'>".$_row['media_name']."-".$_row['sitename']."</option>";
		}

?>
				
	   </select>
</td>
</tr>
<tr>
<td width="100" height="25"  bgcolor="#F0F0F0" align="right">上传文件：</td>
  <td bgcolor="#FFFFFF"> <input type="file" id="file" name="execl" accept=".xlsx,.xls" /></td>


</tr>
      <tr>
        <td height="40" colspan="2" align="center" bgcolor="#F0F0F0"><input type="button" name="Submit" value="导入数据"  onclick="import_pa();" />
          &nbsp;&nbsp;&nbsp;&nbsp;<a href="temp.xlsx">导入模板</a></td>
        </tr>
		</form>
    </table></td>
  </tr>
</table>
<?php
}


 

global $_POST,$_GET;

	
if( passCheck("userid",true) ){

	$func = $_POST['func'] ? $_POST['func'] : $_GET['func'];
	if($func=="")$func="M";
	

		switch($func){
			case "D":
				import_pa_del($DB);
				break;
			case "S":
				import_pa_execl($DB);
				break;
			default:
				headhtml("平安寿名单上载","..");
			//	menu();
				showForm($DB);
				foothtml();
				break;
		}

	
	

}



/**
 *@目的： 过滤非法字符
 *@参数： $str		=>	源字符

 *@返回： 
*/
function strFilter($str){
    $str = str_replace('`', '', $str);
	$str = str_replace('’', '', $str);
    $str = str_replace('·', '', $str);
    $str = str_replace('~', '', $str);
    $str = str_replace('!', '', $str);
    $str = str_replace('！', '', $str);
    $str = str_replace('@', '', $str);
    $str = str_replace('#', '', $str);
    $str = str_replace('$', '', $str);
    $str = str_replace('￥', '', $str);
    $str = str_replace('%', '', $str);
    $str = str_replace('^', '', $str);
    $str = str_replace('……', '', $str);
    $str = str_replace('&', '', $str);
    $str = str_replace('*', '', $str);
    $str = str_replace('(', '', $str);
    $str = str_replace(')', '', $str);
    $str = str_replace('（', '', $str);
    $str = str_replace('）', '', $str);
    $str = str_replace('-', '', $str);
    $str = str_replace('_', '', $str);
    $str = str_replace('——', '', $str);
    $str = str_replace('+', '', $str);
    $str = str_replace('=', '', $str);
    $str = str_replace('|', '', $str);
    $str = str_replace('\\', '', $str);
    $str = str_replace('[', '', $str);
    $str = str_replace(']', '', $str);
    $str = str_replace('【', '', $str);
    $str = str_replace('】', '', $str);
    $str = str_replace('{', '', $str);
    $str = str_replace('}', '', $str);
    $str = str_replace(';', '', $str);
    $str = str_replace('；', '', $str);
    $str = str_replace(':', '', $str);
    $str = str_replace('：', '', $str);
    $str = str_replace('\'', '', $str);
    $str = str_replace('"', '', $str);
    $str = str_replace('“', '', $str);
    $str = str_replace('”', '', $str);
    $str = str_replace(',', '', $str);
    $str = str_replace('，', '', $str);
    $str = str_replace('<', '', $str);
    $str = str_replace('>', '', $str);
    $str = str_replace('《', '', $str);
    $str = str_replace('》', '', $str);
    $str = str_replace('.', '', $str);
    $str = str_replace('。', '', $str);
    $str = str_replace('/', '', $str);
    $str = str_replace('、', '', $str);
    $str = str_replace('?', '', $str);
    $str = str_replace('？', '', $str);
    return trim($str);
}

?>

