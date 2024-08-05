package main

import (
	"html/template"

	"procmon.perryfanks.nerd/internal/models"
)

// lookup for names of functions and their functions
var functions = template.FuncMap{}

// holding structure for any data we want to pass to our
type templateData struct {
	DisplayVars   *DisplayVars
	Processes     *[]models.Process
	FinishedProcs *[]models.Process
}
