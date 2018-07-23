#!/bin/bash

git add -A
git commit -am "更新内容"
git push origin master

#更新public库
hugo --theme=jane --buildDrafts --baseUrl="https://zouhuigang.github.io/"
cd public
git add -A
git commit -am "更新网站内容"
git push origin master
