package block

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func New(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func NewGenesis() *Block {
	return New("Genesis", []byte{})
}

func (b *Block) DeriveHash() {
	hash := sha256.Sum256(bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{}))
	b.Hash = hash[:]
}
