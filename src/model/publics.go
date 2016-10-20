package model

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/polaris1119/config"
	"github.com/polaris1119/goutils"
	"github.com/zouhuigang/logger"
	"github.com/zouhuigang/sessions"
	"net/http"
	"net/smtp"
	"regexp"
	"strings"
	"time"
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

// SendMail 发送电子邮件,subject:标题,content内容,tos目标对象邮箱
func (this PublicsModel) SendMail(subject string, content string, tos []string) error {
	emailConfig, _ := config.ConfigFile.GetSection("email")
	mailtype := "html"
	content_type := "Content-Type: text/" + mailtype + "; charset=UTF-8"
	//message := `From: 安橙网|永久免费的娱乐学习平台<` + emailConfig["from_email"] + `>To: ` + strings.Join(tos, ",") + `Subject: ` + subject + `Content-Type: text/html;charset=UTF-8` + content
	message := "To: " + strings.Join(tos, ",") +
		"\r\nFrom:安橙网|永久免费的娱乐学习平台<" + emailConfig["from_email"] +
		">\r\nSubject: " + subject + "\r\n" +
		content_type + "\r\n\r\n" + content

	smtpAddr := emailConfig["smtp_host"] + ":" + emailConfig["smtp_port"]
	auth := smtp.PlainAuth("", emailConfig["smtp_username"], "lioeopjnwvnebefd", emailConfig["smtp_host"])
	err := smtp.SendMail(smtpAddr, auth, emailConfig["from_email"], tos, []byte(message))
	if err != nil {
		logger.Errorln("Send Mail to", strings.Join(tos, ","), "error:", err)
		return err
	}
	logger.Infoln("Send Mail to", strings.Join(tos, ","), "Successfully")
	return nil
}

// SendActivateMail 发送激活邮件
func (this PublicsModel) SendActivateMail(email, uuid string) {

	timestamp := time.Now().Unix()
	sign := this.genActivateSign(email, uuid, timestamp)
	fmt.Println("进入了发送邮件流程")
	param := goutils.Base64Encode(fmt.Sprintf("uuid=%s&timestamp=%d&sign=%s", uuid, timestamp, sign))

	domain := config.ConfigFile.MustValue("global", "domain")
	activeUrl := fmt.Sprintf("http://%s/account/activate?param=%s", domain, param)

	content := `
尊敬的anooc用户：<br/><br/>
欢迎您注册安橙网，请点击下面的地址激活你的帐号（有效期4小时）：<br/><br/>
<a href="` + activeUrl + `">` + activeUrl + `</a><br/><br/>
<div style="text-align:right;">&copy;2014-2017 anooc.com </div>`
	this.SendMail("安橙网邮箱激活", content, []string{email})
}

func (this PublicsModel) genActivateSign(email, uuid string, ts int64) string {
	emailSignSalt := config.ConfigFile.MustValue("security", "activate_sign_salt")
	origStr := fmt.Sprintf("uuid=%semail=%stimestamp=%d%s", uuid, email, ts, emailSignSalt)
	return goutils.Md5(origStr)
}

//session
var Store = sessions.NewCookieStore([]byte(config.ConfigFile.MustValue("global", "cookie_secret")))

func (this PublicsModel) SetCookie(ctx echo.Context, username string) {
	Store.Options.HttpOnly = true

	session := GetCookieSession(ctx)
	if ctx.FormValue("remember_me") != "1" {
		// 浏览器关闭，cookie删除，否则保存30天(github.com/gorilla/sessions 包的默认值)
		session.Options = &sessions.Options{
			Path:     "/",
			HttpOnly: true,
		}
	}
	session.Values["username"] = username
	req := Request(ctx)
	resp := ResponseWriter(ctx)
	session.Save(req, resp)
}

// 必须是 http.Request
func GetCookieSession(ctx echo.Context) *sessions.Session {
	session, _ := Store.Get(Request(ctx), "anooc_user")
	return session
}

func Request(ctx echo.Context) *http.Request {
	return ctx.Request().(*standard.Request).Request
}

func ResponseWriter(ctx echo.Context) http.ResponseWriter {
	return ctx.Response().(*standard.Response).ResponseWriter
}
