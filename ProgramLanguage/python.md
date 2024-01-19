# Mac上 OSX pip install - error: the clang compiler does not support '-march=native'
尝试添加 `ARCHFLAGS="-arch x86_64" pip3 install ...`
如果是M1 添加 `ARCHFLAGS="-arch arm64"`

# requirements.txt
`pip3 install pipreqs`
设置alias
`alias pipreqs='python3 -m pipreqs.pipreqs'`
导出
`pipreqs ./ --encoding=utf-8 --force`
自动导入
`pip install -r requirements.txt`

# pip安装组件的目录
`pip show xxx`
比如
`pip show -f pipreqs | grep Location:`

# 开代理的时候无法pip
`ERROR: Could not install packages due to an EnvironmentError: Missing dependencies for SOCKS support.`
先在zshrc或者.bashrc里把代理all_proxy注释掉，然后重进bash或zsh，然后装个pysocks `pip install pysocks` ，然后反注释all_proxy，重进bash即可

# conda无法create
先用`python3 -c "import requests; print(requests.__version__)"`查下版本，看是不是太低了
然后用`pip3 install requests --upgrade`更新一下requests