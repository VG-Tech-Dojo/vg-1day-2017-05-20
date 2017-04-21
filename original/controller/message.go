package controller

import (
	"database/sql"
	"errors"
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
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, msgs)
}

func (m *Message) GetByID(c *gin.Context) {
	msg, err := model.MessageByID(m.DB, c.Param("id"))

	switch {
	case err == sql.ErrNoRows:
		errorResponse(c, http.StatusNotFound, err)
	case err != nil:
		errorResponse(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, msg)
}

func (m *Message) Create(c *gin.Context) {
	var msg model.Message
	if err := c.BindJSON(&msg); err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if msg.Body == "" {
		errorResponse(c, http.StatusBadRequest, errors.New("body is missing"))
		return
	}

	inserted, err := msg.Insert(m.DB)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"result": gin.H{
			"message": inserted,
		},
	})
}

// errorResponse return json response for api error
func errorResponse(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"result": nil,
		"error": gin.H{
			"message": err.Error(),
		},
	})
}
