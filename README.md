# zouhuigang.github.io


访问域名:

[zouhuigang.github.io](zouhuigang.github.io)


发布到master:

	cd D:\mnt\zouhuigang.github.io
	hugo --theme=hugo-travelify-theme --buildDrafts --baseUrl="https://zouhuigang.github.io/"
	cd public

	git add -A
	git commit -am "备注"
	git push

发布到md分支:

	创建md分支
	git branch md

	切换分支
	git checkout md

	将分支上传
	git push origin md