```shell
curl -k --cert /etc/kubernetes/pki/apiserver-kubelete-client.crt --key /etc/kubenetes/pki/apiserver-kubelet-client.key  https://127.0.0.1:10250/api/v1/nodes/proxy/metrics/cadvisor
```