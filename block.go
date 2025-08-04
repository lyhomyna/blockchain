package main

import (
	"time"
)

type Block struct {
    Timestamp		int64
    Data		[]byte
    PrevBlockHash	[]byte
    Hash		[]byte
    Nonce		int  // required to verify a proof
}

func NewGenesisBlock() *Block {
    return NewBlock("Genesis Block", []byte{})
}

func NewBlock(data string, prevBlockHash []byte) *Block {
    block := &Block {
	Timestamp: time.Now().Unix(),
	Data: []byte(data),
	PrevBlockHash: prevBlockHash,
    }

    pow := NewProofOfWork(block)
    nonce, hash := pow.Run()  // MINING!!1

    block.Hash = hash[:]
    block.Nonce = nonce

    return block
}
