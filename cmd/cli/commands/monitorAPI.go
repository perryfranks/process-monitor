package commands

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"

	monitorapi "procmon.perryfanks.nerd/internal/monitorAPI"
)

const baseUrl = "http://localhost:4000"

var monitorID int

func startPayload(name string, workspaceName string, user string) []byte {
	s := monitorapi.StartMonitor{
		Name:      name,
		Workspace: workspaceName,
		User:      user,
	}
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
	if err != nil {
		userName = ""
	}
	userName = useruser.Name

	return hostname, userName

}

// send the message to start the api service
func sendStart(name string) {

	workspace, user := getProcEnv()

	log.Print(baseUrl)

	payload := startPayload(name, workspace, user)
	resp, err := http.Post(baseUrl+"/api/start", "application/json", bytes.NewReader(payload))
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

func endPayload(id int, output string) []byte {
	s := monitorapi.EndMonitor{
		Id:          id,
		ReturnValue: output,
	}
	payload, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	return payload
}

func sendEnd(id int, output string) {

	payload := endPayload(id, output)
	resp, err := http.Post(baseUrl+"/api/end", "application/json", bytes.NewReader(payload))
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
