package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/spf13/viper"
	"github.com/wianvos/xlr"
)

//GetClient returns a xlr client object using the configuration from viper
func GetClient() *xlr.Client {
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

//RenderJSON returns a json representation of a struct
func RenderJSON(l interface{}) string {

	b, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		panic(err)
	}
	s := string(b)

	return s
}

//WriteToFile dumps a string to a file on the filesystem
func WriteToFile(s string, f string) {
	d1 := []byte(s + "\n")
	err := ioutil.WriteFile(f, d1, 0644)
	if err != nil {
		panic(err)
	}
}

//WriteJSONToFile writes a struct to a file as json
func WriteJSONToFile(l interface{}, file string) {
	WriteToFile(RenderJSON(l), file)
}
