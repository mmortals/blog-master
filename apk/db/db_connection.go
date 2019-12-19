package db

import (
	"bytes"
	"context"
	sql2 "database/sql"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"reflect"
	"strings"
)

// 单个Mysql连接
type DBConn struct {
	gorm *gorm.DB
	tx   bool // 是否是事务
}

// 错误码
var (
	ERR_DB_GROUP_NOT_FOUND   = errors.New("此DB不存在")
	ERR_DB_CONN_NOT_FOUND    = errors.New("没有可用DB连接")
	ERR_QUERY_RESULT_INVALID = errors.New("result传参类型必须是*[]*ElemType")
	ERR_RECURSION_TX         = errors.New("嵌套开启了事务")
	ERR_INVALID_TX           = errors.New("非事务不能提交或回滚")
)

//查找
func (dbConn *DBConn) Query(context context.Context, result interface{}, sql string, values ...interface{}) (err error) {

	var (
		type1, type2, type3 reflect.Type
	)
	if type1 = reflect.TypeOf(result); type1.Kind() != reflect.Ptr { // type1是*[]*Element
		err = ERR_QUERY_RESULT_INVALID
		return
	}
	if type2 = type1.Elem(); type2.Kind() != reflect.Slice { // type2是[]*Element
		err = ERR_QUERY_RESULT_INVALID
		return
	}
	if type3 = type2.Elem(); type3.Kind() != reflect.Ptr { // type3是*Element
		err = ERR_QUERY_RESULT_INVALID
		return
	}
	var rows *sql2.Rows
	if rows, err = dbConn.gorm.Raw(sql, values...).Rows(); err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		elem := reflect.New(type3.Elem())                                   // 创建*Element
		if err = dbConn.gorm.ScanRows(rows, elem.Interface()); err != nil { // 填充*Element
			return
		}
		newSlice := reflect.Append(reflect.ValueOf(result).Elem(), elem) // 将*Element追加到*result
		reflect.ValueOf(result).Elem().Set(newSlice)                     // 将新slice赋值给*result
	}
	return
}

func (dbConn *DBConn) QueryRow(context context.Context, sql string, values ...interface{}) (result2 *sql2.Row) {

	return dbConn.gorm.Raw(sql, values...).Row()
}

//func (dbConn *DBConn) Insert(context context.Context, sql string, values ...interface{}) (result2 *sql2.Result, err error) {
//	var result sql2.Result
//	if result, err = dbConn.gorm.CommonDB().Exec(sql, values...); nil != err {
//
//		return nil, err
//	}
//	return &result, nil
//}
//
//func (dbConn *DBConn) Delete(context context.Context, sql string, values ...interface{}) (result2 *sql2.Result, err error) {
//	var result sql2.Result
//	if result, err = dbConn.gorm.CommonDB().Exec(sql, values...); nil != err {
//
//		return nil, err
//	}
//	return &result, nil
//}
//
//func (dbConn *DBConn) Update(context context.Context, sql string, values ...interface{}) (result2 *sql2.Result, err error) {
//	var result sql2.Result
//	if result, err = dbConn.gorm.CommonDB().Exec(sql, values...); nil != err {
//
//		return nil, err
//	}
//	return &result, nil
//}
//SQL 写入
func (dbConn *DBConn) Exec(context context.Context, sql string, values ...interface{}) (result int64, err error) {

	var sqlResult sql2.Result
	sqlType := dbConn.sqlType(sql)

	// 执行SQL
	if sqlResult, err = dbConn.gorm.CommonDB().Exec(sql, values...); err != nil {
		return
	}

	// 判断SQL类型取不同结果
	if sqlType == "INSERT" {
		result, err = sqlResult.LastInsertId()
	} else {
		result, err = sqlResult.RowsAffected()
	}

	return
}

func NewDbConnection() (dbConn *DBConn, err error) {

	conn, err := gorm.Open("mysql", "root:Mimajiushi2.@(106.54.101.82:3306)/vueblog2?charset=utf8&parseTime=True&loc=Local")

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
	_ = dbConn.gorm.Close()
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

// 判断SQL类型
func (dbConn *DBConn) sqlType(sql string) string {
	sql = strings.TrimLeft(sql, " \t\r\n")

	buf := bytes.Buffer{}
	for i := 0; i < len(sql); i++ {
		if sql[i] != ' ' && sql[i] != '\t' && sql[i] != '\r' && sql[i] != '\n' {
			buf.WriteByte(sql[i])
		} else {
			break
		}
	}
	return strings.ToUpper(buf.String())
}
