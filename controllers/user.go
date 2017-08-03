package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/opiumated/officeSpace/models"
)

type (
	UserController struct{}
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

func NewUserController() *UserController {
	return &UserController{}
}

//Retrieve an individual user's resource
func (userController UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := models.User{
		ID:          p.ByName("user_id"),
		FirtName:    "Ofonime",
		MiddleName:  "Francis",
		LastName:    "Usoro",
		Username:    "Baba",
		Email:       "baba.usoro@gloo.ng",
		Password:    "Phoenix#01",
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	userJSON, err := json.Marshal(user)
	checkError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", userJSON)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.ID = "foo" //this will be set by the user
	uj, err := json.Marshal(u)
	checkError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: only write status for now
	w.WriteHeader(200)
}
