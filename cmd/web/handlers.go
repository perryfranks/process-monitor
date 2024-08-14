package main

import (
	"net/http"

	"procmon.perryfanks.nerd/internal/models"
	"procmon.perryfanks.nerd/internal/templates"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	page := templates.BasePage(app.ProcessList, app.FinishedList)
	app.renderTempl(w, http.StatusOK, page)
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

	content := templates.FinishedPolledProcessList(procList, poll, app.Paused)
	// content := templates.PollProcessList(procList, "components/finishedprocs", poll, "")
	app.renderTempl(w, http.StatusOK, content)

}

// values will be AUTO/STOP
func (app *application) finishedCardPollSet(w http.ResponseWriter, r *http.Request) {

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

// components/clearfinished | Clear the finished procs array completely
func (app *application) clearFinishedProcs(w http.ResponseWriter, r *http.Request) {
	app.FinishedList = []models.Process{}

	// TODO: if we have a display counter then this needs to be updated
	w.WriteHeader(http.StatusNoContent)
}
