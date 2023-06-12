# docker镜像压缩
压缩 docker save myimage | gzip > somename.tgz
解压 gunzip -c somename.tgz | docker load

# Docker cgroup driver使用systemd
```shell
vi /etc/docker/daemon.json
加
{
  "exec-opts":["native.cgroupdriver=systemd"]
}
sudo systemctl restart docker
```
