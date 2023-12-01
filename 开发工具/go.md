# Linux安装
[官网](https://go.dev/dl/)下载tar包
```
tar -C /usr/local -zxvf go*.linux-amd64.tar.gz
vim /etc/profile

export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
source /etc/profile
go version
# 新版改成如下链接
go env -w GO111MODULE=on
go env -w GOPROXY=https://proxy.golang.com.cn,direct
```

## 在vscode中配置
```
Ctrl + Shift + P
GO: Install/Updata Tools
install all
# 可以进行debug
go mod init name 
```