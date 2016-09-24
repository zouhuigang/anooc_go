package controller

import (
	"github.com/labstack/echo"
	"logic"
	"model"
)

type BlogController struct{}

//注册路由
func (this *BlogController) RegisterRoute(g *echo.Group) {
	g.Get("/blog", this.index)
}

// Create 新建主题
func (BlogController) index(ctx echo.Context) error {
	nodes := logic.GenNodes()

	title := ctx.FormValue("title")
	// 请求新建主题页面
	if title == "" || ctx.Request().Method() != "POST" {
		return render(ctx, "blog/index.html", map[string]interface{}{"nodes": nodes, "activeTopics": "active"})
	}

	me := ctx.Get("user").(*model.Me)
	err := logic.DefaultTopic.Publish(ctx, me, ctx.FormParams())
	if err != nil {
		return fail(ctx, 1, "内部服务错误")
	}

	return success(ctx, nil)
}
