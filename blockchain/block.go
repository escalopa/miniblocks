package blockchain

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"
)

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

func (b *Block) Serialize() ([]byte, error) {
	var result bytes.Buffer
	err := gob.NewEncoder(&result).Encode(b)
	if err != nil {
		return []byte{}, err
	}
	return result.Bytes(), nil
}

func (b *Block) Deserialize(data []byte) (*Block, error) {
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ToInt(b []byte) int {
	var result int
	binary.Read(bytes.NewReader(b), binary.BigEndian, &result)
	return result
}
