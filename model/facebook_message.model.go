package model

import "time"

type FacebookMessage struct {
	ID        int64     `db:"id"`
	Sender    string    `db:"sender"`
	Receiver  string    `db:"receiver"`
	Message   string    `db:"message"`
	ThreadID  string    `db:"thread_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (fm *FacebookMessage) GetTableName() string {
	return "facebook_messages"
}

func (fm *FacebookMessage) GetPrimaryKey() string {
	return "id"
}
