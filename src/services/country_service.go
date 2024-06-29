package services

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/db"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
	"gorm.io/gorm"
)

type CountryService struct {
	DB     *gorm.DB
	Logger logging.Logger
}

func NewCountryService(cfg *config.Config) *CountryService {

	return &CountryService{
		DB:     db.GetDB(),
		Logger: logging.NewLogger(cfg),
	}
}

//create

//update

//delete

//get by ID
