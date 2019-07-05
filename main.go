package main

import (
    "net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/ldarren/agogo/routes/wiki"
)

type HostSwitch map[string]http.Handler

func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(req.URL.Path, "/wiki") {
		if handler := hs['/wiki']; handler != nil {
			handler.ServeHTTP(w, r)
		} else {
			// Handle host names for which no handler is registered
			http.Error(w, "Forbidden", 403) // Or Redirect?
		}
	}
}

func main() {
	hs := make(HostSwitch)
	hs["/wiki"] = Wiki
    http.ListenAndServe(":8800", hs)
}
