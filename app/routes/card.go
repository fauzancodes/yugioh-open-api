package routes

import (
	"github.com/fauzancodes/yugioh-open-api/app/controllers"
	"github.com/fauzancodes/yugioh-open-api/app/middlewares"
	"github.com/labstack/echo/v4"
)

func CardRoute(api *echo.Group) {
	card := api.Group("/card")
	{
		card.POST("", controllers.CreateCard, middlewares.CheckAPIKey, middlewares.Auth)
		card.POST("/upload-picture", controllers.UploadCardPicture, middlewares.CheckAPIKey, middlewares.Auth)
		card.GET("", controllers.GetCards)
		card.GET("/:id", controllers.GetCardByID)
		card.PATCH("/:id", controllers.UpdateCard, middlewares.CheckAPIKey, middlewares.Auth)
		card.DELETE("/:id", controllers.DeleteCard, middlewares.CheckAPIKey, middlewares.Auth)
		card.GET("/utility", controllers.GetCardUtility)
	}
}
