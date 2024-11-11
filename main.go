package main

import (
	"log"

	"github.com/fauzancodes/yugioh-open-api/app/config"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	app := Init()

	port := config.LoadConfig().Port

	log.Printf("Server: " + config.LoadConfig().BaseUrl + ":" + port)
	app.Logger.Fatal(app.Start(":" + port))
}

func Init() *echo.Echo {
	app := echo.New()

	config.Database()

	return app
}
