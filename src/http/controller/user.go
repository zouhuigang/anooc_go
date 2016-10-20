package controller

import (
	"github.com/dchest/captcha"
	"github.com/labstack/echo"
	"github.com/polaris1119/logger"
	guuid "github.com/zouhuigang/uuid"
	"model"
	"net/http"
	"net/url"
	"util"
)

type UserController struct{}

// 保存uuid和email的对应关系（TODO:重启如何处理，有效期问题）
var regActivateCodeMap = map[string]string{}

//注册路由
func (this *UserController) RegisterRoute(g *echo.Group) {
	g.Any("/register", this.Register)
	g.Any("/login", this.login)
	g.Any("/emailexit", this.Emailexit)

}

// 注册用户页面
func (this UserController) Register(ctx echo.Context) error {
	data := map[string]interface{}{} //php数组相似的类型
	registerTpl := "user/registration.html,common/template.html"
	// 请求注册页面
	if ctx.Request().Method() != "POST" {
		return render(ctx, registerTpl, map[string]interface{}{"captchaId": captcha.NewLen(4)})
	}
	if model.Publics.Trim(ctx.FormValue("email")) == "" {
		return model.Publics.ReturnJson(ctx, 501, "邮箱不能为空", data)
	}
	if !model.Publics.IsEmail(ctx.FormValue("email")) {
		return model.Publics.ReturnJson(ctx, 501, "亲,邮箱格式不对哟。", data)
	}
	if model.Publics.Trim(ctx.FormValue("passwd")) == "" {
		return model.Publics.ReturnJson(ctx, 501, "密码不能为空", data)
	}

	data["email"] = ctx.FormValue("email")
	data["captchaId"] = captcha.NewLen(4)
	fields := []string{"email", "passwd"}
	form := url.Values{}
	for _, field := range fields {
		form.Set(field, ctx.FormValue(field))
	}
	form.Set("username", ctx.FormValue("email"))
	//{"data":{"email":["952750121@qq.com"],"password":["123456"]},"info":"form数据","status":501}

	// 入库
	errMsg, err := model.Users.CreateUser(ctx, form)
	if err != nil {
		return model.Publics.ReturnJson(ctx, 501, errMsg, err)
	}

	// 校验验证码
	//if !captcha.VerifyString(ctx.FormValue("captchaid"), ctx.FormValue("captchaSolution")) {
	//	return model.Publics.ReturnJson(ctx, 200, "验证码错误", data)
	//	}
	// 需要检验邮箱的正确性
	email := ctx.FormValue("email")
	uuid := this.genUUID(email)
	go model.Publics.SendActivateMail(email, uuid)
	return model.Publics.ReturnJson(ctx, 200, "注册成功", data)

}

func (UserController) genUUID(email string) string {
	var uuid string
	for {
		uuid = guuid.NewV4().String()
		if _, ok := regActivateCodeMap[uuid]; !ok {
			regActivateCodeMap[uuid] = email
			break
		}
		logger.Errorln("GenUUID 冲突....")
	}
	return uuid
}

//检测邮箱是否被注册
func (UserController) Emailexit(ctx echo.Context) error {
	data := map[string]interface{}{}
	data["info"] = "邮箱已被注册.."
	return success(ctx, data)
}

//登录
func (UserController) login(ctx echo.Context) error {
	data := map[string]interface{}{}
	loginTpl := "user/login.html,common/template.html"

	if _, ok := ctx.Get("anooc_user").(*model.Me); ok {
		return ctx.Redirect(http.StatusSeeOther, "/")
	}

	if ctx.Request().Method() != "POST" {
		return render(ctx, loginTpl, data)
	}

	//用户登录
	// 处理用户登录
	passwd := ctx.FormValue("passwd")
	username := ctx.FormValue("username")
	userLogin, err := model.Users.Login(ctx, username, passwd)
	if err != nil {
		data["username"] = username
		data["error"] = err.Error()

		if util.IsAjax(ctx) {
			return fail(ctx, 1, err.Error())
		}

		return render(ctx, loginTpl, data)
	}

	// 登录成功，设置cookie
	model.Publics.SetCookie(ctx, userLogin.Username)

	return model.Publics.ReturnJson(ctx, 200, "登录成功", data)

}
