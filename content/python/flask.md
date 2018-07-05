### 安装flask

	pip install flask

pip install -r requirements.txt

运行:

	python manage.py runserver
	
	python manage.py -c development  # 开发环境运行
	python manage.py -c testing      # 测试环境运行


问题汇总:

Q:
	安装gevent
	  error: command 'x86_64-linux-gnu-gcc' failed with exit status 1

A:
	(py3) zhg@zhg-ThinkPad-E450c:~/workspaces/py-test$ python -V
	Python 3.6.5


	For Python 2.x use:

	   $ sudo apt-get install python-dev
	For Python 2.7 use:

	   $ sudo apt-get install libffi-dev
	For Python 3.x use:

	   $ sudo apt-get install python3-dev
	For Python 3.4 use:

	   $ sudo apt-get install python3.4-dev
	For Python 3.5 use:

	   $ sudo apt-get install python3.5-dev
	For Python 3.6 use:

	   $ sudo apt-get install python3.6-dev
	
