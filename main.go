// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/wianvos/goxlr/cmdrelease"
	"github.com/wianvos/goxlr/cmdtemplate"

	//external libraries
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config environmental variables.
var Host string
var Context string
var User string
var Password string
var Port string
var Scheme string

var goxlr = &cobra.Command{
	Use:   "goxlr",
	Short: "goxlr provides a command line interface to work with XL-Release",
}

func init() {
	goxlr.AddCommand(cmdrelease.GetCommands())
	goxlr.AddCommand(cmdtemplate.GetCommands())
	goxlr.PersistentFlags().StringVarP(&Host, "host", "x", "blah", "XL-Release hostname")
	goxlr.PersistentFlags().StringVarP(&Context, "context", "c", "/xl-release", "XL-Release context")
	goxlr.PersistentFlags().StringVarP(&User, "user", "u", "", "XL-Release username")
	goxlr.PersistentFlags().StringVarP(&Password, "password", "p", "", "XL-Release password")
	goxlr.PersistentFlags().StringVarP(&Port, "port", "P", "5516", "portnumber to reach XL-Release on")
	goxlr.PersistentFlags().StringVarP(&Scheme, "scheme", "s", "http", "http scheme to user")
	viper.BindPFlag("port", goxlr.PersistentFlags().Lookup("port"))
	viper.BindPFlag("host", goxlr.PersistentFlags().Lookup("host"))
	viper.BindPFlag("context", goxlr.PersistentFlags().Lookup("context"))
	viper.BindPFlag("user", goxlr.PersistentFlags().Lookup("user"))
	viper.BindPFlag("password", goxlr.PersistentFlags().Lookup("password"))
	viper.BindPFlag("scheme", goxlr.PersistentFlags().Lookup("scheme"))

}
func main() {

	// initialze config
	initializeConfig()

	goxlr.Execute()
}

//initialize the viper config
func initializeConfig() {
	// get input from config files
	// configfile name is goxlr
	viper.SetConfigName("goxlr")

	// add the filepaths that will be used
	viper.AddConfigPath("/etc/goxlr/")
	viper.AddConfigPath("$HOME/.goxlr")
	viper.AddConfigPath(".")
	// Handle errors reading the config file
	viper.ReadInConfig()
	//if err != nil {
	//	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	//}

}
