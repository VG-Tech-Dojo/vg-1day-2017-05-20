package controller

import 	(
	"database/sql"
	"errors"
	"net/http"
	"github.com/VG-Tech-Dojo/vg-1day-2017-05-20/takumi-n/httputil"
	"github.com/VG-Tech-Dojo/vg-1day-2017-05-20/takumi-n/model"
	"github.com/gin-gonic/gin"
	"strconv"
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
	// 1-3. メッセージを編集しよう
	// ...
	var msg model.Message
	if err := c.BindJSON(&msg); err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	id, err1 := strconv.ParseInt(c.Param("id"), 10, 64)

	if err1 != nil {
		resp := httputil.NewErrorResponse(err1)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	msg.ID = id

	_, err2 := msg.Update(m.DB)

	if err2 != nil {
		resp := httputil.NewErrorResponse(err2)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// DeleteByID は...
func (m *Message) DeleteByID(c *gin.Context) {
	// 1-4. メッセージを削除しよう
	// ...
	c.JSON(http.StatusOK, gin.H{})
}
