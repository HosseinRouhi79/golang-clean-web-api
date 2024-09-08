package handlers

import (
	"strconv"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/dto"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/services"
	"github.com/gin-gonic/gin"
)

type Country struct {
	Name string `form:"name"`
}

type CountryDelete struct {
	Id string
}

type CountryID struct {
	Id string
}


// Country godoc
// @Summary Country
// @Description Create a new country
// @Tags country
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "country name"
// @Success 200 {object} helper.HTTPResponse "Success"
// @Failure 400 {object} helper.HTTPResponse "Failed"
// @Router /c/create [post]
// @Security AuthBearer
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

func (ci CountryID) GetByID(c *gin.Context){
	cfg := config.GetConfig()
	cs := services.NewCountryService(cfg)

	err := c.ShouldBind(&ci)
	if err != nil{
		cs.Logger.Infof("err get country id :%s", err.Error())
	}
	cs.GetByID(c, ci.Id)
}
