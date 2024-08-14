package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"procmon.perryfanks.nerd/internal/models"
	"procmon.perryfanks.nerd/internal/templates"
)

// Render a list of the cards in the
func (app *application) morphRunningProcsUpdate(w http.ResponseWriter, r *http.Request) {

	// Get the procs between the current head of the DisplayedProcesses and the real procs list
	if app.MostRecentRunningID == displayEmpty {
		// then we can just render the entire currently running list
		// does this actually get us anything since we're just going to push the new stuff to the client
		// app.DisplayedProcesses = app.ProcessList
		app.MostRecentRunningID = app.ProcessList[len(app.ProcessList)-1].Id

		// render & exit
		// render the whole list as new
		// return
		fmt.Println("Rending whole list")
		page := templates.MorphRunningProcsList(app.ProcessList, "fade-in", "every 2s", "/components/poll/finished", "")
		app.renderTempl(w, http.StatusOK, page)
		return
	}

	var newProcs []models.Process
	// then we need the difference between the current head the rest of the procs in the list
	for i, v := range app.ProcessList {
		if v.Id == app.MostRecentRunningID {
			// check for this being the end of the list. -> display is up to date
			if i == len(app.ProcessList)-1 {
				fmt.Println("returning no content")
				w.WriteHeader(http.StatusNoContent)
				return

			}

			newProcs = app.ProcessList[i+1:]
			app.MostRecentRunningID = newProcs[len(newProcs)-1].Id

			fmt.Println("newProcs:  ", newProcs)

			page := templates.MorphRunningProcsList(newProcs, "fade-in", "every 2s", "/components/poll/finished", "")
			app.renderTempl(w, http.StatusOK, page)
			return
		}
	}

	// otherwise we didn't match now what?
	// on finished proc we should update the index so we should only find or have displayEmpty
	app.serverError(w, errors.New("Bad logic. Inconsistent process lists"))
}

// Handles the poll from an individual processchecking if it should be counted as finished and removed from display
// URL: /components/poll/finished/:id
func (app *application) procAmFinished(w http.ResponseWriter, r *http.Request) {
	ps := httprouter.ParamsFromContext(r.Context())
	if ps == nil {
		app.errorLog.Println("Params to procAmFinished was null")
		app.clientError(w, http.StatusInternalServerError)
	}

	idRaw := ps.ByName("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		app.errorLog.Println("Error converting id. Could be malformed. id: ", idRaw)
		app.serverError(w, err)
		return
	}

	if app.checkFinished(id) {
		// ASSUMPTION: this would be called in the api we simply have to remove it from the displayList
		// app.finishProc(id) && app.finishDisplay()

		// return nothing
		fmt.Printf("id: %v was found to be finished returning no content\n", id)

		proc := *app.getFinishedId(id)
		card := templates.MorphCard(proc, "fade-out", "every 1s", "/components/end", "delete")
		app.renderTempl(w, http.StatusOK, card)
	} else {
		// fmt.Printf("id: %v just quiet ending\n", id)
		proc := *app.getRunningProc(id)
		card := templates.MorphCard(proc, "", "every 2s", "/components/poll/finished", "")
		app.renderTempl(w, http.StatusOK, card)

	}

}

// basically a nop
// components/end
func (app *application) clearCard(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
