package main

import "fmt"

// 定义结构
// 前区块哈希
// 当前区块哈希
// 数据

// 创建区块
// 生成哈希
// 引入区块链
// 添加区块
// 重构代码

type Block struct {
	PrevBlockHash []byte //前一个hash
	Hash          []byte //当前hash
	Data          []byte //数据
}

// 创建区块 对Block的每一个字段填充数据
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
	}
	return &block
}

func main() {
	fmt.Printf("helloworld\n")
}
