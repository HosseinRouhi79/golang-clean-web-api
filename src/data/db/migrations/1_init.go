package migrations

import (
	"reflect"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/constants"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/db"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/models"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
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

	password := "hashedPass"

	role := models.Role{Name: constants.AdminRole, UserRole: &ur}

	tables = addNewTable(db, countries, tables)
	tables = addNewTable(db, cities, tables)
	tables = addNewTable(db, users, tables)
	tables = addNewTable(db, roles, tables)
	tables = addNewTable(db, userRoles, tables)

	err := db.Migrator().CreateTable(tables...)
	if err != nil {
		zap.Infof("%v", "Error creating table", err)
	}

	admin := models.User{FirstName: "admin", LastName: "admin", Email: "admin@gmail.com", Password: password,
		Enabled: true, Username: "admin"}

	ur := []models.UserRole{{User: admin}}
	admin.UserRoles = &ur

	// db.Create(&ur[0])
	db.Create(&admin)
	// db.Create(role)

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
