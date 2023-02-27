package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/baumgarb/httpstarter"
	"github.com/baumgarb/httpstarter/inmem"
	"github.com/baumgarb/httpstarter/myhttp"
)

type todoAPI struct {
	ts  *httpstarter.TodoService
	mux *myhttp.Mux
}

func RegisterTodoAPI(mux *myhttp.Mux) {
	seed := []httpstarter.Todo{
		{ID: 1, Name: "Go Shopping", Done: false},
		{ID: 2, Name: "Go Beer Drinking", Done: true},
		{ID: 3, Name: "Play Tennis", Done: false},
	}
	store := inmem.NewInMemoryStore(seed, func(t httpstarter.Todo, id int) httpstarter.Todo {
		t.ID = id
		return t
	})
	svc := httpstarter.NewTodoService(store)
	api := &todoAPI{
		ts:  svc,
		mux: mux,
	}

	mux.Get("/api/v1/todos", api.GetAll)
	mux.Post("/api/v1/todos", api.Add)
	mux.Put("/api/v1/todos/{id}", api.Update)
}

func (ta *todoAPI) GetAll(w http.ResponseWriter, r *http.Request, u httpstarter.User) (myhttp.Result, error) {
	all, err := ta.ts.GetAll(u)
	if err != nil {
		return myhttp.EmptyResult, err
	}
	return myhttp.NewOkResult(http.StatusOK, all), nil
}

func (ta *todoAPI) Add(w http.ResponseWriter, r *http.Request, u httpstarter.User) (myhttp.Result, error) {
	var t httpstarter.Todo
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return myhttp.EmptyResult, fmt.Errorf("todoapi: error decoding request body: %w", err)
	}

	t, err = ta.ts.Add(t, u)
	if err != nil {
		return myhttp.EmptyResult, err
	}

	return myhttp.NewOkResult(http.StatusCreated, t), nil
}

func (ta *todoAPI) Update(w http.ResponseWriter, r *http.Request, u httpstarter.User) (myhttp.Result, error) {
	id, err := ta.mux.URLParamInt(r, "id")
	if err != nil {
		return myhttp.EmptyResult, myhttp.NewClientErrorf("todoapi: parameter '%v' is not a valid ID", ta.mux.URLParam(r, "id"))
	}

	var t httpstarter.Todo
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return myhttp.EmptyResult, fmt.Errorf("todoapi: error decoding request body: %w", err)
	}

	err = ta.ts.Update(id, t, u)
	if err != nil {
		return myhttp.EmptyResult, err
	}

	return myhttp.NewOkResult(http.StatusNoContent, nil), nil
}
