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

# 有的时候/var/run/docker.sock不见了
[方案](https://superuser.com/questions/1741326/how-to-connect-to-docker-daemon-if-unix-var-run-docker-sock-is-not-available)
核心思路
```shell
sudo vim /lib/systemd/system/docker.service
在打开的文件中找到ExecStart=这一行配置，把fd://改成unix:///var/run/docker.sock
然后
sudo systemctl daemon-reload
sudo systemctl restart docker
查看
ll /var/run/*.sock 可以找到这个文件
```

# 制定docker的工作目录
修改/etc/docker/daemon.json文件，添加
```shell
{                                                                          
     "data-root": "/home/Userlist/docker/data"
}
```