package blockchain

import "errors"

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func NewGenesis() *Block {
	b, err := NewBlock("Genesis", []byte{})
	if err != nil {
		panic("Create genesis block failed")
	}
	return b
}

func NewBlock(data string, prevHash []byte) (*Block, error) {
	b := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(b)
	nonce, hash := pow.Run()
	b.Hash = hash[:]
	b.Nonce = nonce
	if !pow.Validate() {
		return nil, errors.New("PoW validation failed")
	}
	return b, nil
}
