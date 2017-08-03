package main

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
	"github.com/opiumated/officeSpace/controllers"
)

func main() {
	router := httprouter.New()

	controller := controllers.NewUserController(getSession())
	router.GET("/api/user/:user_id", controller.GetUser)
	router.POST("/api/user", controller.CreateUser)
	router.DELETE("/api/user/:user_id", controller.RemoveUser)
	http.ListenAndServe(":8080", router)
}

func getSession() *mgo.Session {
	//connect to local mongodb
	mongoSession, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer mongoSession.Close()
	return mongoSession
}
