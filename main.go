package main

import (
	"flag"
	"fmt"
)

var (
	blockIndex uint
)

func init() {
	flag.UintVar(&blockIndex, "height", 0, "区块号")
}

func main() {
	flag.Parse()
	fmt.Printf("验证第%d个区块中......\n", blockIndex)

	// 获取区块信息
	bData, err := getBlock(blockIndex)
	if err != nil {
		fmt.Println(err)
		return
	}
	bHeader := bData.getBlockHeader()
	fmt.Printf("区块头信息：\n%v\n", bHeader)
	bHash := bData.getBlockHash()
	fmt.Println("区块哈希值：" + litBytes2BigStr(bHash[:]))

	// 计算区块哈希值
	bHashCal, err := bHeader.calBlockHash()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("计算的区块哈希值：" + litBytes2BigStr(bHashCal[:]))

	// 验证
	if string(bHashCal[:]) != string(bHash[:]) {
		fmt.Println("验证失败！")
		return
	}
	fmt.Println("验证成功！")
}
