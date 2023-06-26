> 有些项目会发布一些helm chart，如果想从repo里把chart包拖下来，需要做一些操作
```shell
helm repo add xxxx https://xxxx
helm repo update
helm pull xxxx/xxxxxx
```