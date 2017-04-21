package model

import "database/sql"

type Message struct {
	ID   int64  `json:"id"`
	Body string `json:"body"`
}

// MessagesAll は全てのメッセージを返します
func MessagesAll(db *sql.DB) ([]*Message, error) {
	rows, err := db.Query(`select id, body from message`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms []*Message
	for rows.Next() {
		m := &Message{}
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

	if err := db.QueryRow(`select id, body from message where id = ?`, id).Scan(&m.ID, &m.Body); err != nil {
		return nil, err
	}

	return m, nil
}

// Insertはmessageテーブルに新規データを1件追加します
func (m *Message) Insert(db *sql.DB) (sql.Result, error) {
	return db.Exec(`insert into message (body) values (?)`, m.Body)
}
