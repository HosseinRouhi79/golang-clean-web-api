package services

import (
	"context"
	"fmt"
	"time"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/helper"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/db"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
	"gorm.io/gorm"
)

type BaseService[T any, Tc any, Tr any] struct {
	DB     *gorm.DB
	Logger logging.Logger
}

func NewBaseService[T any, Tc any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tr] {
	return &BaseService[T, Tc, Tr]{
		DB:     db.GetDB(),
		Logger: logging.NewLogger(cfg),
	}
}

func (s BaseService[T, Tc, Tr]) Create(c context.Context, req Tc) (res *Tr, err error) {
	model, err := helper.TypeConverter[map[string]interface{}](req)
	if err != nil {
		return nil, err
	}
	(*model)["createdby"] = int(c.Value("id").(float64))
	(*model)["createdat"] = time.Now().UTC()
	fmt.Println(model)
	tx := s.DB.WithContext(c).Begin()
	model2, err := helper.TypeConverter[T](model)
	if err != nil {
		tx.Rollback()
		s.Logger.Info(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	fmt.Println(model2)
	
	err = tx.Create(&model2).Error

	if err != nil {
		tx.Rollback()
		s.Logger.Info(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	res, err = helper.TypeConverter[Tr](model2)
	if err != nil {
		tx.Rollback()
		s.Logger.Info(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	return res, nil

}
