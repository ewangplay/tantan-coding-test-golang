package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//SQL Statements
const (
	GET_USER_RELATIONSHIP_SQL          string = `SELECT peer_user_id,state,type from relationship_tbl WHERE user_id=?`
	GET_STATE_ONE_TO_ONE               string = `SELECT state from relationship_tbl WHERE user_id=? AND peer_user_id = ?`
	UPDATE_STATE_ONE_TO_ONE            string = `UPDATE relationship_tbl SET state = ? WHERE user_id=? AND peer_user_id=?`
	CREATE_OR_UPDATE_USER_RELATIONSHIP string = `INSERT INTO relationship_tbl(user_id, peer_user_id, state, type) VALUES(?, ?, ?, ?) ON CONFLICT (user_id, peer_user_id) DO UPDATE SET state = ?`
)

///////////////////////////////////////////////////////////////////////////
//Get user relationships handler
func GetRelationshipsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]

	var result string
	var err error

	result, err = GetUserRelationship(user_id)
	if err != nil {
		result = fmt.Sprintf("[ERROR] %v", err)
	}

	io.WriteString(w, result)
}

func GetUserRelationship(user_id string) (result string, err error) {
	var relationships []TT_Relationship

	err = g_pgAdaptor.Query(&relationships, GET_USER_RELATIONSHIP_SQL, user_id)
	if err != nil {
		return "", err
	}

	rb, err := json.Marshal(relationships)
	if err != nil {
		return "", err
	}

	return string(rb), nil
}

//////////////////////////////////////////////////////////////////////////
//Set user relationships handler
func SetRelationshipsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	peer_user_id := vars["peer_user_id"]

	var result string
	var err error

	result, err = SetUserRelationship(user_id, peer_user_id, r)
	if err != nil {
		result = fmt.Sprintf("[ERROR] %v", err)
	}

	io.WriteString(w, result)
}

func SetUserRelationship(user_id, peer_user_id string, r *http.Request) (result string, err error) {
	//read the body data
	body, err := ioutil.ReadAll(r.Body)
	if err != nil && err != io.EOF {
		return "", err
	}

	//parse the body data
	var rs TT_Relationship
	rs.Peer_user_id = peer_user_id
	rs.Type = "relationship"

	err = json.Unmarshal(body, &rs)
	if err != nil {
		return "", err
	}

	//check the data field
	if rs.State == "" {
		return "", fmt.Errorf("state field empty!")
	}

	//Retrieve the state for peer_user_id to user_id
	var peer_rs TT_Relationship
	err = g_pgAdaptor.QueryOne(&peer_rs, GET_STATE_ONE_TO_ONE, peer_user_id, user_id)
	if err != nil {
		if !strings.Contains(err.Error(), "no rows in result set") {
			return "", err
		}
	} else {

		//Determine the state
		var isStateChanged bool
		if rs.State == "liked" && peer_rs.State == "liked" {
			rs.State = "matched"
			peer_rs.State = "matched"
			isStateChanged = true
		} else if rs.State == "disliked" && peer_rs.State == "matched" {
			peer_rs.State = "liked"
			isStateChanged = true
		}

		//Update the state for peer_user_id to user_id if matched
		if isStateChanged {
			err = g_pgAdaptor.Exec(UPDATE_STATE_ONE_TO_ONE, peer_rs.State, peer_user_id, user_id)
			if err != nil {
				return "", err
			}
		}

	}

	//Create or update the state for user_id to peer_user_id
	err = g_pgAdaptor.Exec(CREATE_OR_UPDATE_USER_RELATIONSHIP, user_id, peer_user_id, rs.State, rs.Type, rs.State)
	if err != nil {
		return "", err
	}

	rb, err := json.Marshal(&rs)
	if err != nil {
		return "", err
	}

	return string(rb), nil
}
