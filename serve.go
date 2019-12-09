package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", logger(cors(fs)))

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(requestDump))

		next.ServeHTTP(w, r)
	})
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
