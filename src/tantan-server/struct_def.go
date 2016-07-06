package main

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Relationship struct {
	User_id string `json:"user_id"`
	State   string `json:"state"`
	Type    string `json:"type"`
}
