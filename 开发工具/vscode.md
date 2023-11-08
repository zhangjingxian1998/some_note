开源，体积小，开启快

[vscode下载](https://code.visualstudio.com/)

[入门教程](https://blog.csdn.net/jiong9412/article/details/123216812)

> #### 推荐插件
```
Python
Jupyter
Chinese (Simplified)
Remote-SSH
Python Image Preview
```

> #### 快捷键
```
ctrl+/      # 注释选中的行
tab         # 选中行右移
shift+tab   # 选中行左移
```

> #### 选择python解释器
python环境安装好，python插件安装好后，创建.py文会自动检测python解释器，选择在右下角Python右边。如果检测不到，检查python[环境变量](https://blog.csdn.net/m0_67484548/article/details/123434033)是否配置好。

> #### 远程设置
```
1、cmd输入ssh查看系统是否安装服务，如果没有使用Add-WindowsCapability -Online -Name OpenSSH.Client~~~~0.0.1.0命令安装
2、vscode安装Remote-SSH插件
3、左侧插件栏远程资源管理器(Remote Explore)
4、选中SSH设置齿轮打开SSH配置文件(Open SSH Config File)
5、选中*****\.ssh\config
```
```
Host alias              # 服务器名字，随便起
    HostName hostname   # 服务器ip
    User user           # 登录用户名
```
例：
```
Host X9dri-f
    HostName 192.168.2.164
    User zhangjx
    port 22             # 22为默认端口，有时不是22
```
这样设置后每次登录只需要输入用户密码即可。
> #### 远程免密登录
只需要将本地的ssh公钥上传至服务器即可。[原理](https://blog.csdn.net/weixin_44981916/article/details/103497631)
```
# 本地
ssh-keygen -t rsa       # 一路回车，在C:user/username/.ssh文件中生成.pub公钥，复制公钥信息。

# 服务器
ssh-keygen -t rsa       # 一路回车，生成~/.ssh文件夹。
cd .ssh
vim authorized_keys     # 将本地公钥信息粘贴进去，保存退出，即可免密登录。
```
如果依旧需要输入密码，更改Host设置，指定公钥位置
```
Host X9dri-f
    HostName 192.168.2.164
    User zhangjx
    port 22
    IdentityFile ~/.ssh/id_rsa
```