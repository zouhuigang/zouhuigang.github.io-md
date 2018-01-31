#!/bin/bash
#mysqldump -P 4306 -h 127.0.0.1 -u 'root' -p'yy2017622' -R yy_sys > backup.sql
#定时任务mysqldump -P 4306 -h 127.0.0.1 -u 'root' -p'yy2017622' -R yy_sys | gzip > /mnt2/mysql-backup/yysys-WWW-`date+\%Y-\%m-\%d_\%H.\%M.\%S`.sql.gz
nowtime=`date +%Y-%m-%d_%H.%M.%S`
#echo /mnt2/mysql-backup/yysys-WWW-${nowtime}.sql.gz
mysqldump -P 4306 -h 127.0.0.1 -u 'root' -p'yy2017622' -R yy_sys | gzip > /mnt2/mysql-backup/yysys-WWW-${nowtime}.sql.gz
