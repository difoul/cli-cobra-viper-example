package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(catalogCmd)
	catalogCmd.AddCommand(orderCmd)
}

var catalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "The catalog subcommand",
	Long:  `The catalog subcommand`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Orders a service",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
