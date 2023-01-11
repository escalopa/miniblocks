package application

import "github.com/escalopa/myblocks/pkg"

type IBlockchain interface {
	AddBlock(data string) error
	NewIterator() (IBlockchainIterator, error)
}

type IBlockchainIterator interface {
	Next() (*pkg.Block, error)
}

type IDatabase interface {
	SaveBlock(data *pkg.Block) error
	GetBlock(hash []byte) (*pkg.Block, error)
	GetLastHash() ([]byte, error)
	Close() error
}

type ICli interface {
	Run()
}
