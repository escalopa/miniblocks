package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:     "print",
	Short:   "Prints `n` amount of blocks in the chain",
	Long:    `Prints all the blocks in the chain`,
	Run:     printBlocks,
	Example: `block print -n 10`,
}

var count int

func init() {
	printCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of blocks to print")
}

func printBlocks(cmd *cobra.Command, args []string) {
	iterator, err := bc.NewIterator()
	cobra.CheckErr(err)

	for {
		block, err := iterator.Next()
		cobra.CheckErr(err)

		line := fmt.Sprintf("---------------->>\nPrev. hash: %x\nHash: %x\nData: %s\nNonce: %d\n<<----------------",
			block.PrevHash, block.Hash, block.Data, block.Nonce)
		fmt.Println(line)
		if len(block.PrevHash) == 0 || count <= 1 {
			break
		}
		count--
	}
}
