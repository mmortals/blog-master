package dao

import (
	"blog-master/apk/db"
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

func (*userDao) Delete(c *public.MyfContext) {
	dbConn, err := db.NewDbConnection()

	if nil != err {
		return
	}
	id := c.Gin.Param("id")
	_, err = dbConn.Begin(c.Context)
	if nil != err {
		return
	}
	_, err = dbConn.Delete(c.Context, "delete from user where id=?", id)
	if nil != err {
		return
	}
	err = dbConn.Commit(c.Context)
	if nil != err {
		return

	}
	c.Gin.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
