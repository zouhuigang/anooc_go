package controller

import "github.com/labstack/echo"

func RegisterRoutes(g *echo.Group) {
	new(IndexController).RegisterRoute(g)
	new(AccountController).RegisterRoute(g)
	new(TopicController).RegisterRoute(g)
	new(ArticleController).RegisterRoute(g)
	new(ProjectController).RegisterRoute(g)
	new(ResourceController).RegisterRoute(g)
	new(ReadingController).RegisterRoute(g)
	new(WikiController).RegisterRoute(g)
	new(UserController).RegisterRoute(g)
	new(LikeController).RegisterRoute(g)
	new(FavoriteController).RegisterRoute(g)
	new(MessageController).RegisterRoute(g)
	new(SidebarController).RegisterRoute(g)
	new(CommentController).RegisterRoute(g)
	new(SearchController).RegisterRoute(g)
	new(WideController).RegisterRoute(g)
	new(ImageController).RegisterRoute(g)
	new(CaptchaController).RegisterRoute(g)
	new(WebsocketController).RegisterRoute(g)

	new(InstallController).RegisterRoute(g)
}
