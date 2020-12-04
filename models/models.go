package models

type Users struct {
	Id       string `form:"id" json:"id"`
	name     string `form:"name" json:"name"`
	email    string `form:"email" json:"email"`
	password string `form:"password" json:"password"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}
