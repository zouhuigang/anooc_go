package controller

import (
	"github.com/labstack/echo"
)

type UserController struct{}

//注册路由
func (this *UserController) RegisterRoute(g *echo.Group) {
	g.Get("/register", this.register)
	g.Get("/login", this.login)
}

// 注册用户
func (UserController) register(ctx echo.Context) error {

	data := map[string]interface{}{}
	return render(ctx, "user/registration.html,common/template.html", data)
}

//登录
func (UserController) login(ctx echo.Context) error {

	data := map[string]interface{}{}
	return render(ctx, "user/login.html,common/template.html", data)
}
