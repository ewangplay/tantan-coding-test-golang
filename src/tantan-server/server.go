package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
    "fmt"
)

var g_pgAdaptor *PgAdaptor

func main() {
    var err error

    //create PostGreSQL adaptor
    addr := "127.0.0.1:5432"
    user := "dbuser"
    pass := "dogtutu"
    dbname := "tantan_db"
    g_pgAdaptor, err = NewPgAdaptor(addr, user, pass, dbname)
    if err != nil {
        fmt.Printf("Create PostGreSQL adaptor error: %v\n", err)
        os.Exit(1)
    }
    defer g_pgAdaptor.Release()

    //create http request router
	router := mux.NewRouter()
	router.HandleFunc("/users", UsersHandler).Methods("GET", "POST")
	router.HandleFunc("/users/{user_id:[0-9]+}/relationships", GetRelationshipsHandler).Methods("GET")
	router.HandleFunc("/users/{user_id:[0-9]+}/relationships/{peer_user_id:[0-9]+}", SetRelationshipsHandler).Methods("PUT")

    //statup http server
	err = http.ListenAndServe(":8090", router)
	if err != nil {
        fmt.Printf("Startup http server error: %v\n", err)
		os.Exit(1)
	}
}
