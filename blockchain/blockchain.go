package blockchain

import (
	"errors"

	"github.com/escalopa/myblocks/application"
	"github.com/escalopa/myblocks/pkg"
)

type Blockchain struct {
	db application.IDatabase
}

func New(db application.IDatabase) (*Blockchain, error) {
	bc := &Blockchain{db}

	_, err := bc.db.GetLastHash()
	if err == pkg.ErrKeyNotFound {
		genesis, err := newGenesis()
		if err != nil {
			return nil, err
		}
		err = bc.db.SaveBlock(genesis)
		if err != nil {
			return nil, err
		}
	}
	return bc, nil
}

func (bc *Blockchain) AddBlock(data string) error {
	lh, err := bc.db.GetLastHash()
	if err != nil {
		return err
	}
	newBlock, err := newBlock(data, lh)
	if err != nil {
		return err
	}
	err = bc.db.SaveBlock(newBlock)
	if err != nil {
		return err
	}
	return nil
}

// newGenesis creates a new genesis block with empty prevHash
// @return *Block - the new block
// @return error - error if any
func newGenesis() (*pkg.Block, error) {
	b, err := newBlock("Genesis", []byte{})
	if err != nil {
		return nil, err
	}
	return b, nil
}

// newBlock creates a new block using PoW
// @param data - data to be stored in the block
// @param prevHash - hash of the previous block
// @return *Block - the new block
// @return error - error if any
func newBlock(data string, prevHash []byte) (*pkg.Block, error) {
	b := &pkg.Block{Data: []byte(data), PrevHash: prevHash}
	pow := NewProof(b)
	nonce, hash := pow.Run()
	b.Hash = hash[:]
	b.Nonce = nonce
	if !pow.Validate() {
		return nil, errors.New("PoW validation failed")
	}
	return b, nil
}

// // toInt converts a byte slice to an int
// // @param b - byte slice
// // @return int - the int value
// func toInt(b []byte) int {
// 	var result int
// 	binary.Read(bytes.NewReader(b), binary.BigEndian, &result)
// 	return result
// }
