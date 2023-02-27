package main

import (
	"log"
	"net/http"

	"github.com/baumgarb/httpstarter/api"
	"github.com/baumgarb/httpstarter/myhttp"
	"github.com/baumgarb/httpstarter/myhttp/mw"
)

func main() {
	mux := myhttp.NewMux()

	mux.Use(mw.Recover)
	mux.Use(mw.SimulateAuth("token-user-1"))
	mux.Use(mw.Authenticate)

	api.RegisterTodoAPI(mux)
	api.RegisterPingAPI(mux)

	log.Println("Listening on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
