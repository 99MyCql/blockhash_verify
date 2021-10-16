package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
)

type blockHeader struct {
	version       [4]byte  // 版本号
	prevBlockHash [32]byte // 前一个区块的哈希
	merkleRoot    [32]byte // 该区块中交易的merkle树根的哈希值
	time          [4]byte  // 该区块的创建时间戳
	bits          [4]byte  // 该区块工作量证明难度目标
	nonce         [4]byte  // 用于证明工作量的计算随机数
}

// toBytes 将区块头转换成字节数组(小端)
func (b *blockHeader) toBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, b); err != nil {
		return nil, errors.Wrap(err, "区块头转字节数组出错")
	}
	return buf.Bytes(), nil
}

// calBlockHash 计算区块头的哈希值
func (b *blockHeader) calBlockHash() ([32]byte, error) {
	bhBytes, err := b.toBytes()
	if err != nil {
		return [32]byte{}, errors.WithMessage(err, "")
	}
	fmt.Printf("区块头字节序(小端)：%080X\n", bhBytes)

	h1 := sha256.Sum256(bhBytes)
	return sha256.Sum256(h1[:]), nil
}

// String 打印区块头信息(小端字节数组需转换成大端打印)
func (b *blockHeader) String() string {
	str := "+++++++++++++++++++++++++++++++"
	str += "\nversion:" + litBytes2BigStr(b.version[:])
	str += "\nprevBlockHash:" + litBytes2BigStr(b.prevBlockHash[:])
	str += "\nmerkleRoot:" + litBytes2BigStr(b.merkleRoot[:])
	str += "\ntime:" + litBytes2BigStr(b.time[:])
	str += "\nbits:" + litBytes2BigStr(b.bits[:])
	str += "\nnonce:" + litBytes2BigStr(b.nonce[:])
	str += "\n+++++++++++++++++++++++++++++++"
	return str
}
