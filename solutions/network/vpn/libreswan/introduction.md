# 概述
https://libreswan.org/ </p>
libreswan软件实现了IPSec和IKE这种最为广泛和标准化的VPN协议。这两种标准有IETF维护。</p>
## 历史沿革
Libreswan从FreeS/WAN(1995 - 2003)开始，历经Openswan(2003 - 2011)，到今天的libreswan。</p>
## 协议支持
支持IKE v1 v2版本。在linux上使用内建的XFRM的IPSec二进制包，使用NSS加密库
# 安装
## 从源码安装
CentOS8
```shell
yum install audit-libs-devel bison curl-devel flex \
                gcc ldns-devel libcap-ng-devel libevent-devel \
                libseccomp-devel libselinux-devel make nspr-devel nss-devel \
                pam-devel pkgconfig systemd-devel unbound-devel xmlto
```
CentOS7
```shell
yum install audit-libs-devel bison curl-devel fipscheck-devel flex \
                gcc ldns-devel libcap-ng-devel libevent-devel \
                libseccomp-devel libselinux-devel make nspr-devel nss-devel \
                pam-devel pkgconfig systemd-devel unbound-devel xmlto
```
Ubuntu
```shell
apt-get install libnss3-dev libnspr4-dev pkg-config libpam-dev \
                libcap-ng-dev libcap-ng-utils libselinux-dev \
                libcurl3-nss-dev flex bison gcc make libldns-dev \
                libunbound-dev libnss3-tools libevent-dev xmlto \
                libsystemd-dev
```
其他依赖
```shell
nss, iproute2, iptables, sed, awk, bash, cut, procps-ng, which
```
## 编译
源码 https://github.com/libreswan/libreswan/
```shell
    make programs
    sudo make install
```
或者不要说明文档的编译
```shell
    make base
    sudo make install-base
```

## 启动
```shell
systemctl enable ipsec.service
systemctl start ipsec.service
```

安装nss
```shell
ipsec initnss
```