package main

import (
	"log"

	"github.com/escalopa/myblocks/badger"
	"github.com/escalopa/myblocks/blockchain"
)

const (
	dbPath = "./tmp/blocks"
)

func main() {
	var err error

	// Init Database
	db, err := badger.New(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bc, err := blockchain.New(db)
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
