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

在用户下.bashrc下设置
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

## 在vscode中配置
```
进入$GOPATH文件夹中
Ctrl + Shift + P
GO: Install/Updata Tools
install all

在其他GO项目中
go mod init <moudle name>
函数的入口必须是main
在除package main外的package定义的函数必须是大写字母开头

# 解决导入包会自动删除
在setting中 GO:Alternate Tools
添加
"[go]":{
        "editor.insertSpaces": false,
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": false
        }
}
```