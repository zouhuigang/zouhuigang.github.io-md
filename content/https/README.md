### 使用docker生成https证书


	docker pull quay.io/letsencrypt/letsencrypt:latest


方式1：

	docker run --rm -p 80:80 -p 443:443  -v /etc/letsencrypt:/etc/letsencrypt  quay.io/letsencrypt/letsencrypt auth  --standalone -m 952750120@qq.com --agree-tos  -d www.anooc.com -d www.xbaod.com

启动nginx:

	docker run --name nginx -p 80:80 -p 443:443 \
    -v /etc/nginx/conf.d:/etc/nginx/conf.d \
    -v /etc/letsencrypt:/etc/letsencrypt \
    nginx

方式2(手动)：

	docker run -it --rm -p 80:80 -p 443:443 \
    -v /etc/letsencrypt:/etc/letsencrypt \
    quay.io/letsencrypt/letsencrypt auth



自动更新证书：

	docker run --rm -p 80:80 -p 443:443 \
	    -v /etc/letsencrypt:/etc/letsencrypt \
	    quay.io/letsencrypt/letsencrypt renew \
	    --standalone

运行这个命令时，certbot会自动检查确认证书有效期，如果过期时间在一个月之内，就会自动更新。

在CoreOS中，由于没有Cron，我们需要通过systemd的timer来做定时调度，比如每个月运行一次这个renew任务就可以了，不过记得运行之前先停止nginx容器，运行之后再启动nginx容器



### Standalone

使用独立模式进行自动验证，需要在 目标服务器 上运行 Let's Encrypt 客户端，并指定 certonly 和 --standalone参数。本模式需要绑定 80 或 443 端口进行域名验证，所以如果服务器上已有web server运行并侦听这2个端口，则需要先关闭web server。

### Webroot

如果 目标服务器 已有web server运行，并且不能够关闭服务来获取和安装证书，可以使用 Webroot plugin。在运行 Let's Encrypt 客户端时指定 certonly 和 --webroot 参数，并使用 --webroot-path 或 -w 参数来指定 webroot 目录，比如 --webroot-path /usr/share/nginx/html

docker run --rm -p 80:80 -p 443:443  -v /etc/letsencrypt:/etc/letsencrypt  quay.io/letsencrypt/letsencrypt auth  --webroot -m 952750120@qq.com --agree-tos  -d www.anooc.com -d www.xbaod.com



[https://imququ.com/post/letsencrypt-certificate.html](https://imququ.com/post/letsencrypt-certificate.html)

[http://blog.csdn.net/hj7jay/article/details/54405615](http://blog.csdn.net/hj7jay/article/details/54405615)

[http://www.jianshu.com/p/5afc6bbeb28c](http://www.jianshu.com/p/5afc6bbeb28c)