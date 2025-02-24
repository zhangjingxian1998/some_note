![img](https://img-blog.csdnimg.cn/20210410160002817.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0gyMDAxMDI=,size_16,color_FFFFFF,t_70#pic_center)

本地变更记录
这里很好理解，因为git的提交分为commit,本地提交（这里主要是表明需要push的改变需要先commit），远程提交push 将本地提交推送到远程。
本地变更记录可以非常轻易的查看自上次commit或push后存在的变更
本地提交记录
将变更commit后，现存的变更就会发生改变，可以分阶段或分功能的commit,以方便了解每次commit的主要变更原因，例如代码初始化，修复XX功能等
本地未push记录
可以查看都进行过哪些commit的提交，但是还未曾提交到远程
拉取线上变更pull
该功能主要是用于更新当前分支的远程更新
fetch 更新
该功能主要是用于在获取远程分支后，其他人在远程地址上有了新的push的分支情况，本地是没有该分支的，也就无法checkout该分支，此时需要fetch
远程地址管理
这个主要是如果我们项目的远程地址发生变化的时候，需要修改远程地址，一般这种情况比较少，因为checkout项目是一定存在远程地址的，但是当我

# 远程仓库（Remote）

# 本地仓库（Repository）

# 工作区（workspace）

当前分支

# 克隆（clone）

```
git clone SSH
```

将远程仓库(Remote)下载到本地(Repository)

# 抓取（fetch）

```
git fetch --all
```

将远程仓库所有分支加载到本地

# 切换分支（checkout）

```

```

切换分支

# 暂存（add）

将工作区所作修改暂时存放到暂存区（index）

# 提交（commit）

将暂存区修改提交到本地仓库

# 推送（push）

将本地仓库推送到远程仓库

# 拉取（pull）

拉取远程仓库当前分支所有文件

# 分支合并

两个分支的更改合并

# 合并冲突

两个不同分支对同一文件更改合并时冲突，进行选择

# 版本回滚

网页操作

# 分支比较

插件