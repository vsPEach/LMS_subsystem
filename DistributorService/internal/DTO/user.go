package DTO

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func NewUser() *User {
	return nil
}
