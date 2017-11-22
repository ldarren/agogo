package main

import (
    "fmt"
    "strings"
    "net/http"
    "html/template"
    "test/darren/gowiki/pages"
)

var templates = template.Must(template.ParseFiles("./templates/view.html", "./templates/edit.html"))

func router(res http.ResponseWriter, req *http.Request) {
    fmt.Printf("Method: %s, URL: %s\n", req.Method, req.URL.Path)

    if strings.Contains(req.URL.Path, "/wiki/edit/") {
        title := req.URL.Path[len("/wiki/edit/"):]
        editHandler(res, req, title)
    } else {
        title := req.URL.Path[len("/wiki/"):]
        switch req.Method {
        case http.MethodPost:
            createHandler(res, req, title)
        default:
            readHandler(res, req, title)
        }
    }
}

func renderTemplate(res http.ResponseWriter, templ string, p *pages.Page) {
    err := templates.ExecuteTemplate(res, templ + ".html", p)
    if nil != err {
        http.Error(res, err.Error(), http.StatusInternalServerError)
    }
}

func readHandler(res http.ResponseWriter, req *http.Request, title string) {
    p := &pages.Page{Title: title}
    err := p.Load()
    if err != nil {
        http.Redirect(res, req, "/wiki/edit/" + title, http.StatusFound)
        return
    }

    renderTemplate(res, "view", p)
}

func editHandler(res http.ResponseWriter, req *http.Request, title string) {
    p := &pages.Page{Title: title}
    err := p.Load()
    if err != nil {
        p = &pages.Page{Title: title}
    }

    renderTemplate(res, "edit", p)
}

func createHandler(res http.ResponseWriter, req *http.Request, title string) {
    body := req.FormValue("body")
    p := &pages.Page{Title: title, Body: []byte(body)}
    err := p.Save()
    if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(res, req, "/wiki/" + title, http.StatusFound)
}

func main() {
    http.HandleFunc("/", router)
    http.ListenAndServe(":8080", nil)
}
