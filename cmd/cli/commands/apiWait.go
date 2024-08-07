package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Wait the specified time while emitting start & end monitor signals",
	Long:  "Wait the specified time while emitting start & end monitor signals. Will call the ",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

		sleepTime, err := strconv.Atoi(args[0])
		if err != nil {
			panic(1)
		}

		sendStart("SleepTest", "1")
		fmt.Println("ProcessID: ", monitorID)
		time.Sleep(time.Duration(sleepTime) * time.Second)
		fmt.Println("sleep over")
		sendEnd(monitorID, "dud output", 0)

	},
}
