在~/.bash_profile配置alias
```shell
alias etcdctl="ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/eetcd/healthcheck-client.crt --key=/etc/kubernetes/pki/etcd/healthcheck-client.key"
```

```shell
source ~/.bash_profile
etcdctl get / --prefix --keys-only | grep "/pods/"
```
