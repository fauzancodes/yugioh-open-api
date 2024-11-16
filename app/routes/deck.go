package routes

import (
	"github.com/fauzancodes/yugioh-open-api/app/controllers"
	"github.com/fauzancodes/yugioh-open-api/app/middlewares"
	"github.com/labstack/echo/v4"
)

func DeckRoute(api *echo.Group) {
	deck := api.Group("/deck")
	{
		deck.POST("", controllers.CreateDeck, middlewares.CheckAPIKey, middlewares.Auth)
		deck.GET("/public", controllers.GetPublicDecks)
		deck.GET("", controllers.GetDecks, middlewares.CheckAPIKey, middlewares.Auth)
		deck.GET("/:id", controllers.GetDeckByID, middlewares.CheckAPIKey, middlewares.Auth)
		deck.PATCH("/:id", controllers.UpdateDeck, middlewares.CheckAPIKey, middlewares.Auth)
		deck.DELETE("/:id", controllers.DeleteDeck, middlewares.CheckAPIKey, middlewares.Auth)
	}
}
