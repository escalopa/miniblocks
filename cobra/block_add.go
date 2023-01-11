package cobra

import "github.com/spf13/cobra"

var (
	addBlockCmd = &cobra.Command{
		Use:     "add",
		Args:    cobra.ExactArgs(0),
		Short:   "Adds a block to the blockchain",
		Long:    `Adds a block to the blockchain with the given data after meeting the proof of work requirements`,
		Example: `block add -d "Hello World"`,
		Aliases: []string{"create", "new"},
		Run:     addBlock,
	}
)

var blockData string

func init() {
	addBlockCmd.Flags().StringVarP(&blockData, "data", "d", "", "Block data to store in the new created block")
}

func addBlock(cmd *cobra.Command, args []string) {
	if len(blockData) == 0 {
		cmd.Help()
		return
	}
	cobra.CheckErr(bc.AddBlock(blockData))
}
