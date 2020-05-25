package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func routeHandler(w http.ResponseWriter, r *http.Request, route string) {
	var (
		p []byte
		err error)

	if route == "all" {
		p, err = getAllRoutes()
	} else {
		p, err = getRoute(route)
	}
	
	if err != nil {
		fmt.Fprintf(w, "%s", "ERROR")
		return
	}
	
	fmt.Fprintf(w, "%s", p)
}

func getRoute(routeId string) ([]byte, error) {
	filename := "routes/" + routeId + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getAllRoutes() ([]byte, error) {
	filename := "routes/all.txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return body, nil
}

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