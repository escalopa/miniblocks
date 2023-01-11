package blockchain

import (
	"github.com/escalopa/myblocks/application"
	"github.com/escalopa/myblocks/pkg"
)

type BlockchainIterator struct {
	ch []byte
	db application.IDatabase
}

func (bc *Blockchain) NewIterator() (*BlockchainIterator, error) {
	lh, err := bc.db.GetLastHash()
	if err != nil {
		return nil, err
	}
	return &BlockchainIterator{lh, bc.db}, nil
}

func (bci *BlockchainIterator) Next() (block *pkg.Block, err error) {
	block, err = bci.db.GetBlock(bci.ch)
	if err != nil {
		return nil, err
	}
	bci.ch = block.PrevHash
	return block, err
}
