package pages

import (
    "testing"
    "bytes"
)

func TestPage(t *testing.T) {
    title := "TestPage"
    body := []byte("This is a test page.")
    p1 := &Page{Title: title, Body: body}
    err1 := p1.Save()
    if err1 != nil {
        t.Error(err1.Error())
    }

    p2 := &Page{Title: title}
    err2 := p2.Load()
    if err2 != nil {
        t.Error(err2.Error())
    }

    if 0 != bytes.Compare(body, p2.Body) {
        t.Fail()
    }
}
