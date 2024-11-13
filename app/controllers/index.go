package controllers

import "github.com/labstack/echo/v4"

func Index(c echo.Context) error {
	return c.HTML(200, "<h1>Welcome!</h1>")
}
