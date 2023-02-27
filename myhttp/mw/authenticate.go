package mw

import (
	"context"
	"net/http"

	"github.com/baumgarb/httpstarter"
	"github.com/baumgarb/httpstarter/myhttp"
)

var dummySessions map[string]httpstarter.User = map[string]httpstarter.User{
	"token-user-1": {ID: 1, Email: "bb@dt.com"},
	"token-user-2": {ID: 2, Email: "sb@dt.com"},
}

func SimulateAuth(token string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Add("X-Auth-Token", token)
			next.ServeHTTP(w, r)
		})
	}
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth-Token")
		user := httpstarter.AnonymousUser
		if len(token) > 0 {
			u, ok := dummySessions[token]
			if ok {
				user = u
			}
		}

		ctx := context.WithValue(r.Context(), myhttp.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
