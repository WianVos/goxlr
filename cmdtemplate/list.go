package cmdtemplate

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wianvos/goxlr/utils"
)

var listLong = `Return a list of templates in the system
Example:
  templates list
`

type listOutput interface {
	render() string
	renderJSON()
}

// define the output
type listOutputShort struct {
	ID    string
	Title string
}

type listOutputLong struct {
	listOutputShort
	CreatedAt      string
	LastModifiedAt string
	Description    string
}

func addList() {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List templates in the system",
		Long:  listLong,
		Run:   runList,
	}

	//add local long listing flag to the Command
	cmd.Flags().BoolVarP(&flagJSON, "json", "j", false, "display in json format")
	cmd.Flags().StringVarP(&flagOutFile, "outfile", "o", "", "File to use for output")
	relCmd.AddCommand(cmd)

}

func runList(cmd *cobra.Command, args []string) {
	//declare function variables
	// output will hold the output ... haha u guessed it

	//get the much needed client object for xlr
	client := utils.GetClient()

	// query for a full list of the available templates
	templates, err := client.Templates.List()
	// deal with any thrown errors
	if err != nil {
		panic(fmt.Errorf("Unable to retrieve templates: %s \n", err))
	}

	// check if we need to come up with a long or a short answer

	//render

	for _, s := range templates {
		fmt.Println(s.ID)
	}

}
