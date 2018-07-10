### 将服务打包成ubuntu下可直接安装的程序

	sudo apt-get install dh-make
	

	
配置文件:

sudo vi DEBIAN/control:

	Package: ZphoneServer
	Version: 1.0.0
	Section: gnome
	Priority: gnome
	Architecture: amd64
	Depends:
	Maintainer: zouhuigang888@gmail.com
	Homepage: https://github.com/zouhuigang
	Description: zouhuigang zphone server



说明:

	Package:	#包名
	Version:	#版本号
	Section:	gnome
	Priority:	extra #优先级
	Maintainer:	name<mailaddress> #维护者
	Homepage:	http://... #主页地址
	Architecture:	amd64   #这里添加 i386   amd64等参数
	Depends: #如果存在依赖的库在这里填写
	Description:	#描述信息


	
	sudo apt-get install ./package.deb