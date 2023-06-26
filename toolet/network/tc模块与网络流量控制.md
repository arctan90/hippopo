tc qdisc add dev eth0 parent 10:1 handle 100:tbf rate 10mbit burst 32kbit latency 400ms
```shell
tc qdisc show dev eth0 可以查询eth0网卡的队列编号
```