package cmdtemplate

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wianvos/xlr/datamodels/template"
)

// define constants

var deleteLong = `Delete a template
Example:
  template delete templateID
`

// flag variables

func addDelete() {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete  a certain template",
		Long:  deleteLong,
		Run:   runDelete,
	}
	cmd.Flags().BoolVarP(&flagSearchByTitle, "byTitle", "t", false, "retrieve a template by name instead of by id")
	relCmd.AddCommand(cmd)
}

func runDelete(cmd *cobra.Command, args []string) {

	// check the nr of arguments
	if len(args) != 1 {
		fmt.Printf("invalid number of arguments: %d\n", len(args))
		os.Exit(3)
	}

	// declare variables
	var t template.Template
	var success bool
	var err error
	var applicationID string

	// instantiate the xlr client
	client := getClient()

	if flagSearchByTitle == true {
		applicationName := strings.Join(args, " ")

		t, err = client.Templates.GetByTitle(applicationName)

		if err != nil {
			panic(fmt.Errorf("goxlr: there was an error trying to retrieve id for : %s : %s", applicationName, err))
		}

		applicationID = t.ID

	} else {
		// args is our application ID here
		applicationID = strings.Join(args, " ")

		if strings.HasPrefix(applicationID, IDPrefix) == false {
			applicationID = IDPrefix + "/" + applicationID
		}

	}
	// query for a full list of the available templates
	success, err = client.Templates.Delete(applicationID)

	if err != nil {
		panic(fmt.Errorf("goxlr: there was an error trying to delete: %s : %s", applicationID, err))
	}

	if success {
		fmt.Println("Succesfully deleted template: ", applicationID)

	} else {
		fmt.Println("Unable to delete template: ", applicationID)
	}

}
