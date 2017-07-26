package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID         int64  `json:"uid"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	UserName   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsAdmin    bool   `json:"is_admin"`
}

var database *sql.DB
var err error

func main() {
	database, err = sql.Open("mysql", "root:glootian@/office_space")
	routes := mux.NewRouter()
	routes.HandleFunc("/api/user/create", CreateUser).Methods("GET")
	routes.HandleFunc("/api/user/read/{uid:\\d+}", GetUser).Methods("GET")
	http.Handle("/", routes)
	http.ListenAndServe(":4000", nil)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Pragma", "no-cache")
	urlParams := mux.Vars(r)
	id := urlParams["uid"]
	ReadUser := User{}
	err := database.QueryRow("select * from users where uid = ?", id).Scan(&ReadUser.ID, &ReadUser.FirstName, &ReadUser.MiddleName, &ReadUser.LastName, &ReadUser.Email, &ReadUser.UserName, &ReadUser.Password)
	switch {
	case err == sql.ErrNoRows:
		fmt.Fprint(w, "No such User")
	case err != nil:
		log.Fatal(err.Error())
	default:
		output, _ := json.Marshal(ReadUser)
		fmt.Fprintf(w, string(output))
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	NewUser := User{}
	NewUser.FirstName = r.FormValue("first_name")
	NewUser.MiddleName = r.FormValue("middle_name")
	NewUser.LastName = r.FormValue("last_name")
	NewUser.UserName = r.FormValue("username")
	//NewUser.Password = r.FormValue("password")
	NewUser.Email = r.FormValue("email")
	output, err := json.Marshal(NewUser)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	sql := "INSERT	INTO	users	set	first_name='" + NewUser.FirstName + "', middle_name='" + NewUser.MiddleName + "',	last_name='" + NewUser.LastName + "',username='" + NewUser.UserName + "',email='" + NewUser.Email + "',password='" + NewUser.Password + "'"
	q, err := database.Exec(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(q)
}

func initialTest() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		james := User{1, "Ofonime", "Francis", "Usoro", "Jiggaseige", "all4usoro@gmail.com", "#%@&@#(((@#)))", true}
		output, err := json.Marshal(james)
		if err != nil {
			fmt.Println("Something went wrong")
		}
		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":4000", nil)
}
