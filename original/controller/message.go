package controller

import (
	"database/sql"
	"net/http"

	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
	"github.com/gin-gonic/gin"
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
