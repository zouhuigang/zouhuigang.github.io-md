#!/bin/bash
echo "git pull mysql backup data...."
#scp root@10.174.155.169:/mnt2/mysql-backup/yysys-WWW-2017-08-16_21.44.03.sql.gz /mnt2/mysql-backup
echo "git pull success...."

echo "restore mysql data...."
gunzip yysys-WWW-2017-08-16_21.44.03.sql.gz
mysql -h 127.0.0.1 -P 3308 -u root -p yy_sys<yysys-WWW-2017-08-16_21.44.03.sql
