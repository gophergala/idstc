package main

import (
	"fmt"
	"net/http"
	"os"

	"idstc/model"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/unrolled/render"
)

type ctx struct {
	Database *sqlx.DB
	r        *render.Render
	User     *model.User
}

func main() {

	store := cookiestore.New([]byte(os.Getenv("cookiestore")))

	//init db
	//db, err := sqlx.Connect("mysql", os.Getenv("connectionstring"))

	//if err != nil {
	//log.Print(err)
	//log.Fatal("Error Initializing Database...")
	//}

	//setup render
	ren := render.New(render.Options{
		Layout:        "shared/layout",
		IndentJSON:    true,
		IsDevelopment: true,
	})

	session := ctx{nil, ren, nil}

	//create routers
	router := mux.NewRouter()

	n := negroni.Classic()

	//register session middleware
	n.Use(sessions.Sessions("idstc", store))
	n.UseHandler(router)

	//not found handler
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	//open routes
	router.HandleFunc("/", session.Home)
	router.HandleFunc("/people", session.People)
	router.HandleFunc("/sprints", session.Sprints)
	router.HandleFunc("/sprints/detail/{id:[0-9]+}", session.SprintDetail)
	router.HandleFunc("/account", session.Account)
	router.HandleFunc("/analytics", session.Analytics)
	router.HandleFunc("/sprints/tasks/add/{id:[0-9]+}", session.AddTask).Methods("GET")
	router.HandleFunc("/sprints/add/{id:[0-9]+}", session.AddSprint).Methods("GET")

	n.Run(
		fmt.Sprint(":", os.Getenv("PORT")),
	)
}
