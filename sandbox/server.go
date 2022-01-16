package main

import (
	"log"
	"net/http"
)

func cors(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fs.ServeHTTP(w, r)
	}
}

func csp(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'unsafe-inline' 'self'; "+
			"script-src *.gstatic.com 'unsafe-inline' 'self'; "+
			"worker-src 'self'")
		fs.ServeHTTP(w, r)
	}
}

func main() {
	fs := http.FileServer(http.Dir("."))
	log.Println("Listening on :8081")
	http.ListenAndServe(":8081", csp(cors(fs)))
}
