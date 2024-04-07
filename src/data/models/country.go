package models

type Country struct {
	BaseModel
	Name string `gorm:"not null"`
	Cities *[]City `gorm:"not null"`
}
