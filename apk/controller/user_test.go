package controller

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestUserController_Register(t *testing.T) {
	var url = "http://localhost:8080/user/register"
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader("username=chenweiliang&password=mimajiushi2&nickName=陈承橙&email=wycdahaoren@gmail.com"))
	if nil != err {
		fmt.Println(err)
		log.Fatal(">>>Http Request Error<<<")
	}
	fmt.Println(resp.Request.RequestURI)
}

func Test_md5(t *testing.T) {
	password := "duanzhengchun"

	fmt.Println(password)
}

func TestUserController_Delete(t *testing.T) {

	var url = "http://localhost:8080/user/delete/26"
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		nil)
	if nil != err {
		fmt.Println(err)
		log.Fatal(">>>Http Request Error<<<")
	}
	fmt.Println(resp.Request.RequestURI)
}
