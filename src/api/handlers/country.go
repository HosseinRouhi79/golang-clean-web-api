package handlers

import (
	"fmt"
	"time"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/helper"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/data/models"
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
type CitiesAssined struct {
	Cities    []string `form:"cities"`
	CountryID int      `form:"countryid"`
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

// Country godoc
// @Summary Country
// @Description Get country
// @Tags country
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param id path string true "country id"
// @Success 200 {object} helper.HTTPResponse "Success"
// @Failure 400 {object} helper.HTTPResponse "Failed"
// @Router /c/get/{id} [get]
// @Security AuthBearer
func GetByID(c *gin.Context) {
	cfg := config.GetConfig()
	cs := services.NewCountryService(cfg)
	id := c.Param("id")

	res, err := cs.GetByID(c, id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, res)

}

// Country godoc
// @Summary Country
// @Description Get country
// @Tags country
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} helper.HTTPResponse "Success"
// @Failure 400 {object} helper.HTTPResponse "Failed"
// @Router /c/get/all [get]
// @Security AuthBearer
func GetAllCountries(c *gin.Context) {
	cfg := config.GetConfig()
	cs := services.NewCountryService(cfg)

	res, err := cs.GetAllCountries()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, res)

}

// AssignCityToCountry godoc
// @Summary Assign cities to a country
// @Description Assign a list of cities to a specified country by country ID
// @Tags country
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   cities     formData  []string  true  "Array of city names"
// @Param   countryid  formData  int       true  "Country ID"
// @Success 200 {object} helper.HTTPResponse "Success"
// @Failure 400 {object} helper.HTTPResponse "Failed"
// @Router /c/assign/cities [post]
func (cities CitiesAssined) AssignCityToCountry(c *gin.Context) {
	fmt.Println(c.Request)
	modelCityList := []models.City{}
	cfg := config.GetConfig()

	err := c.ShouldBind(&cities)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	cs := services.NewCountryService(cfg)
	for index := range cities.Cities{
		cMap := map[string]string{
			"name": cities.Cities[index],
		}
		res, err := helper.TypeConverter[models.City](cMap)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		modelCityList = append(modelCityList, *res)
	}

	res, err := cs.AssignCity(c, modelCityList, cities.CountryID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, res)

}
