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