package main

import (
	"bytes"
	"encoding/gob"
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

func (b *Block) Serialize() []byte {
    var result bytes.Buffer

    encoder := gob.NewEncoder(&result)
    encoder.Encode(b)

    return result.Bytes()
}

func DeserializeBlock(b []byte) *Block {
    var block Block

    decoder := gob.NewDecoder(bytes.NewReader(b))
    decoder.Decode(&block)

    return &block
}
