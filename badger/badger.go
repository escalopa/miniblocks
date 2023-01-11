package badger

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/escalopa/myblocks/pkg"
)

type BadgerDB struct {
	db *badger.DB
	lh []byte
}

// New creates a new BadgerDB instance
// @param path - The path to the database, if not db exists, it will be created, But the directory must exist
// @return *BadgerDB - The new BadgerDB instance
// @return error - Any error that occurs
func New(path string) (*BadgerDB, error) {
	opts := badger.DefaultOptions(path)
	opts.Logger = nil
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &BadgerDB{db: db, lh: []byte("lh")}, nil
}

// SaveBlock saves a block to the database
// Also updates the last hash value
// @param block - The block to save
// @return error - Any error that occurs
func (b *BadgerDB) SaveBlock(block *pkg.Block) error {
	err := b.db.Update(func(txn *badger.Txn) error {
		// Serialize the new block
		data, err := block.Serialize()
		if err != nil {
			return err
		}
		// Save the new block, Where blcok.Hash >> block
		err = txn.Set(block.Hash, data)
		if err != nil {
			return err
		}
		// Udpate the last hash
		err = txn.Set(b.lh, block.Hash)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetBlock returns a block from the database
// @param hash - The hash of the block to get
// @return block - The block
// @return error - Any error that occurs
func (b *BadgerDB) GetBlock(hash []byte) (*pkg.Block, error) {
	block := &pkg.Block{}
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(hash)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			data := append([]byte{}, val...)
			err = block.Deserialize(data)
			return err
		})
		return err
	})
	return block, err
}

func (b *BadgerDB) GetLastHash() ([]byte, error) {
	var lh []byte
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(b.lh)
		if err == badger.ErrKeyNotFound {
			return pkg.ErrKeyNotFound
		}
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			lh = append([]byte{}, val...)
			return nil
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return lh, err
}

// Close closes the database
// @return error - Any error that occurs
func (b *BadgerDB) Close() error {
	return b.db.Close()
}
