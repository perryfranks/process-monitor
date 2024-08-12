package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"procmon.perryfanks.nerd/internal/models"
	monitorapi "procmon.perryfanks.nerd/internal/monitorAPI"
	"procmon.perryfanks.nerd/internal/templates"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	page := templates.BasePage(app.ProcessList, app.FinishedList)
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

	process.Name = startMsg.Name
	process.Workspace = startMsg.Workspace
	process.User = startMsg.User
	process.Id = app.idCount
	process.IdString = strconv.Itoa(process.Id)
	process.Pid = startMsg.Pid
	app.idCount++

	// assign a start time
	process.StartTime = time.Now()

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

	// delete that process
	finishedProc, err := app.finishProc(endMsg.Id)
	if err != nil {
		fmt.Println("couldn't delete")
		app.clientError(w, http.StatusBadRequest)
		return

	}
	// update anything
	finishedProc.FinishTime = time.Now()
	finishedProc.Finished = true
	finishedProc.CapturedOut = endMsg.Output
	finishedProc.ExitStatus = endMsg.ExitStatus

	spew.Dump(finishedProc)

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

// Render a list of the cards in the
func (app *application) cardList(w http.ResponseWriter, r *http.Request) {

	// get all the procs from app
	// pass to the temple function
	procList := app.ProcessList
	cardList := templates.ProcessList(procList)
	app.renderTempl(w, http.StatusOK, cardList)

}

func (app *application) finishedProcsCardList(w http.ResponseWriter, r *http.Request) {

	// get all the procs from app
	// pass to the temple function

	procList := app.FinishedList
	// cardList := templates.ProcessList(procList)
	var poll string
	if app.StatusVars.FinishedProcsListAuto == true {
		poll = "every 2s"
	} else {
		poll = ""
	}

	app.infoLog.Println("Trigger text: ", poll)

	content := templates.FinishedPolledProcessList(procList, poll)
	// content := templates.PollProcessList(procList, "components/finishedprocs", poll, "")
	app.renderTempl(w, http.StatusOK, content)

}

// values will be AUTO/STOP
func (app *application) finishedCardPollSet(w http.ResponseWriter, r *http.Request) {

	//

	app.infoLog.Println("Setting the poll rate to: ", app.StatusVars.FinishedProcsListAuto)
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// get the damage
	auto := r.PostForm.Get("auto")

	switch auto {
	case "auto":
		app.StatusVars.FinishedProcsListAuto = true
	case "flip":
		app.StatusVars.FinishedProcsListAuto = !app.StatusVars.FinishedProcsListAuto
	default:
		// error no change
		app.infoLog.Println("Incorrect value passed to finishedCardPollSet, value = ", auto)
	}

	app.infoLog.Println("Setting the poll rate to: ", app.StatusVars.FinishedProcsListAuto)

	http.Redirect(w, r, "/components/finishedprocs", http.StatusSeeOther)

}

// /api/checkHealth
// just return a checkHealth message with sucess = true
func (app *application) checkHealth(w http.ResponseWriter, r *http.Request) {
	result := monitorapi.CheckHealthMsg{
		Connection: true,
	}

	jResult, err := json.Marshal(result)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write(jResult)
}

// components/clearfinished | Clear the finished procs array completely
func (app *application) clearFinishedProcs(w http.ResponseWriter, r *http.Request) {
	app.FinishedList = []models.Process{}

	w.WriteHeader(http.StatusNoContent)
}
