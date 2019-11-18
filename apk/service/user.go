package service

import (
	"blog-master/apk/dao"
	"blog-master/public"
)

type userService struct{}

var UserService = new(userService)

func (*userService) Regist(c *public.MyfContext) {

	dao.UserDao.Regist(c)

}
