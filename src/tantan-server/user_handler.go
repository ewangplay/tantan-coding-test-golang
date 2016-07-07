package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var g_userIdSeed int64

const (
	GET_ALL_USERS_SQL string = `SELECT id,name,type from user_tbl`
	ADD_USER_SQL      string = `INSERT INTO user_tbl(id, name, type) VALUES(?,?,?)`
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	var err error

	switch r.Method {
	case "GET":
		result, err = ListAllUsers()
	case "POST":
		result, err = AddUser(r)
	default:
		err = fmt.Errorf("unsupported http method: %v", r.Method)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		result = fmt.Sprintf("[ERROR] %v", err)
	} else {
		//Set the http properties
		header := w.Header()
		header.Add("Content-Type", "application/json")
		header.Add("charset", "UTF-8")
	}

	io.WriteString(w, result)
}

func ListAllUsers() (result string, err error) {

	var users []TT_User

	err = g_pgAdaptor.Query(&users, GET_ALL_USERS_SQL)
	if err != nil {
		return "", err
	}

	r, err := json.Marshal(users)
	if err != nil {
		return "", err
	}

	return string(r), nil
}

func AddUser(r *http.Request) (result string, err error) {
	//read the body data
	body, err := ioutil.ReadAll(r.Body)
	if err != nil && err != io.EOF {
		return "", err
	}

	//parse the body data
	var user TT_User
	user.Type = "user"
	err = json.Unmarshal(body, &user)
	if err != nil {
		return "", err
	}

	//check the data field
	if user.Name == "" {
		return "", fmt.Errorf("user name empty!")
	}

	//increment user id
	user.Id, err = GenerateUserID()
	if err != nil {
		return "", err
	}

	//insert data into db
	err = g_pgAdaptor.Exec(ADD_USER_SQL, user.Id, user.Name, user.Type)
	if err != nil {
		return "", err
	}

	rb, err := json.Marshal(&user)
	if err != nil {
		return "", err
	}

	return string(rb), nil
}

func GenerateUserID() (id string, err error) {
	//This just a temporary solution.
	//In fact, here need to lock，
	// and need to save the last value at somewhere when server exias.
	g_userIdSeed++
	return fmt.Sprintf("%v", g_userIdSeed), nil
}
