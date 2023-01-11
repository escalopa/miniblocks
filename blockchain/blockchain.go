package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

type Blockchain struct {
	lh []byte
	db badger.DB
}

const (
	dbPath = "./tmp/blocks"
)

func New() (*Blockchain, error) {
	lh := []byte("lh")

	// Open the database connnection
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		return nil, err
	}

	// Load lastHash from database or create genesis block
	err = db.Update(func(txn *badger.Txn) error {
		b, err := txn.Get(lh)
		if err != nil && err != badger.ErrKeyNotFound {
			return err
		}

		// Create the new genesis block and save it has the last hash in database
		if err == badger.ErrKeyNotFound {
			// Create new genesis block
			genesis := NewGenesis()
			data, err := genesis.Serialize()
			if err != nil {
				return err
			}
			// Save the genesis block, Where block.Hash >> block
			err = txn.Set(genesis.Hash, data)
			if err != nil {
				return err
			}
			// Save lh,
			err = txn.Set(lh, genesis.Hash)
			if err != nil {
				return err
			}
			lh = genesis.Hash
			return nil
		} else {
			// Set `lh` to the latest hash in the database
			err = b.Value(func(val []byte) error {
				lh = append([]byte{}, val...)
				// DONT DO THIS
				// lh = val
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &Blockchain{lh, *db}, nil
}

func (bc *Blockchain) AddBlock(data string) error {
	var lh []byte

	fmt.Println("Adding new block to the blockchain, lh = ", bc.lh)
	newBlock, err := NewBlock(data, bc.lh)
	if err != nil {
		return err
	}
	err = bc.db.Update(func(txn *badger.Txn) error {
		// Serialize the new block
		data, err := newBlock.Serialize()
		if err != nil {
			return err
		}
		// Save the new block, Where blcok.Hash >> block
		err = txn.Set(newBlock.Hash, data)
		if err != nil {
			return err
		}
		// Udpate the last hash
		err = txn.Set([]byte("lh"), newBlock.Hash)
		if err != nil {
			return err
		}
		lh = newBlock.Hash
		return nil
	})
	if err != nil {
		return err
	}
	// Update the last hash
	bc.lh = lh
	return nil
}

type BlockchainIterator struct {
	ch []byte
	db *badger.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.lh, &bc.db}
}

func (bci *BlockchainIterator) Next() (*Block, error) {
	var block *Block
	err := bci.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(bci.ch)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			_, err = block.Deserialize(val)
			return err
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	bci.ch = block.PrevHash
	return block, err
}
