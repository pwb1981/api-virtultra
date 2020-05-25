package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

