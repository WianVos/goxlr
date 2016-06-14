package cmdrelease

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wianvos/xlr"
)

var listLong = `Return a list of releases in the system
Example:
  release list
`

func addList() {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List releases in the system",
		Long:  listLong,
		Run:   runList,
	}
	//add local long listing flag to the Command
	cmd.Flags().BoolVarP(&flagLong, "long", "l", false, "display a long listing")
	cmd.Flags().StringVarP(&flagStatus, "status", "", "", "show only releases with <status>")
	relCmd.AddCommand(cmd)
}

func runList(cmd *cobra.Command, args []string) {

	validateStatusFlag(flagStatus)

	var newreleases []xlr.Release

	//get the much needed config for the xlr client
	config := getConfig()

	// instantiate the xlr client
	client := xlr.NewClient(config)

	// query for a full list of the available releases
	releases, err := client.Releases.List()
	// deal with any thrown errors
	if err != nil {
		panic(fmt.Errorf("Unable to retrieve releases: %s \n", err))
	}

	//for some stupid reason our briljant developers have decided that quering
	//the rest interface should not only render all releases but should also return all templates as well
	//so we need to deal with that as always ... thnx guys ..
	for _, r := range releases {
		if r.OriginTemplateId != "" {
			if flagStatus != "" {
				if r.Status == flagStatus {
					newreleases = append(newreleases, r)
				}
			} else {
				newreleases = append(newreleases, r)
			}
		}
	}
	releases = newreleases
	//totally avoidable

	// check if we need to come up with a long or a short answer
	// format the output according to the flags
	switch flagLong {
	case true:
		for _, r := range releases {
			fmt.Println(r.RenderJSON())
		}
	case false:
		for _, r := range releases {
			fmt.Println(r.RenderJSONShort())
		}
	}

}
