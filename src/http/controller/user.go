package controller

import (
	"github.com/dchest/captcha"
	"github.com/labstack/echo"
	"model"
	"net/url"
)

type UserController struct{}

//注册路由
func (this *UserController) RegisterRoute(g *echo.Group) {
	g.Any("/register", this.Register)
	g.Get("/login", this.login)
	g.Any("/emailexit", this.Emailexit)

}

// 注册用户页面
func (UserController) Register(ctx echo.Context) error {
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
	if model.Publics.Trim(ctx.FormValue("password")) == "" {
		return model.Publics.ReturnJson(ctx, 501, "密码不能为空", data)
	}
	data["email"] = ctx.FormValue("email")
	data["captchaId"] = captcha.NewLen(4)
	fields := []string{"username", "email", "passwd"}
	form := url.Values{}
	for _, field := range fields {
		form.Set(field, ctx.FormValue(field))
	}

	// 入库
	errMsg, err := model.Users.CreateUser(ctx, form)
	if err != nil {
		// bugfix：http://studygolang.com/topics/255
		if errMsg == "" {
			errMsg = err.Error()
		}
		data["error"] = errMsg
		return render(ctx, registerTpl, data)
	}

	// 校验验证码
	if !captcha.VerifyString(ctx.FormValue("captchaid"), ctx.FormValue("captchaSolution")) {
		data["info"] = "验证码错误"
		return render(ctx, registerTpl, data)
	}

	return render(ctx, registerTpl, data)

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
	return render(ctx, "user/login.html,common/template.html", data)
}
