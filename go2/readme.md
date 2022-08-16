### 1、目前免杀情况
可过360、火绒、小红伞、微软、江民、瑞星等

### 2、使用方法

1)使用msf或cs生成shellcode

2)将shellcode从\xfc\x48\x83\xe4.........转化为0xfc,0x48,0x83,0xe4,.......格式（只需ctrl+h替换斜杠为0且末尾补一个英文逗号即可）

3)使用转换好的shellcode 替换shellcodeGen.go文件中的msg，执行go run shellcodeGen.go即可生成加密字符串

4)使用加密字符串替换AesexeGen.go文件中的text变量，运行go build AesexeGen.go生成exe


