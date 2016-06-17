package cmdtemplate

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wianvos/xlr"
	"github.com/wianvos/xlr/datamodels/template"
)

var showLong = `Show details on a certain template
Example:
  template show templateID
`

// flag variables

func addShow() {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show details of a certain template",
		Long:  showLong,
		Run:   runShow,
	}
	cmd.Flags().BoolVarP(&flagLong, "long", "l", false, "display a long listing")
	cmd.Flags().BoolVarP(&flagJSON, "json", "j", false, "display in json format")
	cmd.Flags().StringVarP(&flagOutFile, "outfile", "o", "", "File to use for output")
	cmd.Flags().BoolVarP(&flagSearchByTitle, "byTitle", "t", false, "retrieve a template by name instead of by id")
	relCmd.AddCommand(cmd)
}

func runShow(cmd *cobra.Command, args []string) {

	// check the nr of arguments
	if len(args) != 1 {
		fmt.Printf("invalid number of arguments: %d\n", len(args))
		os.Exit(3)
	}

	// declare variables
	var t template.Template
	var err error

	//get the much needed config for the xlr client
	config := getConfig()

	// instantiate the xlr client
	client := xlr.NewClient(config)

	if len(args) != 1 {
		panic("goxlr: invalid number of arguments: " + strings.Join(args, " "))
	}

	if flagSearchByTitle == true {
		applicationName := strings.Join(args, " ")

		t, err = client.Templates.GetByTitle(applicationName)

		if err != nil {
			fmt.Println(fmt.Errorf("goxlr: there was an error trying to retrieve: %s : %s", applicationName, err))
			os.Exit(1)
		}

	} else {
		// args is our application ID here
		applicationID := strings.Join(args, " ")

		if strings.HasPrefix(applicationID, IDPrefix) == false {
			applicationID = IDPrefix + "/" + applicationID
		}

		// query for a full list of the available templates
		t, err = client.Templates.Get(applicationID)

		if err != nil {
			panic(fmt.Errorf("goxlr: there was an error trying to retrieve: %s : %s", applicationID, err))
		}
	}

	if flagOutFile == "" {
		if flagLong == false {
			fmt.Println(t.RenderJSONShort())
		} else {
			fmt.Println(template.RenderJSON(t))
		}
	} else {
		writeToFile(template.RenderJSON(t), flagOutFile)
	}

}
