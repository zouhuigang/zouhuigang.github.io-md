#cmd�еı��뷽ʽΪANSI�������Ĳ��Ǵ˱��뷽ʽ���������롣
git add -A
git commit -am "��������"
git push origin master

#����public��
hugo --theme=hugo-travelify-theme --buildDrafts --baseUrl="https://zouhuigang.github.io/"
cd public
git add -A
git commit -am "������վ����"
git push origin master
