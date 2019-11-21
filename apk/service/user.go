package service

import (
	"blog-master/apk/dao"
	"blog-master/public"
)

type userService struct{}

var UserService = new(userService)

func (*userService) Register(c *public.MyfContext) {
	var msg string
	username := c.Gin.Request.FormValue("username")

	userSli := dao.UserDao.FindUserByUsername(c, username)

	dao.UserDao.Register(c)

}

func (*userService) login(c *public.MyfContext) {
	dao.UserDao.Login(c)
}

func (*userService) Delete(c *public.MyfContext) {
	dao.UserDao.Delete(c)
}
