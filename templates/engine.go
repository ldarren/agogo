package templates

import (
	"net/http"
    "html/template"
	"github.com/ldarren/agogo/models"
)

type Engine struct {
	Templates *template.Template
}


func (e *Engine) Init(fnames... string) {
	e.Templates = template.Must(template.ParseFiles(fnames...))
}

func (e *Engine) Render(res http.ResponseWriter, fname string, p *models.Page) {
    err := e.Templates.ExecuteTemplate(res, fname + ".html", p)
    if nil != err {
        http.Error(res, err.Error(), http.StatusInternalServerError)
    }
}
