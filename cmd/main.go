package main

import (
	"log"
	"os"

	"github.com/escalopa/myblocks/badger"
	"github.com/escalopa/myblocks/blockchain"
	"github.com/escalopa/myblocks/cobra"
)

const (
	dbPath = "./tmp/blocks"
)

func main() {
	var err error

	defer os.Exit(0)
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

	cmd, err := cobra.New(bc)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Run()
}
