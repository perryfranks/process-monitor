package main

import (
	"log"

	"procmon.perryfanks.nerd/internal/models"
)

const displayEmpty = 0

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger

	ProcessList  []models.Process
	FinishedList []models.Process

	// The processes the web is displaying
	// DisplayedProcesses  []models.Process
	MostRecentRunningID int // Pretty sure this should just be an index
	FinishedProcesses   []models.Process

	idCount     int
	Paused      bool
	DisplayVars DisplayVars
	StatusVars  StatusVars
}
