package controller

import "github.com/labstack/echo"

func RegisterRoutes(g *echo.Group) {
	new(IndexController).RegisterRoute(g)
	new(BlogController).RegisterRoute(g)
	new(MdController).RegisterRoute(g)
	new(ReactController).RegisterRoute(g)
	new(UserController).RegisterRoute(g)
}
