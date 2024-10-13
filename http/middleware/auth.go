package middleware

import (
	"log"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check authenticate and log
		log.Println("Uri:", r.RequestURI, "Body:", r.Body)

		next.ServeHTTP(w, r)
	})
}
