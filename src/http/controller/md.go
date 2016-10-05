package controller

import (
	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
	"model"
)

type MdController struct{}

//注册路由
func (this *MdController) RegisterRoute(g *echo.Group) {
	g.Get("/md", this.welcome)
	g.Get("/view/:id", this.View)
}

func (MdController) welcome(ctx echo.Context) error {
	const (
		md = ``
	)

	unsafe := blackfriday.MarkdownCommon([]byte(md))
	outputHtml := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	data := map[string]interface{}{
		"output": template.HTML(outputHtml),
	}
	return render(ctx, "markdown/welcome.html", data)

}

func (MdController) View(ctx echo.Context) error {
	md, err := model.MdModelFindById(ctx, string(ctx.Param("id")))
	if err != nil {
		return render(ctx, "markdown/view.html", map[string]interface{}{})
	}

	unsafe := blackfriday.MarkdownCommon([]byte(md.Content))
	outputHtml := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	data := map[string]interface{}{
		"output": template.HTML(outputHtml),
	}
	return render(ctx, "markdown/view.html,common/header.html", data)
}
