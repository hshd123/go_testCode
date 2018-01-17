package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

func main() {

	l1 := uint32(12873)
	l2 := uint8(12)
	l3 := uint8(2)

	s := "这是一个字符串"
	sLen := len(s)

	b3 := toByte(uint32(l1))
	buff := make([]byte, 6+sLen)
	copy(buff[0:4], b3)
	buff[4] = byte(l2)
	buff[5] = byte(l3)
	copy(buff[6:sLen+6], []byte(s))

	fmt.Println(binary.BigEndian.Uint32(buff[0:4]))
	fmt.Println("bufflen - ", len(buff), "\n", buff)
	fmt.Println(uint8(buff[4]))
	fmt.Println(uint8(buff[5]))
	fmt.Println(string(buff[6 : sLen+6]))
	//
	//
	//total := 1 + 1 + 4 + len([]byte(s))
	//
	//buff := make([]byte, total)
	//copy(buff[0:3], b3)
	//buff[4] = byte(l2)
	//buff[5] = byte(l3)
	//copy(buff[6:total], []byte(s))
	//
	//fmt.Println("bytes -- ", buff, " --- ", len(buff), len([]byte(s)), total)
	//fmt.Println("to int 32 -- ", buff[0:3])
	//fmt.Println("int8 --", buff[4])
	//fmt.Println("int 8 ", buff[5])
	//fmt.Println(string(buff[6:total]))
	//fmt.Println(l1)

	//fmt.Println("-----------")
	////[0 1 245 138]
	//var a []byte = []byte{0, 1, 245, 138}
	//fmt.Println(a)
	//fmt.Println(binary.BigEndian.Uint32(a))
	////fmt.Println(binary.LittleEndian.Uint32(a))
	//
	//s1 := "012345678"
	////fmt.Println(" ", s1[0:1])
	////fmt.Println(s1[1 : len(s1)])
	//fmt.Println(string(s1[0:3]))
	//fmt.Println("", string(s1[1:3]))
	//fmt.Println("", string(s1[4]))

}

const INT_SIZE int = int(unsafe.Sizeof(0))

//判断我们系统中的字节序类型 ,true , 大端  否则小端
func systemEdian() bool {
	var i int = 0x1
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 {
		//fmt.Println("system edian is little endian")
		return false
	} else {
		//fmt.Println("system edian is big endian")
		return true
	}
}

func toByte(num uint32) []byte {
	isLittle := systemEdian()
	var b1 []byte = make([]byte, 4)
	if isLittle == true { // 大端
		binary.BigEndian.PutUint32(b1, uint32(num))
		fmt.Println(b1)
	} else {
		binary.LittleEndian.PutUint32(b1, uint32(num))
	}
	return b1
}

func toInt32(b1 []byte) uint32 {
	fmt.Println(b1)
	isLittle := systemEdian()
	if isLittle {
		return binary.BigEndian.Uint32(b1)
	} else {
		return binary.LittleEndian.Uint32(b1)
	}
}
