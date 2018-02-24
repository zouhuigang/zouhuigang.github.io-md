#cmd中的编码方式为ANSI，若中文不是此编码方式则会出现乱码。
git add -A
git commit -am "更新内容"
git push origin master

#更新public库
hugo --theme=hugo-travelify-theme --buildDrafts --baseUrl="https://zouhuigang.github.io/"
cd public
git add -A
git commit -am "更新网站内容"
git push origin master
