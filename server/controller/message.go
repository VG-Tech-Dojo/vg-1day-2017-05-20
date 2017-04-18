package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saxsir/vg-1day-2017/server/model"
)

// Message is controller for requests to messages
type Message struct {
	DB *sql.DB
}

func (m *Message) Root(c *gin.Context) {
	messages, err := model.MessagesAll(m.DB)
	if err != nil {
		c.String(500, "%s", err)
		return
	}

	c.JSON(http.StatusOK, messages)
}
