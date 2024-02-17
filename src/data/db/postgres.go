package db

import (
	"fmt"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	logger "github.com/HosseinRouhi79/golang-clean-web-api/src/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type single struct{}

var singleton *single

var zLog = logger.Logger()
var dbClient *gorm.DB

func InitDB(cfg *config.Config) error {
	if singleton == nil {
		var err error
		cnn := fmt.Sprintf("host=%s port=%s password=%s user=%s sslmode=%s dbname=%s timezone=Asia/Tehran",
			cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Password, cfg.Postgres.User, cfg.Postgres.SSLMode, cfg.Postgres.DbName)
		//we need to open a connection
		dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
		if err != nil {
			zLog.Info().Msg("error openning")
			return err
		}
		sqlDB, err := dbClient.DB()
		if err != nil {
			return err
		}
		err = sqlDB.Ping()
		if err != nil {
			zLog.Info().Msg("error connecting to database")
			return err
		}
		sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime)
		zLog.Info().Msg("creating postgres connection...")
		singleton = &single{}
	} else{
		zLog.Info().Msg("postgres is already created")
	}

	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() error {
	cnn, _ := dbClient.DB()
	cnn.Close()
	return nil
}
