```shell
kubectl config view --minify --raw --output 'jsonpath={.cluster.certificate-authority-data}' | base64 -d | openssl x509 - text -out -
```

# 关于证书
证书的认证方式有很多种
1. 通过token认真，token源自serviceAccount绑定role/clusterRole之后在对应的namespace下生成的secret里，它携带了服务的身份，apiserver对这个身份实施鉴权
2. 一种是用证书cert和key，证书中需要签对应的用户信息