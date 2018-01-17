package Common

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

// 协议格式  4 byte 长度 ， 1 字节 类型 ，  1 字节版本 ， 8 字节 授权码 ，data[]byte json数据
type IMMessage struct {
	DataLen  uint32
	MsgType  uint8
	ProtoVer uint8
	AuthCode uint64
	Data     []byte
}

//转码
func (msg *IMMessage) Encode() []byte {
	totalLen := 4 + 1 + 1 + 8 + len(msg.Data)
	b1 := make([]byte, totalLen)
	copy(b1[0:4], Uint32toByte(msg.DataLen))
	b1[4] = byte(msg.MsgType)
	b1[5] = byte(msg.ProtoVer)
	copy(b1[5:5+8], Uint64ToByte(msg.AuthCode))
	copy(b1[5+8:totalLen], msg.Data)
	return b1
}

func (msg *IMMessage) Decode(data []byte) *IMMessage {
	dataLen := len(data)
	if dataLen <= 0 {
		fmt.Println("data is len empty IMMessage Decode")
		return nil
	} else {
		i1 := IMMessage{}
		dLen := BytetoUInt32(data[0:4])
		if (dLen > uint32(dataLen)) || (dLen < uint32(dataLen)) {
			fmt.Println("message package error , datalen is err")
			panic(error("message package error , datalen is err"))
		} else {
			i1.DataLen = dLen
			i1.MsgType = data[4]
			i1.ProtoVer = data[5]
			i1.AuthCode = ByteToUint64(data[6 : 6+8])
			i1.Data = data[14:]
		}
		return &i1
	}
}

//判断我们系统中的字节序类型 ,true , 大端  否则小端
func SystemEdian() bool {
	const INT_SIZE int = int(unsafe.Sizeof(0))
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

func Uint32toByte(num uint32) []byte {
	isLittle := SystemEdian()
	var b1 []byte = make([]byte, 4)
	if isLittle == true { // 大端
		binary.BigEndian.PutUint32(b1, uint32(num))
	} else {
		binary.LittleEndian.PutUint32(b1, uint32(num))
	}
	return b1
}

func Uint64ToByte(num uint64) []byte {
	isBig := SystemEdian()
	var b1 []byte = make([]byte, 8)
	if isBig {
		binary.BigEndian.PutUint64(b1, num)
	} else {
		binary.LittleEndian.PutUint64(b1, num)
	}
	return b1
}

func ByteToUint64(b1 []byte) uint64 {
	isLittle := SystemEdian()
	if isLittle {
		return binary.BigEndian.Uint64(b1)
	} else {
		return binary.LittleEndian.Uint64(b1)
	}
}

func BytetoUInt32(b1 []byte) uint32 {
	isLittle := SystemEdian()
	if isLittle {
		return binary.BigEndian.Uint32(b1)
	} else {
		return binary.LittleEndian.Uint32(b1)
	}
}
