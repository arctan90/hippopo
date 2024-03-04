# 下载模型
[下载说明](https://huggingface.co/docs/huggingface_hub/v0.20.2/guides/download)

直接下载全部模型
```pyton
from huggingface_hub import snapshot_download
snapshot_download(repo_id="lysandre/arxiv-nlp",local_dir="")
```

只下载数据集
```python
from huggingface_hub import snapshot_download
snapshot_download(repo_id="google/fleurs", repo_type="dataset")
```

