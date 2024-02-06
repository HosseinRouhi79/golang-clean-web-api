package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health struct{}

func NewHealth() *Health {
	return &Health{}
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
