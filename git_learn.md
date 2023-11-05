# 与远程仓库绑定
```
git init # 将本地文件夹 git 初始化
git add . # 将本地文件进行添加, . 是全部文件, 可以通过具体文件名进行添加。添加是添加到本地缓存区
git commit -m "提交说明" # 将缓存区的文件进行提交, 并加上提交说明
git branch -M main # -m的作用是重命名, 此处应该是创建一个名为main的分支
git remote add origin git@github.com:zhangjingxian1998/some_note.git # 添加远程分支, 分支名称为oring, 后面的网址是git仓库
git push -u origin main # 推送分支到远程,-u 参数的意思是将这个远程分支设为本地分支的 upstream（上游），以后的 push 或 pull 可以直接使用 git push 或 git pull 命令。如果不使用 -u 参数，则需要在每次 push 或 pull 命令后指定远程分支名。
```