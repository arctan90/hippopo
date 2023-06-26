# 把本机eth0的转发到eth1
前置条件
<p>
网关机器有双网卡，网关的内网地址在eth1上内网地址是192.168.0.1，网关对外的EIP在eth0上。内网的服务地址是192.168.0.33:30888
</p>

```shell
# 先开启转发
echo 1 > /proc/sys/net/ipv4/conf/all/forwarding
# 把eth0的30888端口的所有请求的目的地址改为192.1168.0.33:30888达到转发的目的
iptables -t nat -A PREROUTING -p tcp --dport 30888 -i eth0 -j DNAT --to 192.168.0.33:30888
# 由于网关上有vxlan，此时ipv4的路由走eth1（ip route可以看到），所以这里把出方向为eth1的请求的原地址换成网关的eth1地址
iptables -t nat -A POSTROUTING -p tcp --dport 30383 -o eth1 -j SNAT --to-source 192.168.0.1
```

# 举例
ip route add <目标网络> via <下一跳地址> dev <接口> table <table号>
```shell
ip route add 172.170.0.191 dev raven-wg0 scope link mut 1420 table 9028
```
"adding route" dst="172.24.0.0/20" via="240.17.0.192" src="172.24.0.0/20" table=9027 src是本机/pod IP
```shell
ip route add 172.24.0.0/20 via 240.17.0.192 dev raven0 src 172.17.0.193 onlink mtu 1420 table 9027
```

