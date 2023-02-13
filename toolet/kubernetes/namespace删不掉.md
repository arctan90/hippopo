1. 检查一下apiservice状态，常见的是扩展api对应的instance没了，阻塞住了便利资源
2. 用下面命令便利namespace中的资源
```shell
kubectl api-resources --verbs=list --namespaced -o name | xargs kubectl get -n$NAMESPACE_NAME
```
