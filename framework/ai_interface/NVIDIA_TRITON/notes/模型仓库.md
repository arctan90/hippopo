# 概述
在Triton服务启动后，可以指定用多个模型仓库提供的模型来提供推理服务。（可以动态加载仓库中的模型）。在Triton运行的时候模型可以按模型管理指定的方
式进行修改。
# 仓库布局
在Triton启动的时候，需要通过参数```–model-repository```参数指定仓库的路径。```–model-repository```选项可以多次使用来包含多个仓库。组成
仓库的路径和文件需要参照下述的规则布局。
```shell
$ tritonserver --model-repository=<model-repository-path>
```
```text
  <model-repository-path>/
    <model-name>/
      [config.pbtxt]
      [<output-labels-file> ...]
      <version>/
        <model-definition-file>
      <version>/
        <model-definition-file>
      ...
    <model-name>/
      [config.pbtxt]
      [<output-labels-file> ...]
      <version>/
        <model-definition-file>
      <version>/
        <model-definition-file>
      ...
    ...
```
上面是一个分层在文件夹结构。可以有0到多个model-name的文件夹，每一个代表一个模型。config.pbtxt文件不是必须的，在模型配置章节会给出描述。
每个模型至少需要一个version目录。每个模型都由特定的后端来执行（如何绑定执行后端？）。
# 模型仓库位置
Triton可以通过多个本地、云服务、S3、Azure Storage来访问模型（足够了，我有S3）
## 本地文件方式
需要指定<b><font color=RED>绝对路径</font></b>，比如
```shell
$ tritonserver --model-repository=/path/to/model/repository ...
```
## 使用环境变量的云存储
### google存储
使用指定的schema, gs://
```shell
$ tritonserver --model-repository=gs://bucket/path/to/model/repository ...
```
### S3
使用指定的schema， s3://
```shell
$ tritonserver --model-repository=s3://bucket/path/to/model/repository ...
```
如果是本地自己搭建的s3，需要使用```地址+":"+端口号```
```shell
$ tritonserver --model-repository=s3://host:port/bucket/path/to/model/repository ...
```
使用S3的时候，证书和default region可以通过aws config，或者环境变量来指定，其中环境变量的优先级更高。
### Azure Storage
使用前缀 as://
```shell
$ tritonserver --model-repository=as://account_name/container_name/path/to/model/repository ...
```
此时需要使用环境变量 <font color=Red>AZURE_STORAGE_ACCOUNT</font>和<font color=Red>AZURE_STORAGE_KEY</font>环境变量作为访问账号的设置。
## 云存储的认证文件
<I>实验特性</I>
通过环境变量来设置使用的文件
```shell
export TRITON_CLOUD_CREDENTIAL_PATH="cloud_credential.json"
```
认证文件的格式
```shell
{
  "gs": {
    "": "PATH_TO_GOOGLE_APPLICATION_CREDENTIALS",
    "gs://gcs-bucket-002": "PATH_TO_GOOGLE_APPLICATION_CREDENTIALS_2"
  },
  "s3": {
    "": {
      "secret_key": "AWS_SECRET_ACCESS_KEY",
      "key_id": "AWS_ACCESS_KEY_ID",
      "region": "AWS_DEFAULT_REGION",
      "session_token": "",
      "profile": ""
    },
    "s3://s3-bucket-002": {
      "secret_key": "AWS_SECRET_ACCESS_KEY_2",
      "key_id": "AWS_ACCESS_KEY_ID_2",
      "region": "AWS_DEFAULT_REGION_2",
      "session_token": "AWS_SESSION_TOKEN_2",
      "profile": "AWS_PROFILE_2"
    }
  },
  "as": {
    "": {
      "account_str": "AZURE_STORAGE_ACCOUNT",
      "account_key": "AZURE_STORAGE_KEY"
    },
    "as://Account-002/Container": {
      "account_str": "",
      "account_key": ""
    }
  }
}
```
# 模型版本
首先，每个版本都存储在以数字命名的子目录中，子目录的名称对应于模型的版本号。未以数字命 名或名称以零 (0) 开头的子目录将被忽略。
每个模型配置都指定了一个版本策略，该策略控制模型存储库中的哪些版本在任何给定时间由 Triton 提供。
# 模型文件
模型子目录的内容由 <br><I>模型的类型</I></br> 和 <br><I>使用模型的后端</I></br> 决定。
## TensorRT模型
TensorRT模型定义叫做<I>规划</I>。一个规划是一个单独文件默认必须命名wei```model.plan```。这个默认名字可以用模型定义中的```default_model_filename```来覆盖。
<br>
一个规划指定到特定的CUDA计算能力。因此，TensorRT模型需要在模型中设置```cc_model_filenames```把计算能力和规划文件联系起来。
<br>
一个最小的TensorRT模型为
```shell
  <model-repository-path>/
    <model-name>/
      config.pbtxt
      1/
        model.plan
```
## ONNX(开放神经网络)模型
一个ONNX模型是一个单独文件或者一个包含多文件的目录。默认的这个单独文件或者文件夹应该叫做model.onnx。这个默认行为可以被模型定义中的```default_model_filename```所覆盖
<br>
版本支持方面，过旧的版本不支持。
<br>
最简的单文件模型定义如下
```shell
  <model-repository-path>/
    <model-name>/
      config.pbtxt
      1/
        model.onnx
```
最简的的多文件模型定义如下
```shell
  <model-repository-path>/
    <model-name>/
      config.pbtxt
      1/
        model.onnx/
           model.onnx
           <other model files>
```
## TorchScript 模型
TorchScript模型是单文件，默认名称为model.pt。这个默认名称可以在模型定义的```default_model_filename```来覆盖。
最简的的模型定义
```shell
  <model-repository-path>/
    <model-name>/
      config.pbtxt
      1/
        model.pt
```
## TensorFlow模型
TensorFlow可以将模型保存为两种格式。<I>GraphDef</I>或者<I>SavedModel</I>，Triton对两种格式都支持。
<br>
GraphDef是一个单文件模型叫做model.graphdef。一个tensorflow saveModel文件夹包含多个文件，这个文件夹默认叫model.savedmodel。GraphDef最小配置为
```shell
  <model-repository-path>/
    <model-name>/
      config.pbtxt
      1/
        model.graphdef
```
或者
```shell
  <model-repository-path>/
    <model-name>/
      config.pbtxt
      1/
        model.savedmodel/
           <saved-model files>
```
## OpenVINO 模型
OpenVINO模型需要两个文件叫做 *.xml和*.bin文件。默认model.xml。最小配置为
```shell
  <model-repository-path>/
    <model-name>/
      config.pbtxt
      1/
        model.xml
        model.bin
```
## Python模型
Python后端允许在triton上运行python脚本。Python默认脚本需要命名为model.py，但是可以用模型配置中的default_model_filename覆盖。最小的模型配置
```shell
  <model-repository-path>/
    <model-name>/
      config.pbtxt
      1/
        model.py
```
## DALI模型
DALI后端运行运行DALI流水线。为了使用这个后端，你需要生成一个文件，默认名字为model.dali。最小模型定义为
```shell
  <model-repository-path>/
    <model-name>/
      config.pbtxt
      1/
        model.dali
```