先在目标目录下执行
```shell
go mod init seeplant.com/host-cluster-controller
```
然后执行
```shell
kubebuilder init --domain=""
```
然后执行
```shell
kubebuilder create api --group seeplant.com --version v1alpha1 --kind HostCluster
```
然后执行
```shell
make manifests
```
然后会在config/crd/bases目录下生成crd文件。crd的name是${kind}.${group}.${domain} domain是在第二步指定的。CRD的group是${group}.${domain}
<p>
如果要修改，就修改api/中的xxxtypexxx.go文件，改完执行make manifests
</p>
