package models

type UserRole struct {
	BaseModel
	User   User `gorm:"foreignKey:UserId"`
	Role   Role `gorm:"foreignKey:RoleId"`
	UserId int
	RoleId int
}