# 内网穿透

```
wget https://github.com/fatedier/frp/releases/download/v****/frp_****_linux_amd64.tar.gz
tar -zxvf frp_****_linux_amd64.tar.gz
cd frp_****_linux_amd64
```

```
#云服务器端设置，修改frps.ini文件
vim frps.ini
bind_port 7000
（此处为默认值，就是你要将本地端与云服务器端的转发端口号）云服务器中安全策略里面得设置开放此端口
# 面板访问
dashboard_port = 7500
dashboard_user = admin
dashboard_pwd = admin

# 设置完后启动
nohup ./frps -c ./frps.ini &
```

```
# 本地机设置
vim frpc.ini

[common]
server_addr = x.x.x.x # 你的公网ip
server_port = 7000	# 与server的bind_port设置一致，用于访问端口

[ssh]
type = tcp
local_ip = 127.0.0.1
local_port = 22
remote_port = 6000	# 你在外部访问的接口（此接口为使用6000 port）

# 设置完后启动
nohup ./frpc -c ./frpc.ini &
```

```
全部开启完成后，在确认三方都有网，且都开启ssh服务的情况下即可登录测试
ssh -p remote_port 用户名@x.x.x.公网ip
```

```
ps -ef | grep sshd
#查看服务是否开启
#如果ssh服务未开启，尝试
sudo apt-get install openssh-server
sudo service ssh start

```

> #### 开机自动启动

[Ubuntu下实现Frpc自启动 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/521448626)

在\etc\systemd\system中创建一个**.service文件，内容为

```
[Unit]
Description=My Frp Client Service - %i
After=network.target syslog.target
Wants=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
ExecStart=/bin/bash -c '/home/hsp/Desktop/frpc/frp_0.42.0_linux_amd64/frpc -c /home/hsp/Desktop/frpc/frp_0.42.0_linux_amd64/frpc.ini'
 # 路径需要根据自身情况修改
[Install]
WantedBy=multi-user.target
```

然后执行命令

```
sudo systemctl enable **
sudo systemctl start **
sudo systemctl status **
```

echo $PATH