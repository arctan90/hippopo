# 安装
## Ubuntu
https://launchpad.net/ubuntu/+source/wireguard-linux-compat
https://launchpad.net/ubuntu/+source/wireguard

https://y2k38.github.io/posts/how-to-setup-wireguard-vpn-server/

## openwrt
https://git.openwrt.org/?p=openwrt/openwrt.git;a=blob;f=package/network/utils/wireguard-tools/Makefile
## centOS8
```shell
$ sudo yum install yum-utils epel-release
$ sudo yum-config-manager --setopt=centosplus.includepkgs="kernel-plus, kernel-plus-*" --setopt=centosplus.enabled=1 --save
$ sudo sed -e 's/^DEFAULTKERNEL=kernel-core$/DEFAULTKERNEL=kernel-plus-core/' -i /etc/sysconfig/kernel
$ sudo yum install kernel-plus wireguard-tools
$ sudo reboot
```
或者
```shell
$ sudo yum install elrepo-release epel-release
$ sudo yum install kmod-wireguard wireguard-tools
```
或者
```shell
$ sudo yum install epel-release
$ sudo yum config-manager --set-enabled PowerTools
$ sudo yum copr enable jdoss/wireguard
$ sudo yum install wireguard-dkms wireguard-tools
```
# CentOS7
```shell
# yum install oraclelinux-developer-release-el7
# yum-config-manager --disable ol7_developer
# yum-config-manager --enable ol7_developer_UEKR6
# yum-config-manager --save --setopt=ol7_developer_UEKR6.includepkgs='wireguard-tools*'
# yum install wireguard-tools
```
或者
```shell
yum install epel-release https://www.elrepo.org/elrepo-release-7.el7.elrepo.noarch.rpm
yum install -y yum-plugin-elrepo
yum install kmod-wireguard wireguard-tools -y
```

# 使用
假设有两台机器，A和B
1. 在每台机器上使用 `wg genkey > private` 生成私钥。
2. 在每台机器上用刚才生成的私钥再生成公钥 `wg pubkey < private`
---
3. 在机器A上添加`ip link`和`ip addr`
```shell
ip link add wg0 type wireguard
ip addr add 10.0.0.1/24 dev wg0
```
4. 在机器A上设置私钥
```shell
wg set wg0 private-key ./private
```
5. 启动机器A上的wg0网络设备
```shell
ip link set wg0 up
```
---
6. 在机器B上添加`ip link`和`ip addr`，导入私钥，启动网卡
```shell
ip link add wg0 type wireguard
ip addr add 10.0.0.2/24 dev wg0
wg set wg0 private-key ./private
ip link set wg0 up
```
7. 在机器A和B上用wg查询详情
---
8. 在机器A上添加对端
```shell
wg set wg0 peer B端公钥 allowed-ips B端的wg0虚拟ip即10.0.0.2/24 endpoint B端对外网卡IP+UDP监听端口51820
```
9. 在机器B上添加对端
```shell
wg set wg0 peer A端公钥 allowed-ips A端的wg0虚拟ip即10.0.0.1/24 endpoint A端对外网卡IP+UDP监听端口51820
```
---
10. check 用wg0的IP ping可通，用`wg`查询状态能看到peer信息
## 命令行接口
A new interface can be added via ip-link(8), which should automatically handle module loading:
```shell
ip link add dev wg0 type wireguard
```
(Non-Linux users will instead write wireguard-go wg0.)

An IP address and peer can be assigned with ifconfig(8) or ip-address(8)
```shell
ip address add dev wg0 192.168.2.1/24
```
Or, if there are only two peers total, something like this might be more desirable:
```shell
ip address add dev wg0 192.168.2.1 peer 192.168.2.2
```
The interface can be configured with keys and peer endpoints with the included wg(8) utility:
```shell
wg setconf wg0 myconfig.conf
```
or
```shell
wg set wg0 listen-port 51820 private-key /path/to/private-key peer ABCDEF... allowed-ips 192.168.88.0/24 endpoint 209.202.254.14:8172
```
Finally, the interface can then be activated with ifconfig(8) or ip-link(8):
```shell
ip link set up dev wg0
```

帮助信息
```shell
wg show
```
