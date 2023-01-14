package cobra

import (
	"log"

	"github.com/escalopa/miniblocks/application"
	"github.com/spf13/cobra"
)

type BlockchainCmd struct {
	c *cobra.Command
}

var (
	bc application.IBlockchain
)

var (
	rootCmd = &cobra.Command{
		Use:     "miniblock [command]",
		Short:   "Blockchain operations",
		Long:    `Blockchain operations`,
		Example: "miniblock print -c 1\nminiblock block create -d 'some data'",
	}
)

func New(ibc application.IBlockchain) (application.ICli, error) {
	bc = ibc
	rootCmd.AddCommand(blockCmd, printCmd)
	return &BlockchainCmd{c: rootCmd}, nil
}

func (bcc *BlockchainCmd) Run() {
	log.Println(bcc.c.Execute())
}
