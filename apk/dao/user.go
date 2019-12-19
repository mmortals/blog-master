package dao

import (
	"blog-master/apk/db"
	"blog-master/apk/model"
	"blog-master/public"
	"crypto/md5"
	"fmt"
)

type userDao struct{}

var UserDao = new(userDao)

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
	var sql = "insert into user(username,nickname,password,enabled,email,userface) values (?,?,?,?,?,?)"
	result, err := dbConn.Exec(c.Context, sql, user.UserName, user.NickName, password, 0, user.Email, user.UserFace)
	if result > 0 {
		fmt.Printf(">>>>> 注册成功<<<<<")
	}
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
	var user_arr = []*model.User{}
	var sql = "select  id,username,nickname,password,enabled,email,userface,regTime from user where id=?"
	if err = dbConn.Query(c.Context, &user_arr, sql, id); nil != err {
		return
	}
	if len(user_arr) > 0 {
		user = *user_arr[0]
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
	var user_arr = []*model.User{}
	var sql = "select username,nickname,password,enabled,email,userface from user where username=?"
	if err = dbConn.Query(c.Context, user_arr, sql, username); nil != err {
		return
	}
	if len(user_arr) > 0 {
		user = *user_arr[0]
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
	_, err = dbConn.Exec(c.Context, "delete from user where id=?", id)
	if nil != err {
		return err
	}
	err = dbConn.Commit(c.Context)
	if nil != err {
		return err
	}
	return err
}
