# 概述
在多个subnet中的两个endpoint上创建tunnel和主机到主机的VPN类似，区别在于添加的是left subnet和right subnet。
我们使用also=关键字避免向每个连接中添加相同信息。

## 两个主机的网络
192.0.2.254/24 eth0 WEST eth1 192.1.2.23 --[internet]-- 192.1.2.45 eth1 EAST eth0 192.0.1.254/24

## 配置
```shell
# /etc/ipsec.conf

config setup
    #logfile=/var/log/pluto.log
# mysubnet用于ipv4的subnet到subnet
conn mysubnet
     also=mytunnel
     leftsubnet=192.0.1.0/24
     rightsubnet=192.0.2.0/24
     auto=start
# mysubnet6用于ipv6的subnet到subnet
conn mysubnet6
     also=mytunnel
     connaddrfamily=ipv6
     leftsubnet=2001:db8:0:1::/64
     rightsubnet=2001:db8:0:2::/64
     auto=start
# mytunnel用于VPN服务器到VPN服务器
conn mytunnel
    leftid=@west
    left=192.1.2.23
    leftrsasigkey=0sAQOrlo+hOafUZDlCQmXFrje/oZm [...] W2n417C/4urYHQkCvuIQ==
    rightid=@east
    right=192.1.2.45
    rightrsasigkey=0sAQO3fwC6nSSGgt64DWiYZzuHbc4 [...] D/v8t5YTQ==
    authby=rsasig
```
# 测试
从west发包
```shell
# ping -n -c 4 -I 192.0.1.254 192.0.2.254
PING 192.0.2.254 (192.0.2.254) from 192.0.1.254 : 56(84) bytes of data.
64 bytes from 192.0.2.254: icmp_seq=1 ttl=64 time=0.XXX ms
64 bytes from 192.0.2.254: icmp_seq=2 ttl=64 time=0.XXX ms
64 bytes from 192.0.2.254: icmp_seq=3 ttl=64 time=0.XXX ms
64 bytes from 192.0.2.254: icmp_seq=4 ttl=64 time=0.XXX ms
--- 192.0.2.254 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time XXXX
rtt min/avg/max/mdev = 0.XXX/0.XXX/0.XXX/0.XXX ms
```
之所以指定ping的源地址192.0.2.254，是因为linux上会自动选择就近ip发起ping。在上面的例子中就近IP是192.1.2.23，并且这个ip不属于子网192.0.2.0/24。
如果希望所有走网关的流量加密，可以使用网关各自的私网ip地址。此时使用附加选项 leftsourceip= 和 rightsourceip=
```shell
conn mysubnet
     also=mytunnel
     leftsubnet=192.0.1.0/24
     leftsourceip=192.0.1.254
     rightsubnet=192.0.2.0/24
     rightsourceip=192.0.2.254
     auto=start
```
这样声明之后libreswan，会使用src <ipaddress>参数为远端subnet自动添加路由达到这个目的。
或者你可以添加主机到主机的tunnel，上面的mytunnel，但你需要添加主机到subnet和subnet到主机的tunnel。这有些笨重。因此人们常用xxxsourceip=选项。