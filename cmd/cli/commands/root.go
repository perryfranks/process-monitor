package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// rootCmd.AddCommand(whatever)
	rootCmd.AddCommand(waitCmd)
}

var rootCmd = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Root cmd")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		var a string
		fmt.Scanf("Press any key %s", a)
		os.Exit(1)
	}
}
