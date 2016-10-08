package controller

import (
	"github.com/labstack/echo"
	"model"
)

type BlogController struct{}

//注册路由
func (this *BlogController) RegisterRoute(g *echo.Group) {
	g.Get("/blog", this.index)
}

// Create 新建主题
func (BlogController) index(ctx echo.Context) error {

	lists, err := model.MdModelList(20)
	if err != nil {
		return render(ctx, "welcome/index.html", map[string]interface{}{})
	}

	data := map[string]interface{}{
		"lists": lists,
	}
	return render(ctx, "blog/welcome.html,common/template.html,common/newslist.html", data)
}
