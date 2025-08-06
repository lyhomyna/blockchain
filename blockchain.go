package main

import "github.com/boltdb/bolt"

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type Blockchain struct {
    tip []byte
    db *bolt.DB
}

func NewBlockchain() *Blockchain {
    var tip []byte
    db, _ := bolt.Open(dbFile, 0600, nil)

    db.Update(func(tx *bolt.Tx) error {
	b := tx.Bucket([]byte(blocksBucket))
    
	if b == nil {
	    genesis := NewGenesisBlock()
	    b, _ := tx.CreateBucket([]byte(blocksBucket))
	    b.Put(genesis.Hash, genesis.Serialize())
	    b.Put([]byte("l"), genesis.Hash)
	    tip = genesis.Hash
	} else {
	    tip = b.Get([]byte("l"))
	}

	return nil
    })

    bc := Blockchain {tip, db}

    return &bc
}

func (bc *Blockchain) AddBlock(data string) {
    var lastHash []byte

    bc.db.View(func(tx *bolt.Tx) error { 
	b := tx.Bucket([]byte(blocksBucket))
	lastHash = b.Get([]byte("l"))
	return nil
    })

    newBlock := NewBlock(data, lastHash)

    bc.db.Update(func(tx *bolt.Tx) error {
	b := tx.Bucket([]byte(blocksBucket))

	b.Put(newBlock.Hash, newBlock.Serialize())
	b.Put([]byte("l"), newBlock.Hash)

	bc.tip = newBlock.Hash

	return nil
    })
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
    return &BlockchainIterator{ bc.tip, bc.db }
}
