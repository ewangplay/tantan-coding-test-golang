package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", UsersHandler)
	router.HandleFunc("/users/{user_id}/relationships", GetRelationShipsHandler)
	router.HandleFunc("/users/{user_id}/relationships/{other_user_id}", PutRelationShipsHandler)

	err := http.ListenAndServe(":8090", router)
	if err != nil {
		os.Exit(1)
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "UsersHandler OK!")
}

func GetRelationShipsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]
    result := fmt.Sprintf("GetRelationShipsHandler %v OK!", user_id)
	io.WriteString(w, result)
}

func PutRelationShipsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]
    other_user_id := vars["other_user_id"]
    result := fmt.Sprintf("PutRelationShipsHandler %v:%v OK!", user_id, other_user_id)

	io.WriteString(w, result)
}
