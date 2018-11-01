package cmds

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const confFilename = "configuration.json"
const urlKey = "quotes.url"
const apiKey = "quotes.api_version"
const usernameKey = "quotes.authorization.username"
const passwordKey = "quotes.authorization.password"

func init() {
	rootCmd.AddCommand(configurationCmd)
	configurationCmd.AddCommand(setConfCmd)
}

var configurationCmd = &cobra.Command{
	Use:   "configuration",
	Short: "Configuration subcommand",
	Long: `Configuration subcommand.
  Sets and updates the CLI configuration.
      - If any subcommand it prints the configuration`,
	Run:     conf,
	Aliases: []string{"conf", "c"},
}

var setConfCmd = &cobra.Command{
	Use:   "set",
	Short: "Initialize configuration.",
	Long:  `Initialize configuration.`,
	Run:   setConf,
}

func conf(cmd *cobra.Command, args []string) {
	fmt.Println(viper.Get("quotes.api_version"))
}

// func initConf(cmd *cobra.Command, args []string) {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Quotes URL: ")
// 	_url, _ := reader.ReadString('\n')
// 	fmt.Print("Authentication username: ")
// 	_username, _ := reader.ReadString('\n')
// 	fmt.Print("Authentication password: ")
// 	_password, _ := terminal.ReadPassword(0)
// 	fmt.Print("Default api version username: ")
// 	_api, _ := reader.ReadString('\n')
//
// 	viper.Set(urlKey, strings.TrimSuffix(_url, "\n"))
// 	viper.Set(usernameKey, strings.TrimSuffix(_username, "\n"))
// 	// TODO: Encypte password before saving it
// 	viper.Set(passwordKey, strings.TrimSuffix(string(_password), "\n"))
// 	viper.Set(apiKey, strings.TrimSuffix(_api, "\n"))
//
// 	err := viper.WriteConfig()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

func setConf(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Please specify a key and value")
		os.Exit(1)
	}
	fmt.Println(args[0])
	fmt.Println(args[1])
}

func initConfig() {

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	confPath := filepath.Join(home, ".foo/")
	// Create directory conf if not exists
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		var err = os.MkdirAll(confPath, os.ModePerm)
		if isError(err) {
			return
		}
	}

	// create configuration file if it does not exist
	confFilePath := filepath.Join(confPath, confFilename)
	if _, err := os.Stat(confFilePath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		file, err := os.Create(confFilePath)
		if isError(err) {
			return
		}
		defer file.Close()
		file.Write([]byte("{}"))
	}

	viper.SetConfigType("json")
	viper.SetConfigFile(confFilePath)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
	value := viper.GetString(urlKey)
	if value == "" {
		fmt.Println("Please specify the URL of Quotes - foo configuration set url https://quotes.domain ")
	}

	value = viper.GetString(usernameKey)
	if value == "" {
		fmt.Println("Please specify the usernaame - foo configuration set authorization.username your_login ")
	}

	value = viper.GetString(urlKey)
	if value == "" {
		fmt.Println("Please specify the password - foo configuration set authorization.password your_password ")
	}

	value = viper.GetString(urlKey)
	if value == "" {
		fmt.Println("Please specify the api version - foo configuration set api_version (v2|v2) ")
	}
}

// Helper function
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
