### phpexcel有内存泄漏的问题，每次执行完成之后，服务器内存就会被占用很多，所以上传excel多了，就会解析有问题



本例子换了个库，改了配置，还是没法解决这个问题，所以先留着测试文件，之后再通过其他语言来解释excel吧



目前的解决方案是，在Linux服务器上，设置一个定时任务，当php-fpm占用内存大于多少的时候，及时把该进程给杀死。



### 将配置文件上传到服务器

​	chmod  +x  kill_php-fpm_big_memory.sh

​	./kill_php-fpm_big_memory.sh