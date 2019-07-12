package routes

import (
    "log"
	"context"
    "net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/ldarren/agogo/models"
)

var user = httprouter.New()

func create(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	ctx := context.Background()
	r, err := ioutil.ReadAll(req.Body)
	if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()
	var obj models.User
	err = json.Unmarshal(r, &obj)
	if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("%+v\n", obj)
	err = obj.Create(ctx)
	if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	output, err := json.Marshal(&obj)
	if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("content-type", "application/json")
	res.Write(output)
	log.Printf("user.write %v\n", output)
}

func read(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Printf("GET /users/:username\n")
	ctx := context.Background()
	obj := models.User{ Username: ps.ByName("username"), Password: "" }
	err := obj.Read(ctx)
	if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(&obj)
	if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("content-type", "application/json")
	res.Write(output)
	log.Printf("user.read %v\n", output)
}

func init(){
	user.POST("/users", create)
	user.GET("/users/:username", read)
}

func GetUsers() http.Handler {
	return user
}
