### 设置环境变量

	vi /etc/profile 
	
在/etc/profile/后面添加 

	export PATH=$PATH:/usr/bin
	export GOROOT=$HOME/go
	export GOBIN=$GOROOT/bin
	export GOARCH=386
	export GOOS=linux
	export GOPATH=$HOME/workspacego
	export PATH=.:$PATH:$GOBIN 

使环境变量生效
 
 	 source /etc/profile 
 	 
 
 
 
 ### 问题
 
q1:sudo source /etc/profile会发现没source命令

a1: 

	source /etc/profile 
	或 
	$ . /etc/profile 
	就可以了，不需要加sudo。 
	source是一个内部命令，使用man builtins查看。

