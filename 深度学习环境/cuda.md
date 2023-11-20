# Updata
如无编译需要，不安装本地CUDA也行

安装的torch中自带有一定的cuda和cudnn算法,可以保证基本程序的运行

```
# 检查安装的torch所对应的cuda版本。该版本不应超过英伟达驱动所支持的cuda版本
torch.version.cuda
# 检查当前环境所使用的cuda
import torch
import torch.utils.cpp_extension
torch.utils.cpp_extension.CUDA_HOME
# 当没有安装cuda, 返回None
# 有安装cuda并且设置正确, 返回对应的文件位置
# 没有cuda不代表着无法使用gpu。但如果需要编译依赖cuda算法的项目将报错。
```
conda和pip都可以指定安装torch时带有某个版本的cudatoolkit, 该cudatoolkit将保证程序的正常运行


# 安装CUDA

[cuda_toolkiot](https://developer.nvidia.com/cuda-toolkit-archive)

使用./run安装，不要安装驱动，也不安装opengl相关内容

将以下内容写入**/etc/bash.bashrc**

```
export PATH=$PATH:/usr/local/cuda-**/bin
export LD_LIBRARY_PATH=:/usr/local/cuda-**/lib64
```
# 安装cudnn

[cudnn](https://developer.nvidia.com/rdp/cudnn-download)

下载cudnn****.xz

tar -xvf cudnn-linux-x86_64-8.6.0.163_cuda11-archive.tar.xz

将对应文件拷入cuda安装路径

```
sudo cp /include/** /usr/local/cuda/include/
sudo cp /lib/libcudnn* /usr/local/cuda/lib64/
sudo chmod a+r /usr/local/cuda/include/cudnn.h
sudo chmod a+r /usr/local/cuda/lib64/libcudnn*
sudo cp include/cudnn_version.h    /usr/local/cuda/include/

验证
cat /usr/local/cuda/include/cudnn_version.h | grep CUDNN_MAJOR -A 2
```
[同一机子多个cuda](https://www.yii666.com/blog/358463.html?action=onAll)