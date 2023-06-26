# 检测
docker run --rm -e NVIDIA_VISIBLE_DEVICES=all nvidia/cuda:11.6.2-base-ubuntu20.04 nvidia-smi
<p>参考 https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html

# 安装k8s
## 每个节点配置/etc/hosts
所有待加入集群的  机器ip 主机名
## 装docker
```shell
yum install -y yum-utils device-mapper-persistent-data lvm2
yum-config-anager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
yum make cache fast
yum install -y docker-ce-19.03.15-3.el7
systemctl start docker
systemctl enable docker
systemctl status docker
docker -v
```
## docker cgroup driver使用systemd
```shell
# vi /etc/docker/daemon.json 加
{
  "exec-opts":["native.cgroupdriver=systemd"]
}
system restart docker
```

## k8s三件套
### 加k8s源
todo
### 安装
对齐版本
```shell
yum install -y kubelet-1.20.11 kubeadm-1.20.11 kubectl-1.20.11
systemctl start kubelet
```
### 关闭防火墙
systemctl stop firewalld
systemctl disable firewalld
###
挑选一台主机(172.16.246.71)运行
```shell
kubeadm init \
--apiserver-advertise-address=172.16.246.71 \
--image-repository xxxx \
--kubernetes-version v1.20.11 \
--service-cidr=10.96.0.0/12 \
--pod-network-cidr=10.244.0.0/16
```
如果装失败了先回滚
```shell
kubeadm reset
```
装好了之后有个打印，在其他节点上执行一下join就行
## GPU驱动
https://docs.nvidia.com/datacenter/tesla/tesla-installation-notes/index.html
```shell
#centOS7的一个bug
yum install -y mesa-libGL.x86_64 mesa-libEGL.x86_64
# 1 一些必要的工具
yum install -y tar bzip2 make auto make gcc gcc-c++ pciutils elfutils-libelf-devel libglvnd-devel iptables firewalld vim bind-utils wget
# 2 装驱动
yum install -y https://dl.fedoraproject.org/pub/epel/eepel-release-latest-7.noarch.rpm
distribution=rhel7
ARCH=$( /bin/arch )
yum -y install yum-utilsyum-config-manager --add-repo http://developer.download.vnidia.com/compute/cuda/repos/$distribution/${ARCH}/cuda-$distribution.repo
yum install -y kernel-devel-$(uname -r)kernel-headers-$(uname -r)
yum clean expire-cache

yum install -y nvidia-driver-latest-dkms
# 3. 设置环境变量
export PATH=/usr/local/cuda-12.0/bin${PATH:+:${PATH}}
export LD_LIBRARY_PATH=/usr/local/cuda-12.0/lib64\
        ${LD_LIBRARY_PATH:+:${LD_LIBRARY_PATH}}
```
## 安装GPU docker运行时 （CDI）
https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/install-guide.html
