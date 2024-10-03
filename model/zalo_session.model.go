package model

import "time"

type ZaloSession struct {
	ID        int64     `db:"id"`
	SenderID  string    `db:"sender_id"`
	ThreadID  string    `db:"thread_id"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func (session *ZaloSession) GetTableName() string {
	return "zalo_sessions"
}

func (session *ZaloSession) GetPrimaryKey() string {
	return "id"
}
