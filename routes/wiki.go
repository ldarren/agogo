package wiki

import (
    "fmt"
    "net/http"
    "github.com/ldarren/agogo/models"
    "github.com/ldarren/agogo/templates"
	"github.com/julienschmidt/httprouter"
)

engine := new Engine{}
engine.Init("./view.html", "./edit.html")

func readHandler(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    fmt.Printf("Method: %s, URL: %s\n", req.Method, req.URL.Path)
	title := ps.ByName("title")
    p := &pages.Page{Title: title}
    err := p.Load()
    if err != nil {
		createHandler(res, req, ps)
        return
    }

    renderTemplate(res, "view", p)
}

func editHandler(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    fmt.Printf("Method: %s, URL: %s\n", req.Method, req.URL.Path)
	title := ps.ByName("title")
    p := &pages.Page{Title: title}
    err := p.Load()
    if err != nil {
        p = &pages.Page{Title: title}
    }

    renderTemplate(res, "edit", p)
}

func createHandler(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    fmt.Printf("Method: %s, URL: %s\n", req.Method, req.URL.Path)
	title := ps.ByName("title")
    body := req.FormValue("body")
    p := &pages.Page{Title: title, Body: []byte(body)}
    err := p.Save()
    if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(res, req, "/wiki/" + title, http.StatusFound)
}

Wiki := httprouter.New()

Wiki.GET("/wiki/:title", readHandler)
Wiki.POST("/wiki/:title", createHandler)
Wiki.GET("/wiki/:title/edit", editHandler)
