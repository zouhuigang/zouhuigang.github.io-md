#!/bin/bash
#日志保存路径
log=/home/wwwlogs/php-fpm.log
echo `date "+%F %H:%M:%S"` >> $log

for PID in `ps aux|grep php-fpm|grep -v grep|awk '{if($4>=1)print $2}'`
do
	kill -9 $PID
	echo $PID >> $log

done

