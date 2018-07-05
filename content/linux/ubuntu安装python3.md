### 升级python3，不要卸载python2,因为一些其他的系统需要python2

去官网下载最新的安装包：

[https://www.python.org/](https://www.python.org/)

打开:

	https://www.python.org/downloads/release/python-370/
	
选择:

	Gzipped source tarball 下载
	
下载完成，得到:Python-3.7.0.tgz 

	tar xfz Python-3.7.0.tgz 

这里使用xfz命令，而不建议使用-xvzf命令，因为其释放的文件夹需要root权限才可以更改或者删除。

	cd  /Python-3.7.0

	./configure   --prefix=/usr/bin/python3.7 
	sudo make
	sudo make install
	
更改软链接,修改默认Python版本 
	sudo rm -rf /usr/bin/python

	ln -s /usr/bin/python3.7/bin/python3.7  /usr/bin/python

查看版本:

	python -V
	
### Virtualenv 多版本python共用
验证是否安装:

	 virtualenv --version
	 
ubuntu:

	pip install virtualenv
	
创建pyhon3虚拟环境（py3为虚拟环境名）

	virtualenv -p /usr/bin/python3 py3
	
创建python2虚拟环境（py2为虚拟环境名）

	virtualenv -p /usr/bin/python py2
	
启动虚拟环境

	cd [虚拟环境名称]/
	source bin/activate
	python -V
### 退出虚拟环境

	deactivate


[https://blog.csdn.net/menciushometown/article/details/77688728](https://blog.csdn.net/menciushometown/article/details/77688728)

[http://www.pythonforbeginners.com/basics/how-to-use-python-virtualenv/](http://www.pythonforbeginners.com/basics/how-to-use-python-virtualenv/)
