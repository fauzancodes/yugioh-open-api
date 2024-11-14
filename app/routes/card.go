package routes

import (
	"github.com/fauzancodes/yugioh-open-api/app/controllers"
	"github.com/fauzancodes/yugioh-open-api/app/middlewares"
	"github.com/labstack/echo/v4"
)

func CardRoute(api *echo.Group) {
	card := api.Group("/card", middlewares.CheckAPIKey)
	{
		card.POST("", controllers.CreateCard, middlewares.Auth)
		card.GET("", controllers.GetCards)
		card.GET("/:id", controllers.GetCardByID)
		card.PATCH("/:id", controllers.UpdateCard, middlewares.Auth)
		card.DELETE("/:id", controllers.DeleteCard, middlewares.Auth)
		card.GET("/utility", controllers.GetCardUtility)
	}
}
