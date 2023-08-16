# 查询显卡型号
思路：通过pci编号查询。
```shell
#先看网卡设备的类型
lspci | grep -i vga
#如果只是pci的数字，可以通过一个网站查
curl http://pci-ids.ucw.cz/mods/PC/10de/2204 | grep itemname
# 其中10de是nvidia的vendor id，2204是pci编号
``` 