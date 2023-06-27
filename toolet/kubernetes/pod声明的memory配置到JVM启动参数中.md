# 解决思路
使用一个脚本作为容器的启动参数
```shell
command: ["/bin/bash"]
args:
- "-c"
- |
  bash "/home/admin/start.sh"
```
## 脚本
```shell
# 先获取cgroup管理的资源的值
if [ ! -d /etc/podinfo ]; then
  # 非k8s，而是docker的方式
  memtotal_mb=$[${cat /sys/fs/cgroup/memory/memory.limit_in_bytes} / 1024 / 1024]
else
  memtotal_mb=$(cat /etc/podinfo/mem_limit)
fi 

if [ $memtotal_mb -ge $[64 * 1024] ];then
  echo "pod memory resource limit[memtotal_mb mb] is too large, plz check if the limit value is set"
  exit 1
fi
date=$(date +%Y%m%d)
#假定gc日子的目录
gc_log_path=/home/admin/app/logs/gc
mkdir -p $gc_log_path

# 动态计算jvm启动参数, 单位是MBytes
xmx=`echo "$memtotal_mb 0.75" | awk '{printf("%g", $1*$2)}'`
xms="$[$xmx / 1]"
xmn=`echo "xmx 0.33" | awk '{printf("%g", $1*$2)}'`
xss="1" # 这里是否需要动态绑定看项目实际情况，根据实践经验暂时是正常的
maxDirectMemorySize=`echo "$memtotal_mb 0.1" | awk '{printf("%d", $1*$2)}'`

# 具体参数信息, 根据实际情况预先定义
# G1ConcRefinementThreads, ConcGCThreads, parallelGCThreads, ParallelGCThreads应该根据cpu核数计算
CATALINA_OPTS="$CATALINA_OPTS -Xms${xms}m -Xmx${xmx}m -Xss${xss}m -XX:MaxDirectMemorySize=${maxDirectMemorySize} \
-XX:+UnlockExperimentalVMOptions -XX:+UnlockDiagnosticVMOptions \
-XX:+UseG1GC -XXG1RSetUpdatingPauseTimePercent=10 -XX:G1ConcRefinementThreads=4 -XX:MaxGCPauseMillis=50 \
-XX:G1HeapWastePercent=10 -XX:G1ReservePercent=10 -XX:G1OldCSetRegionThresholdPercent=30 -XX:+MonitorInUseLists \
-XX:InitiationHeapOccupancyPercent=70 -XX:G1MixedGCLiveThresholdPercent=75 -XX:+G1EagerReclaimHumongousObjects \
-XX:ConcGCThreads=4 -XX:parallelGCThreads=8 \
-XX:-ResizePLAB -XX:-ParallelGCThreads=8 \
-XX:+PrintGCDetails -XX:+PrintGCDteStamps -XX:+PrintGCApplicationStoppedTime \
-XX:+ExitOnOutOfMemoryError -XX:+HeapDumpOnOutOfMemoryError -XX:+PrintFlagsFinal -XX:+PrintReferenceGC \
-XX:HeapDumpPath=/home/admin/app/logs/jvm/ -Xloggc:/home/admin/app/logs/gc/gc-$date.log"

# 重新构建启动脚本
startShell="$JAVA_HOME/bin/java -jar $CATALINA_OPTS -Ojava.security.egd=file:/dev/./urandom \
-Dxxx=xxx  '/home/admin/app/app.jar' \
--server.xxx=xxx"

# 覆盖启动脚本
echo $startShell > /home/adin/java_start.sh

# 增加最大进程数
echo > /etc/security/limits.conf
if [ -d /etc/security/limits.d ];then
  rm -rf /etc/security/limits.d/*
fi
if ! group "options use-vc" /etc/resolv.conf>/dev/null; then
    echo "options use-vc" >> /etc/resolv.conf
fi
# 设置启动用户权限
id -u admin &>/dev/null || useradd admin;
chown admin:admin /homt/admin -R;
su - admin -s /bin/sh /home/admin/java_start.sh
```