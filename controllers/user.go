package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/opiumated/officeSpace/models"
)

type (
	UserController struct {
		session *mgo.Session
	}
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}
}

//Retrieve an individual user's resource
func (userController UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := userController.session.DB("office_space").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.ID = bson.NewObjectId()
	uc.session.DB("office_space").C("users").Insert(u)
	uj, err := json.Marshal(u)
	checkError(err)
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
}
