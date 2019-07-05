package engine

import (
	"net/http"
    "html/template"
)

type Engine struct {
	Templates *template.Template
}


func (e *Engine) Init(fnames... string) {
	e.Templates = template.Must(template.ParseFiles(fnames...))
}

func (e *Engine) Render(res http.ResponseWriter, fname string, p *pages.Page) {
    err := e.Templates.ExecuteTemplate(res, "./" + fname + ".html", p)
    if nil != err {
        http.Error(res, err.Error(), http.StatusInternalServerError)
    }
}
