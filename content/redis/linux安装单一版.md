### 安装

	yum install -y redis


	vi /etc/redis.conf

配置密码：

	requirepass youpwd

查看配置密码:

	 cat /etc/redis.conf |grep requirepass

启动redis:

	systemctl start redis