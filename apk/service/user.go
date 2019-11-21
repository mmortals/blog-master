package service

import (
	"blog-master/apk/dao"
	"blog-master/apk/model"
	"blog-master/public"
	"strconv"
)

type userService struct{}

var UserService = new(userService)

func (*userService) Register(c *public.MyfContext) {
	var user model.User
	username := c.Gin.Request.FormValue("username")
	nickName := c.Gin.Request.FormValue("nickName")
	password := c.Gin.Request.FormValue("password")
	email := c.Gin.Request.FormValue("email")
	user.UserName = username
	user.NickName = nickName
	user.Password = password
	user.Email = email
	_, err := dao.UserDao.FindUserByUsername(c, username)
	if nil != err {
		return
	}
	_ = dao.UserDao.AddUser(c, user)

}

func (*userService) login(c *public.MyfContext) {
	dao.UserDao.Login(c)
}

func (*userService) Delete(c *public.MyfContext) {
	id, _ := strconv.Atoi(c.Gin.Param("id"))
	_ = dao.UserDao.DeleteById(c, int16(id))
}
