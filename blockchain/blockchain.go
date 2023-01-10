package blockchain

import (
	"bytes"
	"errors"
)

type Blockchain struct {
	Blocks []*Block
}

func New() *Blockchain {
	bc := &Blockchain{Blocks: []*Block{NewGenesis()}}
	return bc
}

func (bc *Blockchain) AddBlock(data string) error {
	prevHash := bc.Blocks[len(bc.Blocks)-1].Hash
	newBlock, err := NewBlock(data, prevHash)
	if err != nil {
		return err
	}
	bc.Blocks = append(bc.Blocks, newBlock)
	if !bc.ValidateState() {
		// Drop the latest block
		bc.Blocks = bc.Blocks[:len(bc.Blocks)-1]
		return errors.New("invalid prevHash sequence")
	}
	return nil
}

func (bc *Blockchain) ValidateState() bool {
	prevBlock := bc.Blocks[len(bc.Blocks)-2]
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	return bytes.Equal(prevBlock.Hash, lastBlock.PrevHash)
}
