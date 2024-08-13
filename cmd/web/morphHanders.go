package main

import (
	"errors"
	"fmt"
	"net/http"

	"procmon.perryfanks.nerd/internal/models"
)

// Render a list of the cards in the
func (app *application) morphRunningProcsList(w http.ResponseWriter, r *http.Request) {

	// Get the procs between the current head of the DisplayedProcesses and the real procs list
	if app.MostRecentRunningID == displayEmpty {
		// then we can just render the entire currently running list
		// does this actually get us anything since we're just going to push the new stuff to the client
		// app.DisplayedProcesses = app.ProcessList
		app.MostRecentRunningID = app.DisplayedProcesses[len(app.DisplayedProcesses)-1].Id

		// render & exit
		// render the whole list as new
		// return
	}

	var newProcs []models.Process
	// then we need the difference between the current head the rest of the procs in the list
	for i, v := range app.ProcessList {
		if v.Id == app.MostRecentRunningID {
			// then this is the current head
			// new items to display are app.ProcessList[i:]
			newProcs = app.ProcessList[i:]
			// TODO: rende"Bad logic. Inconsistent process lists"r
			// on finish the index should be updated so we don't need to care about that
			// need something to render
			return
		}
	}

	// otherwise we didn't match now what?
	// on finished proc we should update the index so we should only find or have displayEmpty
	app.serverError(w, errors.New("Bad logic. Inconsistent process lists"))
}
