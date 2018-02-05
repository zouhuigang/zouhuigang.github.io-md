### 安装

	mkdir -p /home/mysql-install && cd /home/mysql-install
	wget http://repo.mysql.com/mysql-community-release-el7-5.noarch.rpm
	sudo rpm -ivh mysql-community-release-el7-5.noarch.rpm



安装这个包后，会获得两个mysql的yum repo源：

	/etc/yum.repos.d/mysql-community.repo
	/etc/yum.repos.d/mysql-community-source.repo


安装mysql

	sudo yum install mysql-server


修改端口:

vi /etc/my.cnf
	