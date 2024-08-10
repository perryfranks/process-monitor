package commands

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"

	"github.com/spf13/cobra"
)

// Flags
var serverURL string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(waitCmd)
	rootCmd.PersistentFlags().StringVarP(&serverURL, "url", "u", "http://localhost:4000", "URL of the monitor server. Defaults to http://localhost:4000")
}

var rootCmd = &cobra.Command{
	Use:     "monitor",
	Short:   "Monitor a process via the matching remotes server. Will capture basic runtime information + output bby default",
	Long:    "Process passed as arguments will be ran and monitored. Please don't require input or pipe commands yet. If parsing args will need to be encased in quotes",
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

	var wg sync.WaitGroup
	wg.Add(1)
	var output []byte
	var err error
	var cmdErr, cmdOut bytes.Buffer

	var cmd *exec.Cmd
	cmd = exec.Command(args[0], args[1:]...)
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr

	// go func() {
	// 	defer wg.Done()
	// 	output, err = cmd.CombinedOutput()
	// }()
	//
	// pid := cmd.Process.Pid
	//
	// fmt.Println(pid)
	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start command - %v", err)
	}
	pid := strconv.Itoa(cmd.Process.Pid)
	sendStart(cmd.Path, pid)

	if err = cmd.Wait(); err != nil {
		sendEnd(monitorID, "Error running command.", cmd.ProcessState.ExitCode())
		log.Fatalf("Command finished unsuccessfully - %v", err)
	}

	outputString := cmdOut.String() + cmdErr.String()
	output = []byte(outputString)

	// when it finishes run the end
	fmt.Println("Output: ", string(output))
	sendEnd(monitorID, string(output), cmd.ProcessState.ExitCode())
	// expand from there

}
