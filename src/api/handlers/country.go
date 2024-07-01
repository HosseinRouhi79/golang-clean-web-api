package handlers

import (
	"strconv"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/dto"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/services"
	"github.com/gin-gonic/gin"
)

type Country struct {
	Name string
}

type CountryDelete struct {
	Id string
}

func (co Country) Create(c *gin.Context) {

	cfg := config.GetConfig()
	cs := services.NewCountryService(cfg)

	c.ShouldBind(&co)
	dto := &dto.CreateUpdateCountry{
		Name: co.Name,
	}

	cs.Create(c, dto)
}

func (cd CountryDelete) Delete(c *gin.Context) {

	cfg := config.GetConfig()
	cs := services.NewCountryService(cfg)

	err := c.ShouldBind(&cd)
	if err != nil{
		cs.Logger.Infof("err: %v", err)
	}
	id, _ := strconv.Atoi(cd.Id)
	cs.Delete(c, id)
}
