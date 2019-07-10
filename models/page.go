package models

import (
	"os"
	"path"
	"runtime"
    "io/ioutil"
)

func init(){
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) Save() error {
    filename := "static/" + p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func (p *Page) Load() error {
    filename := "static/" + p.Title + ".txt"
    var err error
    p.Body, err = ioutil.ReadFile(filename)
    return err
}
