package migrations

import (
	"reflect"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/db"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/models"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
)

func Up_1() {
	var cfg = config.GetConfig()
	var zap = logging.NewLogger(cfg)
	db := db.GetDB()
	tables := []interface{}{}
	var countries = models.Country{}
	var cities = models.City{}

	if !db.Migrator().HasTable(countries) {
		tables = append(tables, countries)
	}

	if !db.Migrator().HasTable(cities) {
		tables = append(tables, cities)
	}
	zap.Infof("%v", reflect.TypeOf(tables[0]), reflect.TypeOf(tables[1]))
	zap.Infof("%v", db.Name())
	err := db.Migrator().CreateTable(tables...)
	if err != nil {
		zap.Infof("%v", "Error creating table", err)
	}
}

func Down_1() {

}
