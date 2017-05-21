package model

import (
	"database/sql"
)

// Message はメッセージの構造体です
type Message struct {
	ID   int64  `json:"id"`
	Body string `json:"body"`
	// 1-1. ユーザー名を表示しよう
}

// MessagesAll は全てのメッセージを返します
func MessagesAll(db *sql.DB) ([]*Message, error) {

	// 1-1. ユーザー名を表示しよう
	rows, err := db.Query(`select id, body from message`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms []*Message
	for rows.Next() {
		m := &Message{}
		// 1-1. ユーザー名を表示しよう
		if err := rows.Scan(&m.ID, &m.Body); err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

// MessageByID は指定されたIDのメッセージを1つ返します
func MessageByID(db *sql.DB, id string) (*Message, error) {
	m := &Message{}

	// 1-1. ユーザー名を表示しよう
	if err := db.QueryRow(`select id, body from message where id = ?`, id).Scan(&m.ID, &m.Body); err != nil {
		return nil, err
	}

	return m, nil
}

// Insert はmessageテーブルに新規データを1件追加します
func (m *Message) Insert(db *sql.DB) (*Message, error) {
	// 1-2. ユーザー名を追加しよう
	res, err := db.Exec(`insert into message (body) values (?)`, m.Body)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Message{
		ID:   id,
		Body: m.Body,
		// 1-2. ユーザー名を追加しよう
	}, nil
}

// 1-3. メッセージを編集しよう
// ...

// 1-4. メッセージを削除しよう
// ...
