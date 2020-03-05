package db

import (
	"database/sql"
	"time"
)

type FileMeta struct{
	FileSha1 string
	FileName string
	FileSize sql.NullInt64
	FileAddr string
	CreateAt time.Time
	UpdateAt time.Time
}
