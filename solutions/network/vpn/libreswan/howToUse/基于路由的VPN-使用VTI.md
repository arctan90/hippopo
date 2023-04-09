# 概述
VPN tunnel一般基于IPSec策略创建。这种叫做基于策略的VPN。在libreswan中，这些策略通过leftsubnet=和rightsubnet=选项以及可选的
leftprotoport= 和 rightprotport=选项指定。libreswan运行创建基于路由的VPN。此时leftsubnet=0.0.0.0/0 并且 rightsubnet=0.0.0.0/0。
此时（谁？）会创建一个irtual Tunnel Interface ("VTI")设备并加载到IPSec策略上。这样路由到这个VTI的包都会被加密。这种方式安全度较低，不过便于管理。
因为你只需要更新路由表而不是添加、修改IPSec策略。
libreswan支持使用ipsec0 interface的KLIPS，当使用XFRM/NETKEY的时候是不支持这种特性的。
使用VTI的另一个好处在于你拥有一个真正的interface(不像 nflogXX interface)，支持iptables的防火墙，运行tcpdump等等。
它甚至在非VTI设备上安装 tcpdump? 所以它真的像是旧KLIPS ipsec0 interface的功能。
## 版本
libreswan-3.18起开始支持VTI。
# 扩展的配置项

| option	                                                              | description                                                                                                                                                                          |
|----------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| mark=	                                                               | The mark number to use for this connection's IPsec SA policy. It will be used for all instances as well.                                                                             |
| markin= / markout=                                                   | 	Same as the above but setting different marks for inbound and outbound.                                                                                                             |
| vti-interface=	The interface name of the VTI device to use, eg vti01 |
| vti-routing=                                                         | 	Whether routes should automatically be created into the VTI device (default yes)                                                                                                    |
| vti-shared=                                                          | 	Whether this vti device is shared with other connections (default no). if not shared, and not a template connection with %any, the VTI device will be deleted when tunnel goes down |
| leftvti=address/mask                                                 | 	Configure the address/mask on the vti-interface when connection is established.                                                                                                     |
# 配置示例
```shell
conn routed-vpn
    left=192.1.2.23
    right=192.1.2.45
    authby=rsasigkey
    leftsubnet=0.0.0.0/0
    rightsubnet=0.0.0.0/0
    auto=start
    # route-based VPN requires marking and an interface
    mark=5/0xffffffff
    vti-interface=vti01
    # do not setup routing because we don't want to send 0.0.0.0/0 over the tunnel
    vti-routing=no
    # If you run a subnet with BGP (quagga) daemons over IPsec, you can configure the VTI interface
    leftvti=10.0.1.1/24

```
这样配置会自动创建interface vti01。如果你想加密所有走10.0.0.0/8的数据，只需要
```shell
ip route add 10.0.0.0/8 dev vti01
```
当然可以配置更复杂的路由规则，如果你想看预封装和封装后的数据表
```shell
tcpdump -i vti01 -n
```
你可以添加防火墙规则
```shell
# do not allow IRC traffic on port 6666
iptables -I INPUT -j DROP -p tcp --dport 6666 -i vti01
```
查看封装后和预解封装，假设使用en0作为数据interface
```shell
tcpdump -i en0 -n esp or udp port 4500
```
还可以查看vti tunnel的信息
```shell
ip tunnel show
ip -s tunnel show
ifconfig vti01
```
# 为所有的VPN Client创建单VTI
如果你运行的是VPN服务上使用了一个VTI设备， 所有预加密和后解密的数据都会出现在VTI设备上。所有后加密加密和预解密数据都会出现在常规物理Interface上。
用来解决监控问题
```shell
conn roadwarriors
    # Regular certificate based VPN server
    left=1.2.3.4
    leftsubnet=0.0.0.0/0
    right=%any
    rightaddresspool=10.0.1.0/24
    authby=rsasig
    leftcert=mycert
    leftid=%fromcert
    auto=add
    rekey=no
    # Create route-based VPN using VTI
    mark=12/0xffffff
    vti-interface=vti02
    vti-routing=yes
```
现在可以监控所有的流量了
```shell
tcpdump -i vti02 -n
```
# VTI相关的issue
updown script
The standard updown script (_updown.netkey) tries to autodetect what scenario is being deployed and will reconfigure the VTI interface and network options accordingly. As this feature is still new, we expect it to not be perfect yet. If your scenario is not covered, please contact the developers and explain your use case.

interface options
Each interface has a standard set of options associated with it. These can be found in /proc/sys/net/ipv4/conf/vtiXXX/*. The default updown script sets some of these (rp_filter, forwarding, policy_disable) but specific use cases might require different settings. Setting disable_xfrm will cause the interface to completely fail, so do not do that.

MTU
We noticed different kernels create different MTU sizes for new VTI devices. Currently the MTU is not set by libreswan.

combing different IPsec gateways into one VTI interface [PARTIALLY SOLVED]
You cannot yet have a vti tunnel that uses local 0.0.0.0 and remote 0.0.0.0. Kernel patches for these are pending. This means that you cannot yet combine two different gateways on the same VTI device, that have a different local IP address. However, you can combine different connections that use the same local address but use a different remote address by using the vti-shared=yes option. Obviously, the marks and interface name must be the same for all shared connections.

```shell
conn west
     left=1.2.3.4
     right=6.7.8.9
     [...]
     mark=10/0xffffff
     vti-interface=vti01
     vti-routing=yes
     vti-shared=yes

conn east
     left=10.11.12.13
     right=9.8.7.6
     [...]
     mark=10/0xffffff
     vti-interface=vti01
     vti-routing=yes
     vti-shared=yes
```