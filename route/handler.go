package route

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RouteHandler(w http.ResponseWriter, r *http.Request, routeRequest string) {

	if routeRequest == "all" {
		routes := getRoutes()
		jsonData, _ := json.Marshal(routes)
		fmt.Fprintf(w, "%s", jsonData)
	} else {
		route := getRouteById(routeRequest)
		jsonData, _ := json.Marshal(route)
		fmt.Fprintf(w, "%s", jsonData)
	}
	
}