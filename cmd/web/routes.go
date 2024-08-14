package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)

	router.HandlerFunc(http.MethodGet, "/components/procs", app.cardList)
	router.HandlerFunc(http.MethodGet, "/components/finishedprocs", app.finishedProcsCardList)
	router.HandlerFunc(http.MethodGet, "/components/poll/finished/:id", app.procAmFinished)
	router.HandlerFunc(http.MethodGet, "/morph/current", app.morphRunningProcsUpdate)
	router.HandlerFunc(http.MethodPost, "/components/set/finishedpollrate", app.finishedCardPollSet)
	router.HandlerFunc(http.MethodPost, "/components/clearfinished", app.clearFinishedProcs)

	// API (CLI) routes
	router.HandlerFunc(http.MethodPost, "/api/start", app.startMonitor)
	router.HandlerFunc(http.MethodPost, "/api/end", app.endMonitor)
	router.HandlerFunc(http.MethodGet, "/api/checkhealth", app.checkHealth)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)

}
