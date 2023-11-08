|指令 |作用 |
|-|-|
|<font color="#FF000">ls</font> |查看当前目录下<font color="#FF000">显示</font>的文件和文件夹|
|<font color="#FF000">ls --all</font>|查看当前目录下<font color="#FF000">所有</font>的文件和文件夹|
|<font color="#FF000">ls -a</font>|查看当前目录下<font color="#FF000">所有</font>的文件和文件夹|
|ls -n|查看文件间的链接信息|
|ls \| grep file_name | 在结果中搜索file_name|
|<font color="#FF000">cd</font>|在文件夹间跳转, 可使用绝对路径和相对路径|
|<font color="#FF000">cd ..</font>| 返回<font color="#FF000">上层目录</font>, 如果不存在上层目录不跳转|
|mv|<font color="#FF000">移动</font>命令move, 可用来重命名|
|mv old new| 将old文件或文件夹移动为new|
|cp| <font color="#FF000">复制</font>命令copy, 可使用绝对路径和相对路径|
|cp old new| 将old<font color="#FF000">文件</font>复制为new|
|cp -r old new| 将old<font color="#FF000">文件夹</font>复制为new|
|rm|<font color="#FF000">删除</font>命令, 可使用绝对路径和相对路径|
|rm file_name|删除文件 remove|
|rm -r folder_name|删除空文件夹|
|rm -rf folder_name|删除文件夹|
|ln -s old new|创建软链接, 相当于windows的快捷方式, 使用绝对路径|
|ps -ef| 查看所有进程|
|kill -9 id|根据进程id杀死进程|