package services

import (
	"context"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/db"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/models"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/dto"
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

// create
func (cs *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountry) (res *dto.CountryResponse, err error) {
	country := models.Country{Name: req.Name}
	tx := cs.DB.WithContext(ctx).Begin()

	err = tx.Create(&country).Error

	if err != nil {
		tx.Rollback()
		cs.Logger.Info(logging.Postgres, logging.Insert, err.Error(), nil)
	}
	tx.Commit()
	res = &dto.CountryResponse{
		Name: country.Name,
	}
	return res, nil

}

//update

//delete

//get by ID
