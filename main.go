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

type Users struct {
	Users []User `json:"users"`
}

type CreateResponse struct {
	Error string `json:"error"`
}

var database *sql.DB
var err error

func main() {
	db, err := sql.Open("mysql", "root:password@/database_name")
	if err != nil {
		log.Fatal(err.Error())
	}
	database = db
	routes := mux.NewRouter()
	routes.HandleFunc("/api/users/{key:[A-Za-z0-9\\-]}", UserRetrieve)
	routes.HandleFunc("/api/users", UserCreate).Methods("POST")
	routes.HandleFunc("/api/users", UserRetrieve).Methods("GET")
	http.Handle("/", routes)
	http.ListenAndServe(":4000", nil)
}

func UserRetrieve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Pragma", "no-cache")
	//variables := mux.Vars(r)
	//key := variables["key"]
	rows, _ := database.Query("select * from users LIMIT 10")
	Response := Users{}
	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.MiddleName, &user.Email, &user.UserName, &user.IsAdmin, &user.Password)
		Response.Users = append(Response.Users, user)
	}
	output, _ := json.Marshal(Response)
	fmt.Fprintln(w, string(output))
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

func UserCreate(w http.ResponseWriter, r *http.Request) {
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

	Response := CreateResponse{}
	sql := "INSERT	INTO	users	set	first_name='" + NewUser.FirstName + "', middle_name='" + NewUser.MiddleName + "',	last_name='" + NewUser.LastName + "',username='" + NewUser.UserName + "',email='" + NewUser.Email + "',password='" + NewUser.Password + "'"
	q, err := database.Exec(sql)
	if err != nil {
		Response.Error = err.Error()
	}
	fmt.Println(q)
	createOutput, _ := json.Marshal(Response)
	fmt.Fprintln(w, string(createOutput))
}
