package api

import (
	"net/http"

	"github.com/baumgarb/httpstarter"
	"github.com/baumgarb/httpstarter/myhttp"
)

func RegisterPingAPI(mux *myhttp.Mux) {
	mux.Get("/api/v1/ping", func(w http.ResponseWriter, r *http.Request, user httpstarter.User) (myhttp.Result, error) {
		result := struct {
			Msg string `json:"msg"`
		}{Msg: "pong"}
		return myhttp.NewOkResult(http.StatusOK, result), nil
	})
}
