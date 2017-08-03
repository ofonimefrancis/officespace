package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/opiumated/officeSpace/controllers"
)

func main() {
	router := httprouter.New()

	controller := controllers.NewUserController()
	router.GET("/api/user/:user_id", controller.GetUser)
	router.POST("/api/user", controller.CreateUser)
	router.DELETE("/api/user/:user_id", controller.RemoveUser)
	http.ListenAndServe(":8080", router)
}
