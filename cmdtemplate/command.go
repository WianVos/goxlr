package cmdtemplate

import (
	"encoding/json"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wianvos/xlr"
)

// flag variables
var flagLong bool
var flagJSON bool
var flagOutFile string
var flagSearchByTitle bool

//constants
const (
	IDPrefix = "Applications"
)

// maskCmd represents the parent for all mask cli commands.
var relCmd = &cobra.Command{
	Use:   "template",
	Short: "template provides and interface to work with active templates",
}

//GetCommands grab and return commands in this package
func GetCommands() *cobra.Command {

	//collect the commands in the package
	addList()
	addShow()
	addCreate()
	addStart()
	addDelete()
	return relCmd
}

func renderJSON(l interface{}) string {

	b, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		panic(err)
	}
	s := string(b)

	return s
}

// lets get the output home shall we
func writeToFile(s string, f string) {
	d1 := []byte(s + "\n")
	err := ioutil.WriteFile(f, d1, 0644)
	if err != nil {
		panic(err)
	}
}

func getConfig() *xlr.Config {
	config := &xlr.Config{
		User:     viper.GetString("user"),
		Password: viper.GetString("password"),
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		Context:  viper.GetString("context"),
		Scheme:   viper.GetString("scheme"),
	}

	return config
}

func getClient() *xlr.Client {
	//get the much needed config for the xlr client
	config := getConfig()

	// instantiate the xlr client
	client := xlr.NewClient(config)

	return client

}
