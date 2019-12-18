package dao

import (
	"blog-master/apk/db"
	"blog-master/apk/model"
	"blog-master/public"
	"crypto/md5"
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
	dbConn, err = dbConn.Begin(c.Context)
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

func (*userDao) Login(username string, password string, c *public.MyfContext) (user model.User, err error) {

	dbConn, err := db.NewDbConnection()
	if nil != err {
		return
	}
	dbConn, err = dbConn.Begin(c.Context)

	if nil != err {
		return
	}
	err = dbConn.QueryRow(c.Context, "select  id,username,nickname,password,enabled,email,userface,regTime from user where username=? and password=? ", username, password).Scan(&user.Id,
		&user.UserName, &user.NickName, &user.Password, &user.Enable, &user.Email, &user.UserFace, &user.RegTime)
	dbConn.Commit(c.Context)
	return
}

func (*userDao) AddUser(c *public.MyfContext, user model.User) error {
	dbConn, err := db.NewDbConnection()
	if nil != err {
		return err
	}
	dbConn, err = dbConn.Begin(c.Context)
	if nil != err {
		return err
	}
	password := fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
	_, err = dbConn.Insert(c.Context, "insert into user(username,nickname,password,enabled,email,userface) values (?,?,?,?,?,?)",
		user.UserName, user.NickName, password, 0, user.Email, user.UserFace)

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

func (*userDao) FindUserById(c *public.MyfContext, id int16) (user model.User, err error) {
	dbConn, err := db.NewDbConnection()
	if err != nil {
		return
	}
	dbConn, err = dbConn.Begin(c.Context)
	if err != nil {
		return
	}

	if result, err := dbConn.Query(c.Context, "select  id,username,nickname,password,enabled,email,userface,regTime from user where id=?", id); nil != err {
		return
	} else {
		if result.Next() {
			err = result.Scan(
				&user.Id,
				&user.UserName, &user.NickName, &user.Password, &user.Enable, &user.Email, &user.UserFace, &user.RegTime)
			if nil != err {
				return
			}
		}
	}

	err = dbConn.Commit(c.Context)

	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (*userDao) FindUserByUsername(c *public.MyfContext, username string) (user model.User, err error) {

	dbConn, err := db.NewDbConnection()
	if err != nil {
		return user, err
	}
	dbConn, err = dbConn.Begin(c.Context)
	if err != nil {
		return
	}
	var sql = "select username,nickname,password,enabled,email,userface from user where username=?"
	if result, err := dbConn.Query(c.Context, sql, username); nil != err {
		result.Scan(&user.Id,
			&user.UserName, &user.NickName, &user.UserFace, &user.Password, &user.Enable, &user.Email, &user.RegTime)
	}
	err = dbConn.Commit(c.Context)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (*userDao) FindUsers(c *public.MyfContext) ([]model.User, error) {
	users := make([]model.User, 0)
	return users, nil
}

func (*userDao) DeleteById(c *public.MyfContext, id int16) error {
	dbConn, err := db.NewDbConnection()
	if nil != err {
		return err
	}
	dbConn, err = dbConn.Begin(c.Context)
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
