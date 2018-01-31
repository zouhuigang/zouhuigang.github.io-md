### 1.第一次使用jumpserver的几个概念

" 公私钥"认证方式简单的解释:首先在客户端上创建一对公私钥 （公钥文件：~/.ssh/id_rsa.pub； 私钥文件：~/.ssh/id_rsa）。然后把公钥放到服务器上（~/.ssh/authorized_keys）, 自己保留好私钥.在使用ssh登录时,ssh程序会发送私钥去和服务器上的公钥做匹配.如果匹配成功就可以登录了。

UI面板中的：

	创建系统用户：创建系统用户，会在服务器上新建一个用户，然后可以用该用户登录服务器。所以不要用root,防止不小心登录不上服务器了。

	管理用户：是各个服务器的登录用户和密码。


a.首先在资产管理->管理用户列表->创建管理用户,将各个服务器的用户名和密码（私钥）保存起来

b.然后创建资产(也就是服务器等信息),填入Ip的选择管理用户之后，可以测试管理用户是否可以登录成功，返回SUCCESS！再刷新硬件信息,完成。

c.之后再创建系统用户，ssh链接登录服务器。

d.权限管理，将2台服务器的控制权，转给当前登录用户。

e.coco管理端，登录不同服务器。

#### 获取服务器的秘钥

	ssh-keygen -t rsa -P '' 

	#把id_rsa.pub追加到授权的key里面去。	
	#cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys


### 2.docker启动服务

### jumpserver 外网推送

	docker login --username=952750120@qq.com registry.cn-hangzhou.aliyuncs.com

	docker tag jumpserver/jumpserver:v0.4.0-beta1 registry.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/jumpserver-jumpserver:v0.4.0-beta1

	docker push registry.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/jumpserver-jumpserver:v0.4.0-beta1

### coco 内网推送，可以选择走内网，速度将大大提升，并且将不会损耗您的公网流量

	docker login --username=952750120@qq.com registry-internal.cn-hangzhou.aliyuncs.com

	docker tag jumpserver/coco:v0.4.0-beta1 registry-internal.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/jumpserver-coco:v0.4.0-beta1

	docker push registry-internal.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/jumpserver-coco:v0.4.0-beta1


#### luna

	docker tag jumpserver/luna:v0.4.0-beta1  registry-internal.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/jumpserver-luna:v0.4.0-beta1

	docker push registry-internal.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/jumpserver-luna:v0.4.0-beta1



### 启动

	docker-compose up


### 使用

  1. 访问 http://你的主机IP:8080 来访问 Jumpserver

  2. 左侧 应用程序接受 Coco和Luna的注册

  3. 添加 管理用户

  4. 添加 资产

  5. 添加授权规则，授权给admin

  6. ssh -p2222 admin@你的主机IP 测试连接服务器

  7. 访问 http://你的主机IP:5000 访问Luna，点击左侧服务器连接测试


demo使用了开发者模式，并发只能为1

    Jumpserver: 访问 账号: admin 密码: admin

    Luna: 访问 同Jumpserver认证

    Coco: ssh -p 2222 admin@demo.jumpserver.org 密码: admin

### SecureCRT登录Coco

ip:192.168.99.100
用户名：zhg
密码:自定义
端口：2222

![https://raw.githubusercontent.com/zouhuigang/zouhuigang.github.io/master/jumpserver/0.gif](https://raw.githubusercontent.com/zouhuigang/zouhuigang.github.io/master/jumpserver/0.gif)


### 在宿主机上生成公钥

	ssh-keygen -t rsa -P ''  

运行该命令后会生成如下两个文件

id_rsa 和id_rsa.pub


	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC7TltnvOrRulPQnDP7s7K8Dxo1U7F0xTZcBvIJr+rqQztL596/ooTUbPo3HKG7lh7BlZXvBD7CaHiYCcDP8rXowEA40DZ+VoNKIzNqjFwHZWwhw2TlmVciDDljzs1RbDTsg38HsLSdNbMKnmiRHQiWsErooVfCpBY7ET2Fz3NsahQouC3az1bhv3gWB3jAlBPeBpQfTimzQ1taHImiHEOlHOXi+xIdwr1A6SMTHIbOiBNzxip5Zx/dkmRSSF5Tvuq6DmfPvXub6trnzKHkLvFoJLDYZlMh2udrZvm8neUGaYhjdiqp51hFmcbTV6sSlc74C/v28MKDrvUf2rD2rH7b docker@default