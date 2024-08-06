package commands

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// cli monitor ls -l /dir

var monitorCmd = &cobra.Command{
	Use:     "monitor",
	Short:   "monitor process via the remote process monitor",
	Long:    "process passed as arguments will be ran and monitored. Please don't require input or pipe commands yet. If parsing args will need to be encased in quotes",
	Example: "monitor \"ls -l tmp\"",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		monitorProcess(args)

	},
}

func monitorProcess(args []string) {
	if len(args) == 1 {
		args = parseQuoteCmd(args[0])
	}
	// run the command provided in the args
	cmd := exec.Command(args[0], args[1:]...)
	// when it runs send the start monitor
	sendStart(cmd.Path)
	// TODO: if this is goroutine then we capture the output we might be able to
	// get PID and output + enable kill feature
	output, err := cmd.CombinedOutput()
	if err != nil {
		sendEnd(monitorID, "Error running command.")
		log.Fatalf("Error running command. Command: %v. Error: %v", cmd, err)
	}

	// when it finishes run the end
	fmt.Println("Output: ", string(output))
	sendEnd(monitorID, string(output))
	// expand from there

}

// take a command as a string and then split it into words
func parseQuoteCmd(s string) []string {
	var invocation []string
	invocation = strings.Split(s, " ")

	return invocation

}
