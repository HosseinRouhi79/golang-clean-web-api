package models

type UserRole struct {
	UserID int
	RoleID int
	User   User `gorm:"foreignKey:UserID; constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Role   Role `gorm:"foreignKey:RoleID; constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
