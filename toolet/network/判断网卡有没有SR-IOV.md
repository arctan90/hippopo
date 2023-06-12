```shell
lspci | grep -i Ethernet 查到id是18:00.0的话
再看是否支持 lspci -s da:00.0 -v 有 有SR-IOV字样就是支持VF
```