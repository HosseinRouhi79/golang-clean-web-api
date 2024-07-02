package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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
	country.CreatedBy = int(ctx.Value("id").(float64))
	country.CreatedAt = time.Now().UTC()
	tx := cs.DB.WithContext(ctx).Begin()

	err = tx.Create(&country).Error

	if err != nil {
		tx.Rollback()
		cs.Logger.Info(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	res = &dto.CountryResponse{
		Name: country.Name,
	}
	return res, nil

}

//update

func (cs *CountryService) Update(ctx context.Context, req *dto.CreateUpdateCountry, id int) (res *dto.CountryResponse, err error) {

	//create updateMap
	updateMap := map[string]interface{}{
		"name":        req.Name,
		"modified_by": sql.NullInt64{Int64: int64(ctx.Value("id").(float64)), Valid: true},
		"modified_at": sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	tx := cs.DB.WithContext(ctx).Begin()
	err = tx.
		Model(models.Country{}).
		Where("id = ?", id).
		Updates(updateMap).
		Error

	if err != nil {
		tx.Rollback()
		cs.Logger.Info(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}

	tx.Commit()
	country := models.Country{}
	cs.DB.Model(models.Country{}).Where("id = ?", id).First(&country)

	res = &dto.CountryResponse{
		Name: country.Name,
	}
	return res, nil
}

// delete
func (cs *CountryService) Delete(ctx context.Context, id int) (res *dto.CountryResponse, err error) {

	deletedMap := map[string]interface{}{
		"deleted_by": sql.NullInt64{Int64: int64(ctx.Value("id").(float64)), Valid: true},
		"deleted_at": sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	fmt.Println(deletedMap)
	tx := cs.DB.WithContext(ctx).Begin()
	err = tx.
		Model(&models.Country{}).
		Where("id = ?", id).
		Updates(deletedMap).
		Error

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		cs.Logger.Info(logging.Postgres, logging.Delete, err.Error(), nil)
		return nil, err
	}

	tx.Commit()
	res = &dto.CountryResponse{
		Name: "", // Name is empty because the country is deleted
	}

	return res, nil
}

//get by ID
