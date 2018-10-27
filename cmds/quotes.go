package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(quotesCmd)
}

var quotesCmd = &cobra.Command{
	Use:   "quotes",
	Short: "Lists all the quotes",
	Long:  `Lists all the quotes`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("==> ALL THE QUOTES SHOULD BE PRINTED <==")
		fmt.Println(apiVersion)
	},
	Aliases: []string{"qt"},
}
