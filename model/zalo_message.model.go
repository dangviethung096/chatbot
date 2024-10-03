package model

import "time"

type ZaloMessage struct {
	ID          int64     `db:"id"`
	MessageID   string    `db:"message_id"`
	SenderID    string    `db:"sender_id"`
	RecipientID string    `db:"recipient_id"`
	Message     string    `db:"message"`
	ThreadID    string    `db:"thread_id"`
	Index       int64     `db:"index"`
	CreatedAt   time.Time `db:"created_at"`
}

func (message *ZaloMessage) GetTableName() string {
	return "zalo_messages"
}

func (message *ZaloMessage) GetPrimaryKey() string {
	return "id"
}
