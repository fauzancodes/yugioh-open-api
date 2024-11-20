package routes

import (
	"github.com/fauzancodes/yugioh-open-api/app/controllers"
	"github.com/fauzancodes/yugioh-open-api/app/middlewares"
	"github.com/labstack/echo/v4"
)

func Route(app *echo.Echo) {
	app.Static("/assets", "assets")
	app.Static("/docs", "docs")

	app.GET("/", controllers.Index, middlewares.StripHTMLMiddleware)
	app.GET("/postman/collection", controllers.DownloadPostmanCollection, middlewares.StripHTMLMiddleware)
	app.GET("/postman/environment", controllers.DownloadPostmanEnvironment, middlewares.StripHTMLMiddleware)
	api := app.Group("/v1", middlewares.StripHTMLMiddleware)
	{
		AuthRoute(api)
		CardRoute(api)
		DeckRoute(api)
	}
}
