package main

import (
	"fmt"
	"log"

	"github.com/escalopa/myblocks/blockchain"
)

func main() {
	var err error
	bc := blockchain.New()
	err = bc.AddBlock("First Block")
	if err != nil {
		log.Fatal(err)
	}
	err = bc.AddBlock("Second Block")
	if err != nil {
		log.Fatal(err)
	}
	err = bc.AddBlock("Third Block")
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range bc.Blocks {
		fmt.Printf("BlochHash: %x\n", b.Hash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("PreHash: %x\n", b.PrevHash)
		fmt.Println("")
	}
}
