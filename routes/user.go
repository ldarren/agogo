package routes

import (
    "fmt"
    "net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

var user = httprouter.New()

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func create(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	r, err := ioutil.ReadAll(req.Body)
	if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()
	var obj User
	err = json.Unmarshal(r, &obj)
	if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%+v\n", obj)
}

func read(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	obj := User{ Username: ps.ByName("username"), Password: "123" }
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
