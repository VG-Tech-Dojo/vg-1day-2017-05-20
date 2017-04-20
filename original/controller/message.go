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

func (m *Message) All(c *gin.Context) {
	msgs, err := model.MessagesAll(m.DB)
	if err != nil {
		c.String(500, "%s", err)
		return
	}

	c.JSON(http.StatusOK, msgs)
}

func (m *Message) GetByID(c *gin.Context) {
	msg, err := model.MessageByID(m.DB, c.Param("id"))

	switch {
	case err == sql.ErrNoRows:
		c.String(500, "%s", err)
	case err != nil:
		c.String(500, "%s", err)
	}

	c.JSON(http.StatusOK, msg)
}

func (m *Message) Create(c *gin.Context) {
	msg := &model.Message{
		Body: "hoge",
	}
	if err := msg.Insert(m.DB); err != nil {
		c.String(500, "%s", err)
	}
	c.String(http.StatusOK, "created")
}
