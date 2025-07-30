package main

import (
	"fmt"
	"time"
)

func main() {

}

func test() {
    bc := NewBlockchain()

    bc.AddBlock("Sent 1 BTC to Misha")
    bc.AddBlock("Sent 1 BTC to James")

    for _, block := range(bc.blocks) {
	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	fmt.Println("Created: ", time.Unix(block.Timestamp, 0))
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Println()
    }
}
