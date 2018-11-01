package cmds

import (
	"fmt"

	"github.com/difoul/cli-cobra-viper-example/helpers"
	"github.com/spf13/cobra"
)

var quoteId string

func init() {
	rootCmd.AddCommand(quotesCmd)

	getQuoteCmd.PersistentFlags().StringVar(&quoteId, "id", "", "The quote id")
	getQuoteCmd.MarkPersistentFlagRequired("id")
	modifyQuoteCmd.PersistentFlags().StringVar(&quoteId, "id", "", "The quote id")
	modifyQuoteCmd.MarkPersistentFlagRequired("id")
	deleteQuoteCmd.PersistentFlags().StringVar(&quoteId, "id", "", "The quote id")
	deleteQuoteCmd.MarkPersistentFlagRequired("id")

	quotesCmd.AddCommand(allQuotesCmd)
	quotesCmd.AddCommand(addQuoteCmd)
	quotesCmd.AddCommand(getQuoteCmd)
	quotesCmd.AddCommand(modifyQuoteCmd)
	quotesCmd.AddCommand(deleteQuoteCmd)
}

var quotesCmd = &cobra.Command{
	Use:   "quotes",
	Short: "Quotes subcommand",
	Long:  `Quotes subcommand. By the default it returns the list of quotes`,
	//Run:     allQutes,
	Aliases: []string{"qt"},
}

var allQuotesCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all quotes",
	Long:    `List all quotes`,
	Run:     allQutes,
	Aliases: []string{"l", "all"},
}

var addQuoteCmd = &cobra.Command{
	Use:     "add",
	Short:   "Adds quote",
	Long:    `Adds quote`,
	Run:     addQuote,
	Aliases: []string{"a", "create", "c"},
}

var getQuoteCmd = &cobra.Command{
	Use:     "get",
	Short:   "Gest quote",
	Long:    `Gets quote`,
	Run:     quote,
	Aliases: []string{"g", "read", "r"},
}

var modifyQuoteCmd = &cobra.Command{
	Use:     "modify",
	Short:   "Modifies quote",
	Long:    `Modifies quote`,
	Run:     updateQuote,
	Aliases: []string{"m", "update", "u"},
}

var deleteQuoteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Deletes quote",
	Long:    `Deletes quote`,
	Run:     deleteQuote,
	Aliases: []string{"d"},
}

func allQutes(cmd *cobra.Command, args []string) {
	_url := fmt.Sprintf("%s/%s/quotes", url, apiVersion)
	res, _ := helpers.HttpGet(_url, "admin", "admin")
	fmt.Println(string(res))
}

func quote(cmd *cobra.Command, args []string) {
	_url := fmt.Sprintf("%s/%s/quotes/%s", url, apiVersion, quoteId)
	res, _ := helpers.HttpGet(_url, "admin", "admin")
	fmt.Println(string(res))
}

func deleteQuote(cmd *cobra.Command, args []string) {
	_url := fmt.Sprintf("%s/%s/quotes/%s", url, apiVersion, quoteId)
	res, _ := helpers.HttpDelete(_url, "admin", "admin")
	fmt.Println(string(res))
}

func addQuote(cmd *cobra.Command, args []string) {
	_url := fmt.Sprintf("%s/%s/quotes", url, apiVersion)
	res, _ := helpers.HttpPost(_url, "admin", "admin", "")
	fmt.Println(string(res))
}

func updateQuote(cmd *cobra.Command, args []string) {
	_url := fmt.Sprintf("%s/%s/quotes/%s", url, apiVersion, quoteId)
	res, _ := helpers.HttpPut(_url, "admin", "admin", "")
	fmt.Println(string(res))
}
