package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"procmon.perryfanks.nerd/internal/models"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
	// partialsCache map[string]*template.Template

	ProcessList []models.Process
	DisplayVars DisplayVars
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// template cache
	templateCache, err := newTemplateData()
	if err != nil {
		errorLog.Fatal(err)
	}

	dv := DisplayVars{}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
		// partialsCache: partialsCache,
		ProcessList: []models.Process{},

		DisplayVars: dv,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Starting server sever on ", *addr)
	err = srv.ListenAndServe()
	// err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
