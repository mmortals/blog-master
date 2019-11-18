package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 单个Mysql连接
type DBConn struct {
	gorm *gorm.DB
	tx   bool // 是否是事务
}

func newDbConnection() (dbConn *DBConn, err error) {

	conn, err := gorm.Open("mysql", "root:Mimajiushi2.@106.54.101.82:3306/vueblog2?charset=utf8&parseTime=True&loc=Local")

	if nil != err {
		fmt.Println(err)
		return
	}
	dbConn = &DBConn{
		gorm: conn,
	}
	return
}

// 开启事务
func (dbConn *DBConn) Begin(context context.Context) (txConn *DBConn, err error) {
	if dbConn.tx {
		return nil, errors.New("tx error")
	}
	clone := *dbConn
	clone.gorm = dbConn.gorm.BeginTx(context, nil)
	clone.tx = true
	txConn = &clone
	return
}

// 提交事务
func (dbConn *DBConn) Commit(context context.Context) (err error) {
	if !dbConn.tx {
		return errors.New("tx error")
	}
	dbConn.gorm.Commit()
	return
}

// 回滚事务
func (dbConn *DBConn) Rollback(context context.Context) (err error) {
	if !dbConn.tx {
		return errors.New("tx error")
	}
	dbConn.gorm.Rollback()
	return
}

//func (dbConn *DBConn) Query(context context.Context, result interface{}, sql string, values ...interface{}) (err error) {

//var rows *sql2.Rows
//if rows, err = dbConn.gorm.Raw(sql, values...).Rows(); err != nil {
//	return
//	}
//for rows.Next() {
//elem := reflect.New(type3.Elem())                                   // 创建*Element
//if err = dbConn.gorm.ScanRows(rows, elem.Interface()); err != nil { // 填充*Element
//	return
//}
//newSlice := reflect.Append(reflect.ValueOf(result).Elem(), elem) // 将*Element追加到*result
//reflect.ValueOf(result).Elem().Set(newSlice)                     // 将新slice赋值给*result
//}
//return

//}
