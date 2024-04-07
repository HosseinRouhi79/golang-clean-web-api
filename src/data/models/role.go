package models

type Role struct {
	BaseModel
	Name     string `gorm:"unique; not null; type:string; size:15"`
	UserRole *[]UserRole
}
