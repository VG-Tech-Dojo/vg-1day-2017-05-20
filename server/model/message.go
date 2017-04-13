package model

type Message struct {
	ID    int64  `json:"id"`
	Value string `json:"value"`
}

func MessagesAll() ([]*Message, error) {
	return []*Message{
		&Message{
			ID:    1,
			Value: "hoge",
		},
		&Message{
			ID:    2,
			Value: "fuga",
		},
		&Message{
			ID:    2,
			Value: "piyo",
		},
	}, nil
}
