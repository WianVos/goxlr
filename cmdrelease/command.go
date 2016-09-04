package cmdrelease

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wianvos/xlr"
)

// flag variables
var flagStatus string
var flagLong bool
var flagJSON bool
var flagOutFile string
var flagTitle bool
var flagShort bool

// maskCmd represents the parent for all mask cli commands.
var relCmd = &cobra.Command{
	Use:   "release",
	Short: "release provides and interface to work with active releases",
}

//GetCommands grab and return commands in this package
func GetCommands() *cobra.Command {

	//collect the commands in the package

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

func validateStatusFlag(s string) bool {
	switch s {
	case
		"",
		"TEMPLATE",
		"PLANNED",
		"IN_PROGRESS",
		"PAUSED",
		"FAILING",
		"FAILED",
		"COMPLETED",
		"ABORTED":
		return true
	}

	fmt.Printf("invalid value provided for Status: %s\n", s)
	os.Exit(3)
	return false
}

func getClient() *xlr.Client {
	//get the much needed config for the xlr client
	config := &xlr.Config{
		User:     viper.GetString("user"),
		Password: viper.GetString("password"),
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		Context:  viper.GetString("context"),
		Scheme:   viper.GetString("scheme"),
	}

	// instantiate the xlr client
	client := xlr.NewClient(config)

	return client

}
