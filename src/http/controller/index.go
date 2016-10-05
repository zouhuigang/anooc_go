// Copyright 2016 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author: polaris	polaris@studygolang.com

package controller

import (
	"github.com/labstack/echo"
)

type IndexController struct{}

// 注册路由
func (self IndexController) RegisterRoute(g *echo.Group) {
	g.GET("/", self.Index)
}

// Index 首页
func (IndexController) Index(ctx echo.Context) error {

	return render(ctx, "welcome/index.html", map[string]interface{}{})
}
