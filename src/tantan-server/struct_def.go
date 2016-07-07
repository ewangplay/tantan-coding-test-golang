package main

//Tantan user struct
type TT_User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

//Tantan user relationship struct
type TT_Relationship struct {
	Peer_user_id string `json:"user_id"`
	State   string `json:"state"`
	Type    string `json:"type"`
}
