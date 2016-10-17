package model

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"regexp"
	"strings"
)

type PublicsModel struct{}

var Publics = PublicsModel{}

func (this PublicsModel) IsEmail(email string) bool {
	regEmail := regexp.MustCompile("^\\w+@\\w+\\.\\w{2,4}$")
	return regEmail.MatchString(email)
}

//返回json格式
//200 – 表示成功
//500 – 表示服务器错误
//501 – 表示请求错误
//502 – 表示Token错误或者过时
//
func (this PublicsModel) ReturnJson(ctx echo.Context, status int, info string, data interface{}) error {
	result := map[string]interface{}{
		"status": status,
		"info":   info,
		"data":   data,
	}

	b, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return ctx.JSONBlob(http.StatusOK, b)
}

//去除空字符串
func (this PublicsModel) Trim(str string) string {
	// 去除空格
	str = strings.Replace(str, " ", "", -1)

	return str
}
