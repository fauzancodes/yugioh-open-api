package routes

import (
	"github.com/fauzancodes/yugioh-open-api/app/controllers"
	"github.com/fauzancodes/yugioh-open-api/app/middlewares"
	"github.com/labstack/echo/v4"
)

func AuthRoute(app *echo.Group) {
	auth := app.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.GET("/user", controllers.GetCurrentUser, middlewares.CheckAPIKey, middlewares.Auth)
		auth.PATCH("/update-profile", controllers.UpdateProfile, middlewares.CheckAPIKey, middlewares.Auth)
		auth.DELETE("/remove-account", controllers.RemoveAccount, middlewares.CheckAPIKey, middlewares.Auth)
		auth.POST("/generate-api-key", controllers.GenerateApiKey, middlewares.CheckAPIKey, middlewares.Auth)
	}
}
