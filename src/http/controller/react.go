package controller

import (
	"github.com/labstack/echo"
)

type ReactController struct{}

//注册路由
func (this *ReactController) RegisterRoute(g *echo.Group) {
	g.Get("/react", this.welcome)
}

func (ReactController) welcome(ctx echo.Context) error {
	data := map[string]interface{}{}
	return render(ctx, "react/welcome.html", data)

}
