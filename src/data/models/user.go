package models

type User struct {
	BaseModel
	UserName  string `gorm:"type:string; not null; unique; size:20"`
	FirstName string `gorm:"not null; size:15; type:string"`
	LastName  string `gorm:"not null; size:25; type:string"`
	Mobile    string `gorm:"null; size:11; default:null; type:string"`
	Password  string `gorm:"not null; size:64; type:string"`
	Email     string `gorm:"not null; size:64; type:string"`
	Enabled   bool   `gorm:"default:true"`
	UserRole  *[]UserRole
}
