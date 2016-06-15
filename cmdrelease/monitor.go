package cmdrelease

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/wianvos/xlr"
	"github.com/wianvos/xlr/datamodels/template"
)

var monitorLong = `Monitors running releases in the system
Example:
  release monitor <release id's>
`

var flagInterval int

func addMonitor() {
	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitors releases in the system",
		Long:  monitorLong,
		Run:   runMonitor,
	}
	//add local long listing flag to the Command
	cmd.Flags().StringVarP(&flagStatus, "status", "", "", "show only releases with <status>")
	cmd.Flags().IntVarP(&flagInterval, "interval", "i", 5, "monitoring interval")
	relCmd.AddCommand(cmd)
}

func runMonitor(cmd *cobra.Command, args []string) {

	validateStatusFlag(flagStatus)

	//get the much needed config for the xlr client
	config := getConfig()

	// instantiate the xlr client
	client := xlr.NewClient(config)

	for {
		var releases template.Templates
		var err error
		// query for a full list of the available releases
		releases, err = client.Releases.List()
		// deal with any thrown errors
		if err != nil {
			panic(fmt.Errorf("Unable to retrieve releases: %s \n", err))
		}

		// seperate releases from the templates
		releases.GetReleases()
		//totally avoidable
		releases.SortByStatus()
		// check if we need to come up with a long or a short answer
		// format the output according to the flags
		switch flagLong {
		case true:
			for _, r := range releases {
				fmt.Println(template.RenderJSON(r))
			}
		case false:
			for _, r := range releases {
				fmt.Println(r.RenderJSONShort())
			}
		}
		time.Sleep(time.Duration(flagInterval) * time.Second)
	}
}
