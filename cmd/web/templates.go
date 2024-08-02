package main

import (
	"fmt"
	"html/template"
	"path/filepath"

	"procmon.perryfanks.nerd/internal/models"
)

// lookup for names of functions and their functions
var functions = template.FuncMap{}

// holding structure for any data we want to pass to our
type templateData struct {
	DisplayVars *DisplayVars
	Processes   *[]models.Process
}

// tmpl -> html
func newTemplateData() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Load page templates
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extract the file name (e.g., "home.html") from the full path
		name := filepath.Base(page)

		// Register the template functions and parse base, partials, and page
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	// Load partials separately
	partials, err := filepath.Glob("./ui/html/partials/*.html")
	if err != nil {
		return nil, err
	}

	for _, partial := range partials {
		name := filepath.Base(partial)

		// Register the template functions and parse the partial template
		ts, err := template.New(name).Funcs(functions).ParseFiles(partial)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	fmt.Println(cache)

	return cache, nil

}
