package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &struct {
			Message string `json:"message"`
		}{
			Message: "pong",
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
