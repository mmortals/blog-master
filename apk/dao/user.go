package dao

import (
	"blog-master/apk/db"
	"blog-master/apk/model"
	"blog-master/public"
	"fmt"
	"github.com/gin-gonic/gin"
)

type userDao struct{}

var UserDao = new(userDao)

func (*userDao) Regist(c *public.MyfContext) {

	var url = c.Gin.Request.RequestURI
	id := c.Gin.Param("id")

	dbConn, err := db.NewDbConnection()
	if nil != err {
		return
	}
	dbConn.Begin(c.Context)
	rows, err := dbConn.Query(c.Context, "select * from user where id = ?", id)
	dbConn.Commit(c.Context)
	if nil != err {
		fmt.Println(err)
	}
	users := make([]model.UserRegister, 0, 5)
	for rows.Next() {
		user := new(model.UserRegister)
		err := rows.Scan(&user.Id, &user.UserName, &user.NickName, &user.Password, &user.Enable, &user.Email, &user.UserFace, &user.RegTime)
		if nil != err {
			fmt.Println(err)
		}
		users = append(users, *user)
	}
	fmt.Printf(">>>>>url is :<<<<<", url)

	c.Gin.JSON(200, gin.H{
		"message": url,
		"result":  users,
	})
}
