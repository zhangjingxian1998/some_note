docker pull ubuntu:20.04
docker run -it -e LANG=C.UTF-8 ubuntu:20.04
# 修改 apt源
sed -i 's@http://archive.ubuntu.com/ubuntu/@http://mirrors.aliyun.com/ubuntu/@g' /etc/apt/sources.list

apt update
apt install -y curl wget vim  # 安装常见软件

apt install python3.9  
apt install -y python3-distutils

curl 'https://bootstrap.pypa.io/get-pip.py' > get-pip.py   #  下载 pip
python3.9 get-pip.py  # 安装pip

ln -s /usr/bin/python3.9 /usr/bin/python  # 建立软连接

# 验证安装
python --version
pip --version

docker commit <container-id> ubuntu:20.04_py39