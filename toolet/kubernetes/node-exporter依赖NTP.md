如果使用的是chrony，默认的配置未对本机开启访问权限，需要进行如下操作
```shell
vi /etc/chrony.conf
# 加allow 127/8
systemctl restart chronyd
```