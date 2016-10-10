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
	g.Match([]string{"GET", "POST"}, "/md/submitmd", this.SubmitMd)
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

//提交笔记
func (MdController) SubmitMd(ctx echo.Context) error {

	if ctx.FormValue("submit") == "1" {
		//user := ctx.Get("user").(*model.Me)
		id, err := model.MdModelAddUpdate(ctx, ctx.FormParams(), "邹慧刚")
		if err != nil {
			return fail(ctx, 1, "保存失败")
		}
		//return success(ctx, nil)
		data := map[string]interface{}{
			"id": id,
		}
		return success(ctx, data)
	}

	return fail(ctx, 2, "保存失败")

	//var data = make(map[string]interface{})
	//id := goutils.MustInt(ctx.QueryParam("id"))
	//if id != 0 {
	//	reading := logic.DefaultReading.FindById(ctx, id)
	//	if reading != nil {
	//		data["reading"] = reading
	//	}
	//}

	//return render(ctx, "reading/modify.html", data)
}
