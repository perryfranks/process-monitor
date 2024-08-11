package commands

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	monitorapi "procmon.perryfanks.nerd/internal/monitorAPI"
)

var checkHealthCmd = &cobra.Command{
	Use:   "checkhealth",
	Short: "Check connection to the to the monitor (path set with --url/-u)",
	Run: func(cmd *cobra.Command, args []string) {
		checkHealth()
	},
}

func checkHealth() {
	// assume there is a /checkhealth endpoint
	// would simply return a {"connection":true} msg
	var response monitorapi.CheckHealthMsg

	resp, err := http.Get(serverURL + "/api/checkhealth")
	if err != nil {
		log.Fatalf("No connection to server.\nServer url: %s\n error: %v", serverURL, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body. %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response. %v", err)
	}

	if response.Connection != true {

		log.Printf("Issue with connection to server.\nURL: %s\nConnection response:%v\n", serverURL, response.Connection)
	} else {
		log.Printf("Message exchanged. Connection to server good.\n")
	}

}
