package handlers

import (
	"fmt"
	"time"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/dto"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/services"
	"github.com/gin-gonic/gin"
)

type Country struct {
	CreatedBy int       `form:"createdby"`
	CreatedAt time.Time `form:"createdat"`
	Name      string    `form:"name"`
}

type CountryUpdate struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
}

type CountryDelete struct {
	Id int `form:"id"`
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
	cs := services.NewBaseService[Country, dto.CreateUpdateCountry, dto.CreateUpdateCountry, dto.CountryResponse](cfg)

	err := c.ShouldBind(&co)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	dto := &dto.CreateUpdateCountry{
		Name: co.Name,
	}
	res, err := cs.Create(c, *dto)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
		"data":   res,
	})
}

// Country godoc
// @Summary Country
// @Description Update country
// @Tags country
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "country name"
// @Param id formData string true "country id"
// @Success 200 {object} helper.HTTPResponse "Success"
// @Failure 400 {object} helper.HTTPResponse "Failed"
// @Router /c/update [put]
// @Security AuthBearer
func (co CountryUpdate) Update(c *gin.Context) {

	cfg := config.GetConfig()
	cs := services.NewBaseService[Country, dto.CreateUpdateCountry, dto.CreateUpdateCountry, dto.CountryResponse](cfg)

	err := c.ShouldBind(&co)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"binding error": err.Error(),
		})
		return
	}
	dto := &dto.CreateUpdateCountry{
		Name: co.Name,
	}
	res, err := cs.Update(c, *dto, co.Id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
		"data":   res,
	})
}

// Country godoc
// @Summary Country
// @Description Delete country
// @Tags country
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param id formData string true "country id"
// @Success 200 {object} helper.HTTPResponse "Success"
// @Failure 400 {object} helper.HTTPResponse "Failed"
// @Router /c/delete [put]
// @Security AuthBearer
func (cd CountryDelete) Delete(c *gin.Context) {

	cfg := config.GetConfig()
	cs := services.NewBaseService[Country, dto.CreateUpdateCountry, dto.CreateUpdateCountry, dto.CountryResponse](cfg)

	err := c.ShouldBind(&cd)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		cs.Logger.Infof("err: %v", err)
	}
	cs.Delete(c, cd.Id)
}

func (ci CountryID) GetByID(c *gin.Context) {
	cfg := config.GetConfig()
	cs := services.NewCountryService(cfg)

	err := c.ShouldBind(&ci)
	if err != nil {
		cs.Logger.Infof("err get country id :%s", err.Error())
	}
	cs.GetByID(c, ci.Id)
}
