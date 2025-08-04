package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// difficulty of mining
// hash must start with { targetBits } zeroes
const targetBits = 24
const maxNonce = math.MaxInt64

type ProofOfWork struct {
    block *Block
    target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
    target := big.NewInt(1)
    target.Lsh(target, uint(256 - targetBits))

    pow := &ProofOfWork{b, target}

    return pow
}

// Run finds the nonce and hash that pass the difficulty and returns them
func (pow *ProofOfWork) Run() (int, []byte) {
    var hashInt big.Int
    var hash [32]byte
    nonce := 0

    fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
    for nonce < maxNonce {
	data := pow.prepareData(nonce)
	hash = sha256.Sum256(data)
	fmt.Printf("\r%x", hash)
	hashInt.SetBytes(hash[:])

	if hashInt.Cmp(pow.target) == -1 {
	    break
	} else {
	    nonce++
	}
    }
    fmt.Print("\n\n")

    return nonce, hash[:]
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
    data := bytes.Join([][]byte{
	pow.block.Data,
	pow.block.PrevBlockHash,
	IntToHex(pow.block.Timestamp),
	IntToHex(int64(nonce)),
	IntToHex(int64(targetBits)),
    }, []byte{})

    return data
}

// If you just compare pow.block.Hash to the target,
// you're trusting the block’s stored hash blindly. 
// That’s dumb and unsafe.
func (pow *ProofOfWork) Validate() bool {
    var hashInt big.Int
    
    data := pow.prepareData(pow.block.Nonce)
    hash := sha256.Sum256(data)
    hashInt.SetBytes(hash[:])
    
    return hashInt.Cmp(pow.target) == -1
}
