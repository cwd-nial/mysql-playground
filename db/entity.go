package db

import (
	"database/sql"
)

type Entity struct {
	Id          uint64       `db:"id"`
	MessageType string       `db:"message_type"`
	SpaceType   string       `db:"space_type"`
	ReceiverId  uint64       `db:"receiver_id"`
	LogData     string       `db:"log_data"`
	CreateTime  sql.NullTime `db:"create_time"`
}
