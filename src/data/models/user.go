package models

type User struct {
	BaseModel
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
}

