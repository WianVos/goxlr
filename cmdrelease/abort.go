package cmdrelease

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wianvos/xlr/datamodels/template"
)

var abortLong = `abort a certain release
Example:
  release abort <id>
`

func addAbort() {
	cmd := &cobra.Command{
		Use:   "abort",
		Short: "abort a certain release",
		Long:  abortLong,
		Run:   runAbort,
	}

	cmd.Flags().BoolVarP(&flagTitle, "title", "t", false, "Use release title instead of ID")

	relCmd.AddCommand(cmd)
}

func runAbort(cmd *cobra.Command, args []string) {

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

	release, err := client.Releases.Abort(releaseID)
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
