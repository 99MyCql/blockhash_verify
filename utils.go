package main

import (
	"fmt"
	"math/big"
)

// litBytes2BigStr 小端字节数组转大端模式字符串
func litBytes2BigStr(b []byte) string {
	str := ""
	for i := len(b) - 1; i >= 0; i-- {
		str += fmt.Sprintf("%02X", b[i])
	}
	return str
}

// reverseCopy 大端字节数组转小端字节数组
func reverseCopy(dest []byte, src []byte) {
	destLen := len(dest)
	srcLen := len(src)
	for i := 0; i < destLen && i < srcLen; i++ {
		dest[i] = src[srcLen-i-1]
	}
}

// str2Bytes32 字符串转字节数组(小端)
func str2Bytes32(s string, base int) [32]byte {
	dest := [32]byte{}
	n := new(big.Int)
	n.SetString(s, base) // 大端
	reverseCopy(dest[:], n.Bytes())
	return dest
}

// uint32toBytes4 uint32转字节数组(小端)
func uint32toBytes4(u uint32) [4]byte {
	dest := [4]byte{}
	n := big.NewInt(int64(u)) // 大端
	reverseCopy(dest[:], n.Bytes())
	return dest
}
