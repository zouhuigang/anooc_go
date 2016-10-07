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
	g.Get("/editor", this.Editor)
	g.Get("/view/:id", this.View)
}

func (MdController) Editor(ctx echo.Context) error {

	data := map[string]interface{}{}
	return render(ctx, "markdown/editor.html", data)

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
	return render(ctx, "markdown/view.html,common/template.html", data)
}
