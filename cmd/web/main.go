package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"procmon.perryfanks.nerd/internal/models"
)

type StatusVars struct {
	FinishedProcsListAuto bool // auto refresh procs (and kill user focus)

}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger

	ProcessList  []models.Process
	FinishedList []models.Process
	idCount      int
	DisplayVars  DisplayVars
	StatusVars   StatusVars
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dv := DisplayVars{}

	app := &application{
		errorLog:     errorLog,
		infoLog:      infoLog,
		ProcessList:  []models.Process{},
		FinishedList: []models.Process{},
		idCount:      1,

		DisplayVars: dv,
		StatusVars: StatusVars{
			FinishedProcsListAuto: true,
		},
	}

	app.ProcessList = append(app.ProcessList, models.Process{
		Name:      "test",
		Workspace: "test",
		Id:        -1,
		IdString:  "-1",
	})

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Starting server sever on ", *addr)
	err := srv.ListenAndServe()
	// err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
