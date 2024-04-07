package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id         int           `gorm:"primarykey; type:int"`
	CreatedAt  time.Time     `gorm:"type: TIMESTAMP with time zone; not null"`
	ModifiedAt *sql.NullTime `gorm:"type: TIMESTAMP with time zone; null"`
	DeletedAt  *sql.NullTime `gorm:"type: TIMESTAMP with time zone; null"`

	CreatedBy  int `gorm:"not null"` //it gets ID
	ModifiedBy *sql.NullInt64 `gorm:"null"` //...
	DeletedBy  *sql.NullInt64 `gorm:"null"` //...
}

// creating hook
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = -1
	// TODO: check userId type
	if value != nil {
		userId = int(value.(float64))
	}
	b.CreatedAt = time.Now().UTC()
	b.CreatedBy = userId
	return
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("userID")
	var userID = &sql.NullInt64{Valid: false}
	if value != nil {
		userID = &sql.NullInt64{Valid: true, Int64: value.(int64)}
	}
	b.ModifiedBy = userID
	b.ModifiedAt = &sql.NullTime{Valid: true, Time: time.Now().UTC()}
	return
}

func (b *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("userID")
	var userID = &sql.NullInt64{Valid: false}
	if value != nil {
		userID = &sql.NullInt64{Valid: true, Int64: value.(int64)}
	}
	b.DeletedBy = userID
	b.DeletedAt = &sql.NullTime{Valid: true, Time: time.Now().UTC()}
	return
}
