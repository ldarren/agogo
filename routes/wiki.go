package routes

import (
    "fmt"
    "net/http"
    "github.com/ldarren/agogo/models"
    "github.com/ldarren/agogo/templates"
	"github.com/julienschmidt/httprouter"
)

var engine templates.Engine
var mux = httprouter.New()

func readHandler(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    fmt.Printf("Method: %s, URL: %s\n", req.Method, req.URL.Path)
	title := ps.ByName("title")
    p := &models.Page{Title: title}
    err := p.Load()
    if err != nil {
		createHandler(res, req, ps)
        return
    }

    engine.Render(res, "view", p)
}

func editHandler(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    fmt.Printf("Method: %s, URL: %s\n", req.Method, req.URL.Path)
	title := ps.ByName("title")
    p := &models.Page{Title: title}
    err := p.Load()
    if err != nil {
        p = &models.Page{Title: title}
    }

    engine.Render(res, "edit", p)
}

func createHandler(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    fmt.Printf("Method: %s, URL: %s\n", req.Method, req.URL.Path)
	title := ps.ByName("title")
    body := req.FormValue("body")
    p := &models.Page{Title: title, Body: []byte(body)}
    err := p.Save()
    if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(res, req, "/wiki/" + title, http.StatusFound)
}

func init(){
	engine.Init("./templates/view.html", "./templates/edit.html")

	mux.GET("/wiki/:title", readHandler)
	mux.POST("/wiki/:title", createHandler)
	mux.GET("/wiki/:title/edit", editHandler)
}

func CreateWiki() http.Handler {
	return mux
}
