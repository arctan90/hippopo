kubectl proxy的作用是通过kubectl对外暴露访问地址，来无密的用http请求访问k8s的rest接口
```shell
kubectl proxy --address='0.0.0.0' --accept-hosts='^*$' --port=8889
```