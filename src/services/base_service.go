package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/helper"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/db"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/pkg/logging"
	"gorm.io/gorm"
)

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	DB     *gorm.DB
	Logger logging.Logger
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		DB:     db.GetDB(),
		Logger: logging.NewLogger(cfg),
	}
}

func (s BaseService[T, Tc, Tu, Tr]) Create(c context.Context, req Tc) (res *Tr, err error) {
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

func (s BaseService[T, Tc, Tu, Tr]) Update(c context.Context, req Tu, id int) (res *Tr, err error) {

	updateMap, _ := helper.TypeConverter[map[string]interface{}](req)
	(*updateMap)["modified_by"] = &sql.NullInt64{Int64: int64(c.Value("id").(float64)), Valid: true}
	(*updateMap)["modified_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}
	model := new(T)
	err = s.DB.Model(model).
		Where("id = ?", id).
		Updates(*updateMap).
		Error

	if err != nil {
		s.Logger.Infof("Error updating: %s", err.Error())
		return nil, err
	}

	type response struct {
		status int
		msg    string
	}
	res2 := response{
		status: 200,
		msg:    "updated successfully",
	}

	response2, err := helper.TypeConverter[Tr](res2)
	if err != nil {
		s.Logger.Infof("Error converting: %s", err.Error())
		return nil, err
	}
	return response2, nil
}
