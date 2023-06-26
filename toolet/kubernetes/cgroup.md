#工作目录
/sys/fs/cgroup
#挂载方式
mount -t cgroup -o cpu,cpuset,memory cpu_and_mem /cgroup/cpu_and_mem
在/cgroup/cpu_and_memory目录下，加载cpu cpuset memory三个子模块，在/cgroup/cpu_and_memory目录下生成三个队友的子目录