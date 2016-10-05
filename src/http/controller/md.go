package controller

import (
	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
)

type MdController struct{}

//注册路由
func (this *MdController) RegisterRoute(g *echo.Group) {
	g.Get("/md", this.welcome)
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
