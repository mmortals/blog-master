package service

import (
	"blog-master/apk/dao"
	"blog-master/apk/model"
	"blog-master/public"
	"crypto/md5"
	"fmt"
	"strconv"
)

type userService struct{}

var UserService = new(userService)

func (*userService) Register(c *public.MyfContext) (err error) {
	var user model.User
	username := c.Gin.Request.FormValue("username")
	nickName := c.Gin.Request.FormValue("nickName")
	password := c.Gin.Request.FormValue("password")
	email := c.Gin.Request.FormValue("email")
	user.UserName = username
	user.NickName = nickName
	user.Password = password
	user.Email = email
	_, err = dao.UserDao.FindUserByUsername(c, username)
	if nil == err {

	}
	err = dao.UserDao.AddUser(c, user)
	if nil != err {
		fmt.Println(err)
	}
	return
}

func (*userService) Login(c *public.MyfContext) (user model.User, err error) {
	username := c.Gin.Request.FormValue("username")
	password := c.Gin.Request.FormValue("password")
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	user, err = dao.UserDao.Login(username, password, c)
	return
}

func (*userService) Delete(c *public.MyfContext) {
	id, _ := strconv.Atoi(c.Gin.Param("id"))
	_ = dao.UserDao.DeleteById(c, int16(id))
}
func (service *userService) FindUserById(c *public.MyfContext) (user model.User, err error) {

	userId, err := strconv.Atoi(c.Gin.Request.FormValue("userId"))
	if nil != err {

	}
	user, err = dao.UserDao.FindUserById(c, int16(userId))

	if nil != err {
		return
	}
	return
}
