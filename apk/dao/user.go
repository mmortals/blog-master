package dao

import (
	"blog-master/apk/db"
	"blog-master/apk/model"
	"blog-master/public"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userDao struct{}

var UserDao = new(userDao)

func (*userDao) Register(c *public.MyfContext) {

	dbConn, err := db.NewDbConnection()

	if nil != err {
		return
	}
	username := c.Gin.Request.FormValue("username")
	nickName := c.Gin.Request.FormValue("nickName")
	password := c.Gin.Request.FormValue("password")
	email := c.Gin.Request.FormValue("email")
	userface := "http://www.baidu.com/"
	_, err = dbConn.Begin(c.Context)
	if nil != err {
		return
	}
	_, err = dbConn.Insert(c.Context, "insert into user(username,nickname,password,enabled,email,userface) values (?,?,?,?,?,?)",
		username, nickName, password, 0, email, userface)

	if nil != err {
		fmt.Println(err)
		return

	}
	err = dbConn.Commit(c.Context)
	if nil != err {
		return

	}
	msg := fmt.Sprintf("insert successful %d", username)
	c.Gin.JSON(http.StatusOK, gin.H{
		"message": msg,
	})

}

func (*userDao) Login(c *public.MyfContext) {

}

func (*userDao) AddUser(c *public.MyfContext, user model.User) error {
	dbConn, err := db.NewDbConnection()
	if nil != err {
		return err
	}
	_, err = dbConn.Begin(c.Context)
	if nil != err {
		return err
	}
	_, err = dbConn.Insert(c.Context, "insert into user(username,nickname,password,enabled,email,userface) values (?,?,?,?,?,?)",
		user.UserName, user.NickName, user.Password, 0, user.Email, user.UserFace)

	if nil != err {
		fmt.Println(err)
		return err
	}
	err = dbConn.Commit(c.Context)
	if nil != err {
		return err
	}
	return err
}

func (*userDao) FindUserById(c *public.MyfContext, id int16) (model.User, error) {
	var user model.User
	dbConn, err := db.NewDbConnection()
	if err != nil {
		return user, err
	}
	_, err = dbConn.Begin(c.Context)
	if err != nil {
		return user, err
	}
	err = dbConn.QueryRow("select username,nickname,password,enabled,email,userface from user where id=?", id).Scan(&user.Id,
		&user.UserName, &user.NickName, &user.UserFace, &user.Password, &user.Enable, &user.Email, &user.RegTime)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	err = dbConn.Commit(c.Context)
	return user, err
}

func (*userDao) FindUserByUsername(c *public.MyfContext, username string) (model.User, error) {
	var user model.User
	dbConn, err := db.NewDbConnection()
	if err != nil {
		return user, err
	}
	_, err = dbConn.Begin(c.Context)
	if err != nil {
		return user, err
	}
	err = dbConn.QueryRow("select username,nickname,password,enabled,email,userface from user where username=?", username).Scan(&user.Id,
		&user.UserName, &user.NickName, &user.UserFace, &user.Password, &user.Enable, &user.Email, &user.RegTime)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	err = dbConn.Commit(c.Context)
	return user, err
}

func (*userDao) FindUsers(c *public.MyfContext) ([]model.User, error) {
	users := make([]model.User, 0)
	return users, nil
}

func (*userDao) DeleteById(c *public.MyfContext) error {
	dbConn, err := db.NewDbConnection()
	if nil != err {
		return err
	}
	id := c.Gin.Param("id")
	_, err = dbConn.Begin(c.Context)
	if nil != err {
		return err
	}
	_, err = dbConn.Delete(c.Context, "delete from user where id=?", id)
	if nil != err {
		return err
	}
	err = dbConn.Commit(c.Context)
	if nil != err {
		return err
	}
	return err
}
