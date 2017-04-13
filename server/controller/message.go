package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/saxsir/vg-1day-2017/server/model"
)

// Message is controller for requests to messages
type Message struct{}

func (m *Message) Root(c echo.Context) error {
	messages, err := model.MessagesAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, messages)
}
