package model

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"regexp"
)

func PublicsIsEmail(email string) bool {
	regEmail := regexp.MustCompile("^\\w+@\\w+\\.\\w{2,4}$")
	return regEmail.MatchString(email)
}

//返回json格式
//200 – 表示成功
//500 – 表示服务器错误
//501 – 表示请求错误
//502 – 表示Token错误或者过时
//
func PublicsReturnJson(ctx echo.Context, status int, info string, data interface{}) error {
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
