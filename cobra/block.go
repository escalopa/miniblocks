package cobra

import "github.com/spf13/cobra"

var (
	blockCmd = &cobra.Command{
		Use:   "block",
		Short: "Manage the state of the blockchain",
		Long:  `Manage the state of the blockchain by adding or removing blocks`,
	}
)

func init() {
	blockCmd.AddCommand(addBlockCmd)
	blockCmd.AddCommand(dropBlocksCmd)
}
