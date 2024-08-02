package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"procmon.perryfanks.nerd/internal/models"
	monitorapi "procmon.perryfanks.nerd/internal/monitorAPI"
	"procmon.perryfanks.nerd/internal/templates"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// data := app.newTemplateData(r)
	// app.render(w, http.StatusOK, "home.html", data)
	// page := templates.Hello("User")
	page := templates.BasePage()
	app.renderTempl(w, http.StatusOK, page)
}

func (app *application) startMonitor(w http.ResponseWriter, r *http.Request) {

	var process models.Process
	var startMsg monitorapi.StartMonitor
	err := json.NewDecoder(r.Body).Decode(&startMsg)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	process.Workspace = startMsg.Workspace
	process.Id = app.idCount
	app.idCount++

	fmt.Println("New process: ", process)
	app.ProcessList = append(app.ProcessList, process)

	// return id as ack
	returnMsg := monitorapi.StartReturn{
		Id:      process.Id,
		Success: true,
	}

	jsonReturn, err := json.Marshal(returnMsg)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write(jsonReturn)

	fmt.Println(app.ProcessList)

}

func (app *application) endMonitor(w http.ResponseWriter, r *http.Request) {

	var endMsg monitorapi.EndMonitor
	err := json.NewDecoder(r.Body).Decode(&endMsg)
	if err != nil {
		fmt.Println("couldn't decode")
		app.clientError(w, http.StatusBadRequest)
		return
	}

	fmt.Println("removing process: ", endMsg.Id)
	// app.ProcessList = append(app.ProcessList, process)

	// delete that process
	err = app.deleteProc(endMsg.Id)
	if err != nil {
		fmt.Println("couldn't delete")
		app.clientError(w, http.StatusBadRequest)
		return

	}

	// return id as ack
	ack := monitorapi.Success{
		Success: true,
	}

	jsonReturn, err := json.Marshal(ack)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write(jsonReturn)

}
