package routes

import (
	"github.com/fauzancodes/yugioh-open-api/app/controllers"
	"github.com/fauzancodes/yugioh-open-api/app/middlewares"
	"github.com/labstack/echo/v4"
)

func Route(app *echo.Echo) {
	app.Static("/assets", "assets")
	app.Static("/docs", "docs")

	app.GET("/", controllers.Index, middlewares.StripHTMLMiddleware, middlewares.CheckAPIKey)
	api := app.Group("/v1", middlewares.StripHTMLMiddleware)
	{
		AuthRoute(api)
		CardRoute(api)
	}
}
