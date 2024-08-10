package commands

import (
	"github.com/spf13/cobra"
	monitorapi "procmon.perryfanks.nerd/internal/monitorAPI"
)

var checkHealthCmd = &cobra.Command{
	Use:   "checkhealth",
	Short: "Check connection to the to the monitor (path set with --url/-u)",
	Run: func(cmd *cobra.Command, args []string) {
		checkHealth(args)
	},
}

func checkHealth(args []string) {
	// assume there is a /checkhealth endpoint
	// would simply return a {"connection":true} msg
	var response monitorapi.CheckHealthMsg

}
