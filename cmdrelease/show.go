package cmdrelease

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wianvos/xlr/datamodels/template"
)

var showLong = `Show details on a certain release
Example:
  release show
`

func addShow() {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show details of a certain release",
		Long:  showLong,
		Run:   runShow,
	}

	cmd.Flags().BoolVarP(&flagTitle, "title", "t", false, "Use release title instead of ID")
	cmd.Flags().BoolVarP(&flagShort, "short", "", false, "show a short overview of the release")

	relCmd.AddCommand(cmd)
}

func runShow(cmd *cobra.Command, args []string) {

	var releaseID string
	var err error

	client := getClient()

	releaseID = args[0]

	if flagTitle == true {
		releaseID, err = client.Releases.GetID(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}

	release, err := client.Releases.Get(releaseID)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if flagShort == true {
		fmt.Println(release.RenderJSONShort())
	} else {
		fmt.Println(template.RenderJSON(release))
	}
}
