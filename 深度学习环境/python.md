# 虚拟环境(conda) [推荐]

## windows
[官网](https://www.anaconda.com/download/#macos)

如果下载不动使用国内镜像源，下载对应的Windows-x86_64.exe文件

[镜像源](https://mirrors.tuna.tsinghua.edu.cn/anaconda/archive/?C=N&O=A)

[安装教程](https://blog.csdn.net/weixin_48697962/article/details/125941609)
```
# 主要命令
conda init                          # 初始化一个base
conda env list                      # 查看所有的虚拟环境，未创建过应该只有一个base
conda create -n env_name python=3.X # 创建名为env_name的虚拟环境，python版本为3.X
conda activate env_name             # 根据虚拟环境名称启动对应的虚拟环境
conda remove -n env_name --all      # 根据虚拟环境名称启动对应的虚拟环境以及安装包
```

如果有报错或者ide无法发现python解释器，检查一下[环境变量](https://blog.csdn.net/m0_67484548/article/details/123434033)
## linux
在[官网](https://repo.anaconda.com/archive)或[镜像网站](https://mirrors.tuna.tsinghua.edu.cn/anaconda/archive)找到对应的.sh文件下载连接
```
如：
wget --no-check-certificate https://repo.anaconda.com/archive/Anaconda3-2023.03-1-Linux-x86_64.sh
或
wget --no-check-certificate https://mirrors.tuna.tsinghua.edu.cn/anaconda/archive/Anaconda3-5.3.1-Linux-x86_64.sh

bash Anaconda*****.sh                   # 提示回车的地方按回车，提示yes/no的地方输入yes，最后安装vscode选no，阅读大段协议处按q退出阅读。此处ctrl+back是回退功能。
source ~/.bashrc                        # 安装完成后刷新环境变量。如果安装中未添加环境变量，需手动添加。
vim ~/.bashrc                           # 在文件末尾处添加
export PATH=~/home/anaconda3/bin:$PATH  # $PATH是所系统中所有的环境变量，可通过echo $PATH打印出来，~/.../bin是新的环境变量，二者用 : 连接。
source ~/.bashrc                        # 即可。
```
# 本地安装

## windows
[官网](https://www.python.org/)

建议3.7版本及以上

[安装教程](https://zhuanlan.zhihu.com/p/569019068)
## linux
需要手动编译，以python3.8.10为例
在[官网](https://www.python.org/)找到想要安装的版本。
```
wget https://www.python.org/ftp/python/3.8.10/Python-3.8.10.tgz
tar -zxvf Python-3.8.10.tgz                                     # 解压
cd Python-3.8.10                                                # 进入刚解压的文件夹
./configure --prefix=~/.local/python3.8 --with-ssl              # 安装位置，随意
make && make install                                            # 编译安装
ln -s /**/.local/python3.8/bin/python3.8 /**/.local/bin/python3.8
ln -s /**/.local/python3.8/bin/pip3 /**/.local/bin/pip3         # 将python和pip放入环境变量中，注意使用绝对路径。
```

[conda 修改pip路径](https://it.cha138.com/nginx/show-4725270.html)

主要是应对在虚拟环境中创建多个相同版本的环境导致共享Pip包

在对应的环境中lib/python3.X/

vim site.py 

更改ENABLE_USER_SITE = False

# 设置镜像源 (重要)
```
pip config set global.index-url https://pypi.mirrors.ustc.edu.cn/simple/
```