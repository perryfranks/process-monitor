package main

import (
	"net/http"

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

}
