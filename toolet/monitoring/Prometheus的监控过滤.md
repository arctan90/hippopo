> metric_relabel_configs是发生在抓取之后，但在数据被插入存储系统之前使用。因此，如果有些想过滤的指标，或者来着抓取本身的指标（比如来自/metrics页面）可以使用metric_relabel_configs来处理。

```yaml
metric_relabel_configs:
# 把__name__属性等于kubevirt_vmi_cpu_affinity的label扔掉
- source_labels:
  - __name__
  regex: kubevirt_vmi_cpu_affinity
  action: drop
# 把__name__属性满足正则的etcd_(debugging|disk|request|server).*这些label扔掉
- source_labels:
  - __name__
  regex: etcd_(debugging|disk|request|server).*
  action: drop
- source_labels:
    - __meta_kubernetes_endpoint_port_name
  action: keep
  regex: web
```