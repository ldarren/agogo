package routes

import (
    "fmt"
    "net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/ldarren/agogo/models"
)

var user = httprouter.New()

func create(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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
	fmt.Printf("%+v\n", obj)
	obj.Create()

	output, err := json.Marshal(&obj)
	if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("content-type", "application/json")
	res.Write(output)
	fmt.Println(string(output))
}

func read(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	obj := models.User{ Username: ps.ByName("username"), Password: "" }
	obj.Read()

	output, err := json.Marshal(&obj)
	if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("content-type", "application/json")
	res.Write(output)
	fmt.Println(string(output))
}

func init(){
	user.POST("/users", create)
	user.GET("/users/:username", read)
}

func GetUsers() http.Handler {
	return user
}
