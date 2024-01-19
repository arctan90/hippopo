# curl 只看状态码
curl -I -m 10 -o /dev/null -s -w%{http_code}

# tcpdump
tcpdump -XvennSs 0 -i eth0 'tcp and (((ip[2:2]-((ip[0]&0xf)<<2))-((tcp[12]&0xf0)>>2))!=0)'

# 查看centos系统
uname -m && cat /etc/redhat-release

# 安装conda
w
# 阿里云ecs磁盘扩容
1. parted /dev/vda print free 可以看到有free空间
2. 由于/dev/vda2挂载到了/

# 查看所有的服务项
systemctl list-unit-files --type=service

# 磁盘数据恢复工具 test
## 安装
sudo apt-get install testdisk

