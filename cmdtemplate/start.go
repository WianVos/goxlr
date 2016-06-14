package cmdtemplate

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wianvos/xlr"
)

var startLong = `Start a release of of a template in the system
Example:
  templates start template_id
`

var flagReleaseTitle string
var flagReleaseVariables string
var flagReleasePasswordVariables string
var flagTemplateID string
var flagTemplateName string
var flagSchedule bool

func addStart() {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a release from a template in the system",
		Long:  startLong,
		Run:   runStart,
	}

	//adding Flags
	cmd.Flags().StringVarP(&flagReleaseTitle, "title", "t", "", "Title of the release to be started")
	cmd.Flags().StringVarP(&flagTemplateID, "id", "i", "", "template id to use to start the release")
	cmd.Flags().StringVarP(&flagReleaseVariables, "vars", "", "", "variables to pass to the release: key:value,key:value")
	cmd.Flags().StringVarP(&flagReleasePasswordVariables, "pvars", "", "", "password variables to pass to the release: key:value,key:value")
	cmd.Flags().StringVarP(&flagTemplateName, "name", "n", "", "template name to use to start the release")
	cmd.Flags().BoolVarP(&flagSchedule, "schedule", "S", false, "create the release but do not start it")

	//cmd.Flags().BoolVarP(&flagOverWrite, "overwrite", "O", false, "overwrite existing templates")

	relCmd.AddCommand(cmd)

}

//runCreate runs the actual create command
func runStart(cmd *cobra.Command, args []string) {

	var release xlr.Release
	var err error
	// setup the client connection
	config := getConfig()
	client := xlr.NewClient(config)

	// Fail Fast
	// check all required flags for proper input
	if flagReleaseTitle == "" {

		fmt.Println("unable to create release without a release Title. Use -t to set one or see help for more detais")
		os.Exit(2)

	}

	// get the template id

	templateID := determineTemplateID(flagTemplateName, flagTemplateID, client)

	// cut up the variable inputs
	fv := formatVariablesString(flagReleaseVariables)
	fpv := formatVariablesString(flagReleasePasswordVariables)

	// start the releases
	if flagSchedule != true {
		fmt.Println("run")
		release, err = client.Templates.Start(templateID, flagReleaseTitle, fv, fpv)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("schedule")

		release, err = client.Templates.Create(templateID, flagReleaseTitle, fv, fpv)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(release.RenderJSONShort())

}

func formatVariablesString(v string) map[string]string {
	//assign m to be an empty map
	m := make(map[string]string)

	if v == "" {
		return m
	}

	//assign m to be an empty map
	//split up the input if any
	ss := strings.Split(v, ",")

	//loop over the split string
	for _, pair := range ss {
		//split into key value pair
		z := strings.Split(pair, ":")
		//drop it in the map
		m[z[0]] = z[1]
	}

	//return the map
	return m

}

func determineTemplateID(name string, id string, client *xlr.Client) string {

	templateID := id

	// if the flagTemplateName is propagated with a template name .. lets retrieve the matching id
	if name != "" {

		template, err := client.Templates.GetByTitle(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		templateID = template.ID
	}

	return templateID
}
