package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id         string       `gorm:"primary_key; not null"`
	CreatedAt  time.Time    `gorm:"type: TIMESTAMP with time zone: not null"`
	ModifiedAt sql.NullTime `gorm:"type: TIMESTAMP with time zone: null"`
	DeletedAt  sql.NullTime `gorm:"type: TIMESTAMP with time zone: null"`

	CreatedBy  *sql.NullInt64 `gorm:"not null"` //it gets ID
	ModifiedBy *sql.NullInt64 `gorm:"null"`     //...
	DeletedBy  *sql.NullInt64 `gorm:"null"`     //...
}

// creating hook
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("userID")
	var userID = &sql.NullInt64{Valid: false}
	if value != nil {
		userID = &sql.NullInt64{Valid: true, Int64: value.(int64)}
	}
	b.CreatedBy = userID
	b.CreatedAt = time.Now().UTC()
	return
}
