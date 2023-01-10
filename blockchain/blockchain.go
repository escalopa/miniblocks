package blockchain

import (
	"errors"
	"reflect"

	"github.com/escalopa/myblocks/block"
)

type Blockchain struct {
	Blocks []*block.Block
}

func New() *Blockchain {
	return &Blockchain{[]*block.Block{block.NewGenesis()}}
}

func (bc *Blockchain) AddBlock(data string) error {
	newBlock := block.New(data, bc.Blocks[len(bc.Blocks)-1].Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
	if err := bc.IsValidBlock(newBlock); err != nil {
		return errors.New("Invalid Block")
	}
	return nil
}

func (bc *Blockchain) IsValidBlock(b *block.Block) error {
	// Compare `b.PrevHash` with the latest hash in the blockchain
	if reflect.DeepEqual(b.PrevHash, bc.Blocks[len(bc.Blocks)-1].Hash) {
		return errors.New("Block.PrevHash != Blockchain latest block hash")
	}
	return nil
}
