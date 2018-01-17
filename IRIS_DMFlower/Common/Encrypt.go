package Common

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

type Encrypt interface {
}

type crypt struct{}

func init() {
}

func Base64Encode(s string) (string, error) {
	if l := len(s); l <= 0 {
		fmt.Println("Base64Encode error string is nil")
	}
	b := base64.StdEncoding.EncodeToString([]byte(s))
	return string(b), nil
}

func Base64Decode(s string) (string, error) {
	if l := len(s); l <= 0 {
		fmt.Println("error s is empty str")
	}
	b, err := base64.StdEncoding.DecodeString(s)
	return string(b), err
}

// key 必须是8位
func DesEncrypt(sec string, key string) (string, error) {
	if secLen := len(sec); secLen <= 0 {
		panic("DesEncrypt sec is error")
	}
	var tempKey string = key
	keyLen := len(key)
	if keyLen < 8 {
		tempKey = "12345678"
	}
	if keyLen > 8 {
		tempKey = tempKey[0:8]
	}
	resByte, err := desEnc([]byte(sec), []byte(tempKey))
	res := string(base64.StdEncoding.EncodeToString(resByte))
	return res, err
}

//key 必须是8 位
func DesDecrypt(secStr string, key string) (string, error) {
	fmt.Println(key)
	if secLen := len(secStr); secLen <= 0 {
		panic("DesDecrypt sec is empty")
	}
	var tempKey string = key
	keyLen := len(key)
	if keyLen < 8 {
		tempKey = "12345678"
	}
	if keyLen > 8 {
		tempKey = tempKey[0:8]
	}
	resByte, err := desDec([]byte(secStr), []byte(tempKey))
	ret := string(resByte)
	fmt.Println("ret ", ret)
	return ret, err
}

//key 必须是32位
func TripleDesEncrypt(sec string, key string) (string, error) {
	if secLen := len(sec); secLen <= 0 {
		panic("TripleDesEncrypt sec is empty")
	}
	var tempKey string = key
	keyLen := len(key)
	if keyLen < 32 {
		for i := 0; i < 32-keyLen; i++ {
			tempKey += "0"
		}
	}
	if keyLen > 32 {
		tempKey = tempKey[0:32]
	}
	resByte, err := tripleDesEncrypt([]byte(sec), []byte(tempKey))
	resStr := base64.StdEncoding.EncodeToString(resByte)
	return resStr, err
}

func TripleDesDecrypt(sec string, key string) (string, error) {
	if secLen := len(sec); secLen <= 0 {
		panic("TripleDesDecrypt sec is empty")
	}
	var tempKey string = key
	keyLen := len(key)
	if keyLen < 32 {
		for i := 0; i < 32-keyLen; i++ {
			tempKey += "0"
		}
	}
	if keyLen > 32 {
		tempKey = tempKey[0:32]
	}
	resByte, err := tripleDesDecrypt([]byte(sec), []byte(tempKey))
	str := string(resByte)
	return str, err
}

func desEnc(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = pkcs5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func desDec(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

// 3DES加密
func tripleDesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = pkcs5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 3DES解密
func tripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func zeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func MD5(s string) string {
	if l := len(s); l <= 0 {
		fmt.Println("MD5 s is empty")
		panic("MD5 s is empty")
	}
	s1 := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	return s1
}

func MD5FromFile(path string) string {
	if state := PathExist(path); state == true {
		f, err := os.Open("file.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		s := fmt.Sprint("%x", h.Sum(nil))
		return s
	}
	return ""
}

func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func Sha512(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	str := fmt.Sprint("%x", h.Sum(nil))
	return str
}
