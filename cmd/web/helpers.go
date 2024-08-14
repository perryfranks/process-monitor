package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/a-h/templ"
	"procmon.perryfanks.nerd/internal/models"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// 404 convience function that just fits with the others
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) renderTempl(w http.ResponseWriter, status int, page templ.Component) {

	// get template from cache if it exists
	buf := new(bytes.Buffer)

	// Render the component into the buffer
	err := page.Render(context.Background(), buf)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		return
	}

	// Write the HTTP status code
	w.WriteHeader(status)

	// Write the rendered content from the buffer to the response writer
	buf.WriteTo(w)
}

// Create a new data struct for any common data that we don't mind passing to all templates
func (app *application) newTemplateData(r *http.Request) *templateData {

	return &templateData{
		Processes:     &app.ProcessList,
		DisplayVars:   &app.DisplayVars,
		FinishedProcs: &app.FinishedList,
	}
}

// Add proc to the FinishedList, remove from ProcessList and update MostRecentRunningID if needed
func (app *application) finishProc(id int) (*models.Process, error) {

	for i, elem := range app.ProcessList {
		if elem.Id == id {

			// move to the finished list
			app.FinishedList = append(app.FinishedList, app.ProcessList[i])
			// update the finished time
			proc := &app.FinishedList[len(app.FinishedList)-1]
			// TODO: is this needed? if so refactor so its not
			if proc.Id != id {
				proc = &app.FinishedList[i-1] // HOPE
			}
			app.ProcessList = append(app.ProcessList[:i], app.ProcessList[i+1:]...)

			// handle the most recent
			if app.MostRecentRunningID == id {
				if i == 0 {
					// just deleted the first element so then update to null
					app.MostRecentRunningID = displayEmpty
				} else {
					// WARN: This seems a bit sus
					app.MostRecentRunningID = app.ProcessList[i-1].Id
				}
			}

			fmt.Println("MostRecentRunningID: ", app.MostRecentRunningID)
			return proc, nil
		}

	}

	return nil, errors.New("No process found with that ID")

}

func (app *application) checkFinished(id int) bool {
	for _, elem := range app.ProcessList {
		if elem.Id == id {
			// Process is still in the running list
			return false
		}
	}

	return true
}

func (app *application) getRunningProc(id int) *models.Process {
	return app.matchListOnId(app.ProcessList, id)
}

func (app *application) getFinishedId(id int) *models.Process {
	return app.matchListOnId(app.FinishedList, id)
}

func (app *application) matchListOnId(procs []models.Process, id int) *models.Process {
	for _, p := range app.ProcessList {
		if p.Id == id {
			return &p
		}
	}
	return nil

}
