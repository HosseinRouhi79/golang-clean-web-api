package migrations

import (
	"reflect"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/db"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/models"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Up_1() {
	var cfg = config.GetConfig()
	var zap = logging.NewLogger(cfg)
	db := db.GetDB()
	tables := []interface{}{}
	var countries = models.Country{}
	var cities = models.City{}
	var users = models.User{}
	var roles = models.Role{}
	var userRoles = models.UserRole{}

	tables = addNewTable(db, countries, tables)
	tables = addNewTable(db, cities, tables)
	tables = addNewTable(db, users, tables)
	tables = addNewTable(db, roles, tables)
	tables = addNewTable(db, userRoles, tables)

	err := db.Migrator().CreateTable(tables...)
	if err != nil {
		zap.Infof("%v", "Error creating table", err)
	}
	{
		ci := []models.City{{Name: "Test City", CountryID: 1}}
		c := models.Country{Name: "Test Country"}

		c.Id = 1
		c.Cities = &ci

		db.Create(&c)
	}
	{
		pass := "12345678"
		hashedPass, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		roles := []*models.Role{
			{Name: "Admin"},
			{Name: "User"},
		}
		for _, role := range roles {
			db.Create(role)
		}
		// Create sample users
		users := []*models.User{
			{
				Username:  "user1",
				FirstName: "test1",
				LastName:  "test1",
				Email:     "test1@example.com",
				Password:  string(hashedPass),
				Enabled:   true,
				UserRoles: &[]models.UserRole{
					{RoleId: roles[0].Id}, // Assigning Admin role
					{RoleId: roles[1].Id}, // Assigning User role
				},
			},
			{
				Username:  "user2",
				FirstName: "test2",
				LastName:  "test2",
				Email:     "test2@example.com",
				Password:  string(hashedPass),
				Enabled:   true,
				UserRoles: &[]models.UserRole{
					{RoleId: roles[1].Id}, // Assigning User role
				},
			},
		}
		for _, user := range users {
			db.Create(user)
		}
	}
	zap.Infof("%v", reflect.TypeOf(tables[0]), reflect.TypeOf(tables[1]))
	zap.Infof("%v", db.Name())

}

func Down_1() {

}

func addNewTable(db *gorm.DB, table interface{}, tables []interface{}) []interface{} {
	if !db.Migrator().HasTable(table) {
		tables = append(tables, table)
	}
	return tables
}
