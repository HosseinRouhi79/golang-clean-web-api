package handlers

import (
	"fmt"
	"net/http"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/helper"
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/validation"
	"github.com/gin-gonic/gin"
)

type Health struct{}
type Test struct {
	UserID  string `json:"userID"`
	Browser string `json:"browser"`
}
type PersonData struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Mobile    string `json:"mobile" binding:"mobile"`
	Password  string `json:"password" binding:"password"`
}

func NewHealth() *Health {
	return &Health{}
}

func NewTest() *Test {
	return &Test{}
}

func (h Health) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "health GET")
}

func (h Health) HealthPost(c *gin.Context) {
	c.JSON(http.StatusOK, "health POST")
}

func (h Health) HealthPostByID(c *gin.Context) {
	// c.JSON(http.StatusOK, "health GET ID")
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("health GET: %v", id))
}

func (t Test) HeaderBind(c *gin.Context) {
	if err := c.BindHeader(&t); err != nil {
		panic("Error binding")
	}
	c.JSON(http.StatusOK, gin.H{
		"Browser": t.Browser,
		"ID":      t.UserID,
	})
}

func (q Test) QueryBind(c *gin.Context) {
	id := c.Query("id") //using QueryArray instead of Query in case we have many ids
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"ID":   id,
		"name": name,
	})
}

func (u Test) UriBind(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"ID":   id,
		"name": name,
	})
}

func (p PersonData) BodyBind(c *gin.Context) {

	var veArr []validation.ValidationError
	err := c.ShouldBind(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,helper.ResponseWithValidationError("", "400", err, veArr))
		return
	}
	c.JSON(200,helper.Response(p, "200"))
	
}
