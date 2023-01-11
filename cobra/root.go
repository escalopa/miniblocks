package cobra

import (
	"log"

	"github.com/escalopa/myblocks/application"
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
		Use:     "blockchain [command]",
		Short:   "Blockchain operations",
		Long:    `Blockchain operations`,
		Example: "blockchain block add --data 'some data' - adds a block to the blockchain",
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
