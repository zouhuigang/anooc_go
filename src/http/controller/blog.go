package controller

import (
	"github.com/labstack/echo"
)

type BlogController struct{}

//注册路由
func (this *BlogController) RegisterRoute(g *echo.Group) {
	g.Get("/blog", this.index)
}

// Create 新建主题
func (BlogController) index(ctx echo.Context) error {

	data := map[string]interface{}{}
	return render(ctx, "blog/welcome.html,common/template.html", data)
}
