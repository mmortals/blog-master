package model

type User struct {
	Id       int
	UserName string `grom:"default:'galeone'"`
	NickName string
	UserFace string
	Password string
	Enable   int
	Email    string
	RegTime  string
}
