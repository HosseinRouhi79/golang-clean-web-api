package handlers

import (
	"strconv"
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
	cs := services.NewBaseService[Country, dto.CreateUpdateCountry, dto.CreateUpdateCountry, dto.CountryResponse](cfg)

	c.ShouldBind(&co)
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

func (cd CountryDelete) Delete(c *gin.Context) {

	cfg := config.GetConfig()
	cs := services.NewCountryService(cfg)

	err := c.ShouldBind(&cd)
	if err != nil {
		cs.Logger.Infof("err: %v", err)
	}
	id, _ := strconv.Atoi(cd.Id)
	cs.Delete(c, id)
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
