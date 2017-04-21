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
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, msgs)
}

func (m *Message) GetByID(c *gin.Context) {
	msg, err := model.MessageByID(m.DB, c.Param("id"))

	switch {
	case err == sql.ErrNoRows:
		c.JSON(http.StatusNotFound, model.APIResponse{
			Error: &model.APIError{
				Message: err.Error(),
			},
		})
	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "error",
			"error":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, msg)
}

func (m *Message) Create(c *gin.Context) {
	var msg model.Message
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "error",
			"error":  err.Error(),
		})
	}

	//NOTE: insert結果受け取ってjsonで何か返す?
	_, err := msg.Insert(m.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "error",
			"error":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Result: &model.APIResult{
			Message: &msg,
		},
	})
}
