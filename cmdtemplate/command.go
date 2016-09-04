package cmdtemplate

import "github.com/spf13/cobra"

// flag variables
var flagLong bool
var flagJSON bool
var flagOutFile string
var flagSearchByTitle bool

//constants
const (
	IDPrefix = "Applications"
)

// maskCmd represents the parent for all mask cli commands.
var relCmd = &cobra.Command{
	Use:   "template",
	Short: "template provides and interface to work with active templates",
}

//GetCommands grab and return commands in this package
func GetCommands() *cobra.Command {

	//collect the commands in the package
	addList()
	//addShow()
	//addCreate()
	//addDelete()
	return relCmd
}
