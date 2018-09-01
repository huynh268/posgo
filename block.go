package main

import "time"

type Block struct {
	Timestamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
	Validator Validator
}

func CreateBlock(data []byte, prevHash []byte, validator []byte) *Block {
	block := &Block{time.Now().Unix(), data, prevHash, []byte{}}

	return block
}
