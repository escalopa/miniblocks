package pkg

import (
	"bytes"
	"encoding/gob"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Serialize serializes the block into a byte array
// @return bb - the byte array representing the block
// @return err - the error if any
func (b *Block) Serialize() (bb []byte, err error) {
	var result bytes.Buffer
	err = gob.NewEncoder(&result).Encode(b)
	if err != nil {
		return []byte{}, err
	}
	bb = result.Bytes()
	return
}

// Deserialize deserializes the block from a byte array
// @param data - the byte array representing the block
// @return bb - the block
// @return err - the error if any
func (b *Block) Deserialize(data []byte) (bb *Block, err error) {
	err = gob.NewDecoder(bytes.NewReader(data)).Decode(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
