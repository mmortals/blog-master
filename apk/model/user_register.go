package model

type UserRegister struct {
	id       int
	UserName string `grom:"default:'galeone'"`
	NickName string
	Age      int
	Password int
}
