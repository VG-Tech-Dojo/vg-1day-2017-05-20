package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/saxsir/vg-1day-2017/server/controller"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler

	// Start server
	message := &controller.Message{}
	e.GET("/", message.Root)

	e.Logger.Fatal(e.Start(":1323"))
}
