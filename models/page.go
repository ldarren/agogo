package models

import (
    "io/ioutil"
	"github.com/ldarren/agogo/config"
)

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) Save() error {
    filename := *config.PATH.Static + p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func (p *Page) Load() error {
    filename := *config.PATH.Static + p.Title + ".txt"
    var err error
    p.Body, err = ioutil.ReadFile(filename)
    return err
}
