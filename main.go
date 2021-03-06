package main

import (
	"log"
	"strings"
    "net/http"
	"github.com/ldarren/agogo/routes"
)

type HostSwitch map[string]http.Handler

func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/wiki") {
		if handler := hs["/wiki"]; handler != nil {
			handler.ServeHTTP(w, r)
		} else {
			// Handle host names for which no handler is registered
			http.Error(w, "Forbidden", 403) // Or Redirect?
		}
	}
	if strings.Contains(r.URL.Path, "/users") {
		if handler := hs["/users"]; handler != nil {
			handler.ServeHTTP(w, r)
		} else {
			// Handle host names for which no handler is registered
			http.Error(w, "Forbidden", 403) // Or Redirect?
		}
	}
}

func main() {
	hs := make(HostSwitch)
	hs["/wiki"] = routes.GetWiki()
	hs["/users"] = routes.GetUsers()

	log.Printf("agogo is serving at 8800\n")
    http.ListenAndServe(":8800", hs)
}
