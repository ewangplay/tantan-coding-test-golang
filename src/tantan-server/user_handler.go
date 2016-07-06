package main

import (
	"fmt"
	"io"
	"net/http"
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
		result = fmt.Sprintf("[ERROR] %v", err)
	}

	io.WriteString(w, result)
}

func ListAllUsers() (result string, err error) {
	return "", nil
}

func AddUser(r *http.Request) (result string, err error) {
	return "", nil
}
