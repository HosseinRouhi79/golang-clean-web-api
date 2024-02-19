package models

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	Id         string       `gorm:"primary_key; not null"`
	CreatedAt  time.Time    `gorm:"type: TIMESTAMP with time zone: not null"`
	ModifiedAt sql.NullTime `gorm:"type: TIMESTAMP with time zone: null"`
	DeleteAt   sql.NullTime `gorm:"type: TIMESTAMP with time zone: null"`

	CreatedBy  int            `gorm:"not null"` //it gets ID
	ModifiedBy *sql.NullInt64 `gorm:"null"`     //...
	DeletedBy  *sql.NullInt64 `gorm:"null"`     //...
}
