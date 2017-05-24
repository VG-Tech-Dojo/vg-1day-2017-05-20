package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/VG-Tech-Dojo/vg-1day-2017/yangwei/httputil"
	"github.com/VG-Tech-Dojo/vg-1day-2017/yangwei/model"
	"github.com/gin-gonic/gin"
)

// Message is controller for requests to messages
type Message struct {
	DB     *sql.DB
	Stream chan *model.Message
}

// All は全てのメッセージを取得してJSONで返します
func (m *Message) All(c *gin.Context) {
	msgs, err := model.MessagesAll(m.DB)
	if err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(msgs) == 0 {
		c.JSON(http.StatusOK, make([]*model.Message, 0))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": msgs,
		"error":  nil,
	})
}

// GetByID はパラメーターで受け取ったidのメッセージを取得してJSONで返します
func (m *Message) GetByID(c *gin.Context) {
	msg, err := model.MessageByID(m.DB, c.Param("id"))

	switch {
	case err == sql.ErrNoRows:
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusNotFound, resp)
		return
	case err != nil:
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": msg,
		"error":  nil,
	})
}

// Create は新しいメッセージ保存し、作成したメッセージをJSONで返します
func (m *Message) Create(c *gin.Context) {
	var msg model.Message
	if err := c.BindJSON(&msg); err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	if msg.Body == "" {
		resp := httputil.NewErrorResponse(errors.New("body is missing"))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// 1-2. ユーザー名を追加しよう
	// ユーザー名が空でも投稿できるようにするかどうかは自分で考えてみよう

	inserted, err := msg.Insert(m.DB)
	if err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// bot対応
	m.Stream <- inserted

	c.JSON(http.StatusCreated, gin.H{
		"result": inserted,
		"error":  nil,
	})
}

// UpdateByID は...
func (m *Message) UpdateByID(c *gin.Context) {
	var msg model.Message
	if err := c.BindJSON(&msg); err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if msg.Body == "" {
		resp := httputil.NewErrorResponse(errors.New("body is missing"))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if id < 1 {
		resp := httputil.NewErrorResponse(errors.New("id should be a positive number"))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	msg.ID = id
	updated, err := msg.Update(m.DB)

	if err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": updated,
		"error":  nil,
	})
}

// DeleteByID は...
func (m *Message) DeleteByID(c *gin.Context) {
  if err := model.DeleteByID(m.DB, c.Param("id")); err != nil {
     resp := httputil.NewErrorResponse(err)
     c.JSON(http.StatusInternalServerError, resp)
     return
   }

   c.JSON(http.StatusOK, gin.H{
     "result": nil,
     "error":  nil,
   })
}
