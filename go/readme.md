### 1、目前免杀情况
实机测试可过360、火绒

沙箱可过360、火绒、小红伞、微软、江民、瑞星等--2022.8.15

### 2、使用方法

1)使用msf或cs生成shellcode

2)将shellcode从\xfc\x48\x83\xe4.........转化为0xfc,0x48,0x83,0xe4,.......格式（只需ctrl+h替换斜杠为0且末尾补一个英文逗号即可）

3)使用转换好的shellcode 替换shellcodeGen.go文件中的msg，修改自己想使用的密码，执行go run shellcodeGen.go即可生成加密字符串

4)使用加密字符串替换exeGen.go文件中的text变量，运行go build exeGen.go生成exe，使用.\exeGen.exe passwd运行(passwd为自己设置的密码)


