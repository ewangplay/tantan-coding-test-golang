package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func GetRelationshipsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]

	result := fmt.Sprintf("GetRelationShipsHandler %v OK!", user_id)

	io.WriteString(w, result)
}

func SetRelationshipsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	peer_user_id := vars["peer_user_id"]

	result := fmt.Sprintf("PutRelationShipsHandler %v:%v OK!", user_id, peer_user_id)

	io.WriteString(w, result)
}
