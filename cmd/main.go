package main

import (
	"log"

	"github.com/escalopa/myblocks/blockchain"
)

func main() {
	var err error
	bc, err := blockchain.New()
	if err != nil {
		log.Fatal(err)
	}
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
}
