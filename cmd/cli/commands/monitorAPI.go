package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"

	"github.com/davecgh/go-spew/spew"
	monitorapi "procmon.perryfanks.nerd/internal/monitorAPI"
)

// const baseUrl = "http://localhost:4000"

var monitorID int

func startPayload(name string, workspaceName string, user string, pid string) []byte {
	s := monitorapi.StartMonitor{
		Name:      name,
		Workspace: workspaceName,
		User:      user,
		Pid:       pid,
	}

	fmt.Println("Sending start message: ", s)

	payload, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	return payload
}

// get the user and workspace
func getProcEnv() (hostname string, userName string) {

	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}

	useruser, err := user.Current()
	// fmt.Println(useruser)
	spew.Dump(useruser)
	if err != nil {
		userName = "Unknown"
	} else {
		// Not always set depending on the system
		if useruser.Name != "" {
			userName = useruser.Name
		} else if useruser.Username != "" {
			userName = useruser.Username
		} else {
			userName = "Unknown"
		}
	}

	return hostname, userName

}

// send the message to start the api service
func sendStart(name string, pid string) {

	workspace, user := getProcEnv()

	// log.Print(baseUrl)

	payload := startPayload(name, workspace, user, pid)
	resp, err := http.Post(serverURL+"/api/start", "application/json", bytes.NewReader(payload))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Start Monitor request failed with status: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body. %v", err)
	}

	log.Println("body: ", body)

	var startReturn monitorapi.StartReturn
	err = json.Unmarshal(body, &startReturn)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body. %v", err)
	}

	if !startReturn.Success {
		log.Fatalf("Transaction unsuccessful")
	}

	monitorID = startReturn.Id
}

func endPayload(id int, output string, exitCode int) []byte {
	s := monitorapi.EndMonitor{
		Id:         id,
		Output:     output,
		ExitStatus: exitCode,
	}

	spew.Dump(s)
	payload, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	return payload
}

func sendEnd(id int, output string, exitCode int) {

	payload := endPayload(id, output, exitCode)
	resp, err := http.Post(serverURL+"/api/end", "application/json", bytes.NewReader(payload))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Start Monitor request failed with status: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body. %v", err)
	}

	var endReturn monitorapi.Success
	err = json.Unmarshal(body, &endReturn)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body. %v", err)
	}
	log.Println(endReturn)

	if !endReturn.Success {
		log.Fatalf("Transaction unsuccessful")
	}

}
