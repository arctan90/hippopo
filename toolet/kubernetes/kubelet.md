# kubelet中cadvisor的路径
```shell
curl -k --cert /etc/kubernetes/pki/apiserver-kubelete-client.crt --key /etc/kubenetes/pki/apiserver-kubelet-client.key  https://127.0.0.1:10250/api/v1/nodes/proxy/metrics/cadvisor
```

# kubelet的api
https://www.deepnetwork.com/blog/2020/01/13/kubelet-api.html#pods