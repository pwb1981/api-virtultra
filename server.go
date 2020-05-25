package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
)

var validPath = regexp.MustCompile("^/api/(route|SOMEOTHERROUTE)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.HandleFunc("/api/route/", makeHandler(routeHandler))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}