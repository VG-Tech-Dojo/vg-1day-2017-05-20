package model

import "database/sql"

type Message struct {
	ID    int64  `json:"id"`
	Value string `json:"value"`
}

func MessagesAll(db *sql.DB) ([]*Message, error) {
	rows, err := db.Query(`select * from message`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms []*Message
	for rows.Next() {
		m := &Message{}
		if err := rows.Scan(&m.ID, &m.Value); err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}
