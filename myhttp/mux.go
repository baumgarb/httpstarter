package myhttp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/baumgarb/httpstarter"
	"github.com/go-chi/chi/v5"
)

type contextKey int

var UserContextKey = contextKey(1)

type Mux struct {
	router *chi.Mux
}

func NewMux() *Mux {
	return &Mux{
		router: chi.NewRouter(),
	}
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}

func (m *Mux) Get(pattern string, handler func(w http.ResponseWriter, r *http.Request, user httpstarter.User) (Result, error)) {
	m.router.Get(pattern, func(w http.ResponseWriter, r *http.Request) {
		m.invokeHandler(w, r, handler)
	})
}

func (m *Mux) Post(pattern string, handler func(w http.ResponseWriter, r *http.Request, user httpstarter.User) (Result, error)) {
	m.router.Post(pattern, func(w http.ResponseWriter, r *http.Request) {
		m.invokeHandler(w, r, handler)
	})
}

func (m *Mux) Put(pattern string, handler func(w http.ResponseWriter, r *http.Request, user httpstarter.User) (Result, error)) {
	m.router.Put(pattern, func(w http.ResponseWriter, r *http.Request) {
		m.invokeHandler(w, r, handler)
	})
}

func (m *Mux) Delete(pattern string, handler func(w http.ResponseWriter, r *http.Request, user httpstarter.User) (Result, error)) {
	m.router.Delete(pattern, func(w http.ResponseWriter, r *http.Request) {
		m.invokeHandler(w, r, handler)
	})
}

func (m *Mux) Handle(pattern string, handler http.Handler) {
	m.router.Handle(pattern, handler)
}

func (m *Mux) invokeHandler(w http.ResponseWriter, r *http.Request, handler func(w http.ResponseWriter, r *http.Request, user httpstarter.User) (Result, error)) {
	user, ok := r.Context().Value(UserContextKey).(httpstarter.User)
	if !ok {
		panic(fmt.Sprintf("mux: invalid type for 'user' in context: '%v'", user))
	}

	result, err := handler(w, r, user)

	if err == nil {
		w.WriteHeader(result.StatusCode)
		if result.Dto == nil {
			return
		}
		err := json.NewEncoder(w).Encode(result.Dto)
		if err != nil {
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			log.Printf("mux: error while encoding DTO in successful result: %v", err)
		}
		return
	}

	switch err := err.(type) {
	case httpstarter.InputValidationError:
		log.Printf("mux: validation error in %v request on '%v': %v", r.Method, r.URL.Path, err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	case ClientErr:
		log.Printf("mux: client error in %v request on '%v': %v", r.Method, r.URL.Path, err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	default:
		log.Printf("mux: unexpected server error in %v request on '%v': %v", r.Method, r.URL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}
}

func (m *Mux) Use(next func(http.Handler) http.Handler) {
	m.router.Use(next)
}

func (m *Mux) URLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func (m *Mux) URLParamInt(r *http.Request, key string) (int, error) {
	p := chi.URLParam(r, key)
	return strconv.Atoi(p)
}
