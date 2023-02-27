package mw

import (
	"log"
	"net/http"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("recover from: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Recover - Internal Server Error Stuff"))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
