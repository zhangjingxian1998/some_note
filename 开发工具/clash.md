# clash for linux
下载解压后
'''
cd clash-for-linux
vim .env
export CLASH_URL='https://9bd4028e.ghelper.me/subs/clash/e64615699bd4028e7ed0e1b6e7eb9655'
export CLASH_SECRET='123'

bash start.sh
vim /etc/bash.bashrc
末尾添加 source /etc/profile.d/clash.sh

http://127.0.0.1:9090/ui

base_URL:http://127.0.0.1:9090
Secrte: $CLASH_SECRET
将9090端口转发
'''

# 开机自启
在/etc/systemd/system中创建一个**.service文件，内容为

```
[Unit]
Description=clash for linux - %i
After=network.target syslog.target
Wants=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
ExecStart=/bin/bash -c '/clash-for-linux/bin/clash-linux-amd64 -d /clash-for-linux/conf'

[Install]
WantedBy=multi-user.target
```

然后执行命令

```
sudo systemctl enable **
sudo systemctl start **
sudo systemctl status **
```