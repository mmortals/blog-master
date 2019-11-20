package service

import (
	"blog-master/apk/dao"
	"blog-master/public"
)

type userService struct{}

var UserService = new(userService)

func (*userService) Register(c *public.MyfContext) {

	dao.UserDao.Register(c)

}

func (*userService) login(c *public.MyfContext) {
	dao.UserDao.Login(c)
}

func (*userService) Delete(c *public.MyfContext) {
	dao.UserDao.Delete(c)
}
