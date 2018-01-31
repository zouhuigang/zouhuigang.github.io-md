### 使用


	git clone https://github.com/letsencrypt/letsencrypt

	$ cd letsencrypt

	$ ./letsencrypt-auto --help



### Let's Encrypt 验证方式

Let's Encrypt 使用两种方式对申请的域名进行验证：

1、 手动验证 按照提示在申请证书的服务器上使用一个指定的URL提供一个指定的文件内容来进行验证，进行手动验证的服务器IP地址会被 Let's Encrypt 服务端记录在案。

2、 自动验证 在 目标服务器 （指域名解析对应的IP地址的服务器，下同）上运行客户端，并启动一个 80 或 443 端口进行自动验证。包括独立模式和其他web sever验证模式，在 Plugins 中详细解释


#### 手动验证

	./letsencrypt-auto certonly --manual -d www.anooc.com


选择 Yes 继续后，便会提示创建一个指定内容的 URL 用来验证对域名及服务器的所有权，注意这个URL仍然需要部署在 目标服务器 上：

	Create a file containing just this data:

	G3UdelazW3uCQSk2JIiBX-X3_nk4uFLCN57Ih6--xUs.-QfnFZBzb2toRS_u4TSdmjy38XELnboJMzn2VBx5Ydo
	
	And make it available on your web server at this URL:
	
	http://www.anooc.com/.well-known/acme-challenge/G3UdelazW3uCQSk2JIiBX-X3_nk4uFLCN57Ih6--xUs

根据提示，创建一个文件目录：

	mkdir -p 网站根目录/.well-known/acme-challenge

创建文件：

	echo 'G3UdelazW3uCQSk2JIiBX-X3_nk4uFLCN57Ih6--xUs.-QfnFZBzb2toRS_u4TSdmjy38XELnboJMzn2VBx5Ydo'>>G3UdelazW3uCQSk2JIiBX-X3_nk4uFLCN57Ih6--xUs



问题：

修改：/etc/nginx/conf.d/default.conf

	由于nginx可能不支持.well-konw的访问，所有在配置文件中添加

	location ^~ /.well-known/acme-challenge/ {
	   default_type "text/plain";
	   root     /usr/share/nginx/html;
	}
	
	location = /.well-known/acme-challenge/ {
	   return 404;
	}

或者：

	location ~ /.well-known/acme-challenge {
        allow all;
    }

可以看到，上面的root，我们指定根目录为：/usr/share/nginx/html，因为我的应用是通过NodeJS的ExpressJS写的，如果修改源代码的话，比较麻烦。因此我就让检验的链接指向了nginx默认的文件夹下。


	location = /.well-known/acme-challenge/ {
	return 404;
	}
	要去掉，不然CertBot会无法验证



输出：


	[root@k8s-master letsencrypt]# ./letsencrypt-auto certonly --manual -d www.xbaod.com
	Saving debug log to /var/log/letsencrypt/letsencrypt.log
	Obtaining a new certificate
	Performing the following challenges:
	http-01 challenge for www.xbaod.com
	
	-------------------------------------------------------------------------------
	NOTE: The IP of this machine will be publicly logged as having requested this
	certificate. If you're running certbot in manual mode on a machine that is not
	your server, please ensure you're okay with that.
	
	Are you OK with your IP being logged?
	-------------------------------------------------------------------------------
	(Y)es/(N)o: Y
	
	-------------------------------------------------------------------------------
	Create a file containing just this data:
	
	9HYamP-gDXrwB4UeXpry9zwL59e3G03VPbxNv_QIsvU.-QfnFZBzb2toRS_u4TSdmjy38XELnboJMzn2VBx5Ydo
	
	And make it available on your web server at this URL:
	
	http://www.xbaod.com/.well-known/acme-challenge/9HYamP-gDXrwB4UeXpry9zwL59e3G03VPbxNv_QIsvU
	
	-------------------------------------------------------------------------------
	Press Enter to Continue9HYamP-gDXrwB4UeXpry9zwL59e3G03VPbxNv_QIsvU.-QfnFZBzb2toRS_u4TSdmjy38XELnboJMzn2VBx5Ydo
	Waiting for verification...
	Cleaning up challenges
	
	IMPORTANT NOTES:
	 - Congratulations! Your certificate and chain have been saved at:
	   /etc/letsencrypt/live/www.xbaod.com/fullchain.pem
	   Your key file has been saved at:
	   /etc/letsencrypt/live/www.xbaod.com/privkey.pem
	   Your cert will expire on 2017-11-20. To obtain a new or tweaked
	   version of this certificate in the future, simply run
	   letsencrypt-auto again. To non-interactively renew *all* of your
	   certificates, run "letsencrypt-auto renew"
	 - If you like Certbot, please consider supporting our work by:
	
	   Donating to ISRG / Let's Encrypt:   https://letsencrypt.org/donate
	   Donating to EFF:                    https://eff.org/donate-le
	
	[root@k8s-master letsencrypt]# 

生成证书成功，证书文件保存在当前服务器的/etc/letsencrypt/live/www.xbaod.com/目录下


参考文档：


[http://www.yuchenw.com/help/show.asp?id=2728](http://www.yuchenw.com/help/show.asp?id=2728)

[https://segmentfault.com/a/1190000005797776](https://segmentfault.com/a/1190000005797776)

[http://www.nginx.cn/115.html](http://www.nginx.cn/115.html)

[http://www.ituring.com.cn/article/217692](http://www.ituring.com.cn/article/217692)

[http://www.restran.net/2017/01/24/nginx-letsencrypt-https/](http://www.restran.net/2017/01/24/nginx-letsencrypt-https/)