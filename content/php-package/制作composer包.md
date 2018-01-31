### 开发 composer 包有以下几个步骤：

    初始化 composer.json 文件
    定义命名空间及包名
    实现包需要实现的功能
    提交到 GitHub
    在 Packagist 注册包


### 1.初始化 composer.json 文件


安装好 composer 后即可在本地运行 composer init 通过交互式命令行设置 composer.json

	cmd 进入D:\www\php-package\elasticsearch

	# composer init 

	Name:zouhuigang/php-package/elasticsearch  #这里填写<包提供者>/<包名>的信息
	Description："es search package"	#包的描述
	Author [zouhuiang <952750120@qq.com>]: #作者信息
	Minimum Stability []: dev	#最低稳定版本
	Package Type (e.g. library, project, metapackage, composer-plugin) []: library #包类型
	License []: MIT #授权协议

	Define your dependencies.

	Would you like to define your dependencies (require) interactively [yes]? no
	Would you like to define your dev dependencies (require-dev) interactively [yes]? no

	#安装当前包所需的依赖。只有所有依赖被安装当前包才会被安装。
	Do you confirm generation [yes]? yes


进过上面的步骤，即可生成一个composer.json


### 2.定义命名空间及包名


### 3.实现包需要实现的功能

### 4.提交到 GitHub


### 5.在 Packagist 注册包

手动提交：

1.打开 [https://packagist.org/packages/submit](https://packagist.org/packages/submit)


输入 https://github.com/zouhuigang/php-esc

提交即可


### 6.下载更新第三方包

cmd进入一个php项目，然后在项目根目录运行:

	#composer require zouhuigang/php-esc
	#composer update zouhuigang/php-esc
	composer require "zouhuigang/php-esc:dev-master" #下载dev版本,所以还是指定Minimum Stability:stable吧


### 7.引用文件，自动加载Php文件

在项目根目录编写加载文件ComposerAutoload.php：


	<?php

	//第 1 行代码会为 vendor 目录下的所有包添加 autoload
	
	$loader = require_once __DIR__.'/vendor/autoload.php';
	//print_r($loader);die;
	 
	// reg my bundles 第 2 行代码则注册了我自己的包(bundles)目录，我只需要在项目目录创建一个 bundles 文件夹
	$loader->add('Bundles', __DIR__ . '/vendor');

	?>

然后在index.php中添加下列代码：

	require_once __DIR__ . '/ComposerAutoload.php';


### 8.测试运行



参考文档：

[http://blog.zhengshuiguang.com/php/packagist.html](http://blog.zhengshuiguang.com/php/packagist.html)


	



