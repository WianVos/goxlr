package cmdtemplate

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wianvos/xlr"
	"github.com/wianvos/xlr/datamodels/template"
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
	cmd.Flags().BoolVarP(&flagLong, "long", "l", false, "display a long listing")
	cmd.Flags().BoolVarP(&flagJSON, "json", "j", false, "display in json format")
	cmd.Flags().StringVarP(&flagOutFile, "outfile", "o", "", "File to use for output")

	relCmd.AddCommand(cmd)

}

func runList(cmd *cobra.Command, args []string) {
	//declare function variables
	// output will hold the output ... haha u guessed it
	var output []string

	//get the much needed config for the xlr client
	config := getConfig()

	// instantiate the xlr client
	client := xlr.NewClient(config)

	// query for a full list of the available templates
	templates, err := client.Templates.List()
	// deal with any thrown errors
	if err != nil {
		panic(fmt.Errorf("Unable to retrieve templates: %s \n", err))
	}

	// check if we need to come up with a long or a short answer

	// format the output according to the flags
	switch flagLong {
	case false:
		for _, t := range templates {
			output = append(output, t.RenderJSONShort())
		}
	case true:
		for _, t := range templates {
			output = append(output, template.RenderJSON(t))
		}
	}

	//render

	if flagOutFile == "" {
		for _, s := range output {
			fmt.Println(s)
		}
	} else {
		var outputString string
		for _, s := range output {
			outputString = outputString + "\n" + s
		}
		writeToFile(outputString, flagOutFile)
	}

}

// render functions
func (l listOutputShort) render() string {
	return fmt.Sprintf("ID:%s;Title:%s\n", l.ID, l.Title)
}

func (l listOutputLong) render() string {
	return fmt.Sprintf("ID:%s;Title:%s;Desc:%s\n", l.ID, l.Title, l.Description)
}

// func renderJSON(l interface{}) string {
//
// 	b, err := json.MarshalIndent(l, "", " ")
// 	if err != nil {
// 		panic(err)
// 	}
// 	s := string(b)
//
// 	return s
// }

// turn listoutput into strings we can work with
func renderString(i []listOutput) string {

	var stringOutput string

	for _, l := range i {
		stringOutput = fmt.Sprintf("%s %s", stringOutput, l)
	}
	return stringOutput
}

// render json ... satisfy interface
func (l listOutputShort) renderJSON() {
	renderJSON(l)
}

// render json ... satisfy interface
func (l listOutputLong) renderJSON() {
	renderJSON(l)
}
