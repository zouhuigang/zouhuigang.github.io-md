
Q1:

使用docker-compose报错:


	pkg_resources.DistributionNotFound: The 'docker-compose==1.7.1' distribution was not found and is required by the application


A1:

	先启动docker mysql

	docker run  -d  -e MYSQL_ROOT_PASSWORD=TYwy2016720 -v /mnt/gitcoding/docker-lnmp-redis/mysql/mysql.cnf:/etc/mysql/conf.d/mysql.cnf:ro -v  /mnt/gitcoding/docker-lnmp-redis/site/mysqldata:/var/lib/mysql -p 3306:3306  registry.aliyuncs.com/zhg_docker_ali_r/mysql 