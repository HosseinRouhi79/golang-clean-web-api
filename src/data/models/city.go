package models

type City struct {
	BaseModel
	Name      string  `gorm:"not null;"`
	CountryID int     
	Country   Country `gorm:"foreignKey:CountryID;"`
}
