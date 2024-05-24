# 安装好系统后

(以ubuntu 20.04.6为例)

```
sudo passwd # 设定sudo密码
su root # 启动超级管理员账号(需要输入sudo密码)
apt-get update
apt-get upgrade
apt-get install build-essential
```
sudo useradd -d /media/dataStore/gjsh -m -s /bin/bash gjsh
由于系统不带pip, 如需使用本机自带python, 可以安装pip

> #### 方法一：

```
sudo apt-get install python3-pip
```

> #### 方法二：

```
sudo apt-get install python3-distutils
wget http://bootstrap.pypa.io/get-pip.py
sudo python3 get-pip.py
sudo pip install launchpadlib -i https://pypi.tuna.tsinghua.edu.cn/simple some-package
sudo pip install --upgrade pip -i https://pypi.tuna.tsinghua.edu.cn/simple some-package
```

# 安装显卡驱动

[nvidia驱动](https://www.nvidia.cn/Download/index.aspx?lang=cn)

直接./*****.run安装

# 取消xorg占用显存

> #### 暂时方法

关闭图形界面

```
sudo init 3
```

> #### 永久方法

```
sudo Xorg -configure # 生成Xorg的可配置文件
cd /usr/share/X11/xorg.conf.d
sudo rm nvidia-drm-outputclass.conf
```

# 挂载硬盘

初次挂载需要格式化

```
sudo umount /dev/sdb1 # 按照子盘符名称先取消挂载
sudo fdisk -l #查看硬盘盘符名称
sudo fdisk /dev/sdb #根据盘符名称考试格式化
按g格式化
按q退出
sudo mkfs.ext4 -F /dev/sdb # 类型变为ext4
sudo mount /dev/sdb /home # 将硬盘挂载到/home 注意，记得提前保存重要文件，挂载完成后再导入回去
```

开机自动挂载

```
blkid # 查看硬盘UUID
sudo gedit /etc/fstab  # 更改启动挂载文件
UUID=****** /home ext4 defaults 0 2   # 加在文档末尾
chown -R ${用户组名}:${用户名} ${文件夹名} # 一定要更改权限，不然进不去
```