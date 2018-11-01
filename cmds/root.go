package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "foo",
	Short: "foo is a cli example based on Cobra.",
	Long: `foo is a cli example based on Cobra.
                different cases are implemented.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// Do Stuff Here
	// },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	confPath   string
	apiVersion string
	url        string
)

func init() {
	cobra.OnInitialize(initConfig)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.foo)")

	// rootCmd.PersistentFlags().StringVarP(&apiVersion, "apiVersion", "", "v1", "The API version")
	// rootCmd.PersistentFlags().StringVarP(&url, "url", "", "http://127.0.0.1:8080", "The url of the server")
	//rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	//rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("version", rootCmd.PersistentFlags().Lookup("apiVersion"))
	// viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	//viper.SetDefault("author", "Difoul BEN")
	//viper.SetDefault("license", "apache 2.0")
}
