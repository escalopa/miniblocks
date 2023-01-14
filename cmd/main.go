package main

import (
	"log"
	"os"

	"github.com/escalopa/miniblocks/badger"
	"github.com/escalopa/miniblocks/blockchain"
	"github.com/escalopa/miniblocks/cobra"
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
