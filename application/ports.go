package application

import "github.com/escalopa/myblocks/pkg"

type IBlockchain interface {
	AddBlock(data string) error
	NewIterator() IBlockchainIterator
}

type IBlockchainIterator interface {
	Next() *pkg.Block
}

type IDatabase interface {
	SaveBlock(data *pkg.Block) error
	GetBlock(hash []byte) (*pkg.Block, error)
	GetLastHash() ([]byte, error)
	Close() error
}
