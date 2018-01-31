###搭建vpn服务器，方便翻墙或访问公司内网

centos7:

	wget https://raw.githubusercontent.com/DanylZhang/VPS/master/CentOS7-pptp-host1plus.sh
	chmod +x ./CentOS7-pptp-host1plus.sh
	./CentOS7-pptp-host1plus.sh -u zouhuigang -p 123456


运行成功：

You can now connect to your VPN via your external IP 101.81.73.142
Username: zouhuigang
Password: c72-rrb-yDH
[root@localhost ~]# 

[http://blog.csdn.net/zdy1127/article/details/50664736](http://blog.csdn.net/zdy1127/article/details/50664736)

[http://www.win7china.com/html/15202.html](http://www.win7china.com/html/15202.html)

[https://www.youtube.com/watch?v=RpmCqxtZjCY](https://www.youtube.com/watch?v=RpmCqxtZjCY)

[http://www.ithtw.com/9246.html](http://www.ithtw.com/9246.html)

[https://hub.docker.com/r/mobtitude/vpn-pptp/](https://hub.docker.com/r/mobtitude/vpn-pptp/)

[http://www.kwor.cn/2017/04/pptp-vpn.html](http://www.kwor.cn/2017/04/pptp-vpn.html)

[https://github.com/hwdsl2/docker-ipsec-vpn-server/blob/master/README-zh.md](https://github.com/hwdsl2/docker-ipsec-vpn-server/blob/master/README-zh.md)


	FROM ubuntu:16.04 MAINTAINER Przemek Szalko <przemek@mobtitude.com> ENV DEBIAN_FRONTEND noninteractive RUN apt-get update && apt-get install -y pptpd iptables
	COPY ./etc/pptpd.conf /etc/pptpd.conf
	COPY ./etc/ppp/pptpd-options /etc/ppp/pptpd-options
	COPY entrypoint.sh /entrypoint.sh
	RUN chmod 0700 /entrypoint.sh
	ENTRYPOINT ["/entrypoint.sh"]
	CMD ["pptpd", "--fg"]



### 第二种安装方式

	docker tag hwdsl2/ipsec-vpn-server registry.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/ipsec-vpn-server:latest
	
	docker tag  registry.cn-hangzhou.aliyuncs.com/zhg_docker_ali_r/ipsec-vpn-server:latest hwdsl2/ipsec-vpn-server

vpn.env文件说明：

	VPN_IPSEC_PSK=这里填一个随机字符串,随便输吧,32位以内
	VPN_USER=这里输入vpn的登录名
	VPN_PASSWORD=这里输入vpn的登录密码


首先在 Docker 主机上加载 IPsec NETKEY 内核模块：

	sudo modprobe af_key

将vpn.env上传到/root文件夹下面

docker run \
    --name ipsec-vpn-server \
    --env-file vpn.env \
    --restart=always \
    -p 500:500/udp \
    -p 4500:4500/udp \
    -v /lib/modules:/lib/modules:ro \
    -d --privileged \
    hwdsl2/ipsec-vpn-server

查看信息：

	docker logs ipsec-vpn-server


查看服务器状态

如需查看你的 IPsec VPN 服务器状态，可以在容器中运行 ipsec status 命令：

	docker exec -it ipsec-vpn-server ipsec status

或者查看当前已建立的 VPN 连接：

	docker exec -it ipsec-vpn-server ipsec whack --trafficstatus





Android

    启动 设置 应用程序。
    在 无线和网络 部分单击 更多...。
    单击 VPN。
    单击 添加VPN配置文件 或窗口右上角的 +。
    在 名称 字段中输入任意内容。
    在 类型 下拉菜单选择 L2TP/IPSec PSK。
    在 服务器地址 字段中输入你的 VPN 服务器 IP。
    在 IPSec 预共享密钥 字段中输入你的 VPN IPsec PSK。
    单击 保存。
    单击新的VPN连接。
    在 用户名 字段中输入你的 VPN 用户名。
    在 密码 字段中输入你的 VPN 密码。
    选中 保存帐户信息 复选框。
    单击 连接。

VPN 连接成功后，会在通知栏显示图标。最后你可以到 这里 检测你的 IP 地址，应该显示为你的 VPN 服务器 IP。




[----win7连接vpn，点击--------](win7连接.md)



我的问题：

我在本地win7电脑的vm虚拟机上搭建了一个vpn，为什么手机连不上去？
已经把win7的防火墙关了，用的极路由的没有关闭防火墙的选项，请问要怎样才能成功呢？是要开什么端口么？

作者答复：

你好！首先你的vm虚拟机必须使用bridge networking，不能用 NAT。然后需要在你的路由器上设置端口转发 UDP 500 和 UDP 4500 到你的虚拟机的 IP 地址。建议设置路由器的 DHCP 的 MAC 地址绑定，以保证虚拟机的 IP 不变。最后你的 ISP 必须分配一个公有 IP （不是 carrier-grade NAT)。






[https://gist.github.com/dferg/2278196d736d5b367682](https://gist.github.com/dferg/2278196d736d5b367682)

[http://webcache.googleusercontent.com/search?q=cache:http://liangshuang.name/2017/06/19/linode-ipsec-vpn/](http://webcache.googleusercontent.com/search?q=cache:http://liangshuang.name/2017/06/19/linode-ipsec-vpn/)

[http://blog.csdn.net/joanna_yan/article/details/50084945](http://blog.csdn.net/joanna_yan/article/details/50084945)

[http://www.winwin7.com/JC/Win7JC-1485.html](http://www.winwin7.com/JC/Win7JC-1485.html)