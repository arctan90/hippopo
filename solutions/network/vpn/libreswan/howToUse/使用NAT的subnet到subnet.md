# Using NAT to resolve an subnet IP conflict
## 概述
VPN经常使用RFC-1918地址空间来连接网络，比如 10.0.0/8, 192.16.8.0.0/16 or 172.16.0.0/12.
当所有端点使用相同的地址空间的时候会出现问题。 
此时其中一方需要将他们的子网 NAT 到其他地方。例如远端使用 10.0.0.0/8 本端使用 10.6.6.0/24。向远端请求一个他们不用的地址范围，比如192.168.0.0/24。

## 实现
用上面这些子网创建vpn
```shell
conn vpn
    left=1.2.3.4
    leftsubnet=192.168.0.0/24
    right=5.6.7.8
    rightsubnet=10.0.0.0/8
    [...]
```
然后添加所需的 iptables NAT 规则，以避免与现有规则或 IPsec 冲突：
```shell
iptables -I POSTROUTING -m policy --dir out --pol ipsec -j ACCEPT
iptables -I POSTROUTING -s 10.6.6.0/24 -d 10.0.0.0/8 -o ethX -j SNAT --to-source 192.168.0.1
```
还可以尝试将 /24 映射到 /24，并让所有机器都可以通过这些备用 IP 地址访问，这应该可以使用另一个 iptables 规则将 DNAT 192.168.0.X/24 到 10.6.6.X/24
