package cobra

import "github.com/spf13/cobra"

var (
	dropBlocksCmd = &cobra.Command{
		Use:   "drop",
		Short: "Drop all blocks from the blockchain",
		Long:  `Drop all blocks from the blockchain`,
		Run:   dropBlocks,
	}
)

func dropBlocks(cmd *cobra.Command, args []string) {}
