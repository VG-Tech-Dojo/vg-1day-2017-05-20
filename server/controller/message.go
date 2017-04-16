package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saxsir/vg-1day-2017/server/model"
)

// Message is controller for requests to messages
type Message struct{}

func (m *Message) Root(c *gin.Context) {
	messages, err := model.MessagesAll()
	if err != nil {
		c.String(500, "%s", err)
		return
	}

	c.JSON(http.StatusOK, messages)
}
