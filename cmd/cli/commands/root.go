package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(waitCmd)
	rootCmd.AddCommand(monitorCmd)
}

var rootCmd = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Root cmd")
	},
}

func Execute() {
	if err := monitorCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
