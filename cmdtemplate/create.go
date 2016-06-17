package cmdtemplate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/wianvos/xlr"
	"github.com/wianvos/xlr/datamodels/template"
)

var createLong = `Return a list of templates in the system
Example:
  templates create <template file>
`

var flagInFile string
var flagInDir string

func addCreate() {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create template in the system",
		Long:  createLong,
		Run:   runCreate,
	}

	//adding Flags
	cmd.Flags().StringVarP(&flagInFile, "infile", "i", "", "File to use for input")
	cmd.Flags().StringVarP(&flagInDir, "directory", "d", "", "Directory to use for input, all json files will be uploaded as it where profiles")
	//cmd.Flags().BoolVarP(&flagOverWrite, "overwrite", "O", false, "overwrite existing templates")

	relCmd.AddCommand(cmd)

}

//runCreate runs the actual create command
func runCreate(cmd *cobra.Command, args []string) {

	// initializing variables
	var inputFiles []string

	// setup the client connection
	config := getConfig()
	client := xlr.NewClient(config)

	// are we running on infile or indir ..
	if flagInFile != "" {
		inputFiles = append(inputFiles, flagInFile)
	}

	if flagInDir != "" {
		files, _ := filepath.Glob(flagInDir + "/*.json")
		for _, f := range files {
			inputFiles = append(inputFiles, f)
		}
	}

	// marshall it into a xlr.Template structure
	for _, file := range inputFiles {
		fmt.Println("attempting to upload: " + file)
		t := readJSONFile(file)
		_, err := client.Templates.CreateTemplate(t)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func readJSONFile(f string) template.Template {

	var template template.Template

	file, e := ioutil.ReadFile(f)

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &template)

	return template

}
