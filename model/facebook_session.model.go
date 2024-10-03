package model

import "time"

type FacebookSession struct {
	ID        int64     `db:"id"`
	Sender    string    `db:"sender"`
	ThreadID  string    `db:"thread_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (fs *FacebookSession) GetTableName() string {
	return "facebook_sessions"
}

func (fs *FacebookSession) GetPrimaryKey() string {
	return "id"
}
