package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/a-h/templ"
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

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {

	// get template from cache if it exists
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("The template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)
	// fmt.Println(buf)
	buf.WriteTo(w)

}

//
// // try and execute the block level template not the base
// // it would be possible to get the name of the template and then take that as the template name
// // but that would be an invisible constraint
// func (app *application) renderBlock(w http.ResponseWriter, status int, page string, templateName string, data *templateData) {
//
// 	// get template from cache if it exists
// 	// ts, ok := app.templateCache[page]
// 	// ts, ok := app.partialsCache[page]
// 	ts, ok := app.templateCache[page]
// 	if !ok {
// 		err := fmt.Errorf("The template %s does not exist", page)
// 		app.serverError(w, err)
// 		return
// 	}
//
// 	buf := new(bytes.Buffer)
//
// 	err := ts.ExecuteTemplate(buf, templateName, data)
// 	// err := ts.Execute(buf, data)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
//
// 	w.WriteHeader(status)
// 	// fmt.Println(buf)
// 	buf.WriteTo(w)
//
// }

func (app *application) renderTempl(w http.ResponseWriter, status int, page templ.Component) {

	// get template from cache if it exists
	// ts, ok := app.templateCache[page]
	// ts, ok := app.partialsCache[page]
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

// func (app *application) renderCompon

// Create a new data struct for any common data that we don't mind passing to all templates
func (app *application) newTemplateData(r *http.Request) *templateData {

	return &templateData{
		Processes:   &app.ProcessList,
		DisplayVars: &app.DisplayVars,
	}
}
