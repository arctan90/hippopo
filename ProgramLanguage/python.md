# Mac上 OSX pip install - error: the clang compiler does not support '-march=native'
尝试添加 `ARCHFLAGS="-arch x86_64" pip3 install ...`
如果是M1 添加 `ARCHFLAGS="-arch arm64"`