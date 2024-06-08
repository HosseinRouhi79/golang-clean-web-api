package routers

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealth()
	r.GET("/", handler.Health)
	r.POST("/", handler.HealthPost)
	r.GET("/:id", handler.HealthPostByID)
}

func Test(r *gin.RouterGroup) {
	handler := handlers.NewTest()
	r.GET("/", handler.HeaderBind)
	r.GET("/query", handler.QueryBind)
	r.GET("/query2/:id/:name", handler.UriBind)
}

func BodyBinder(r *gin.RouterGroup) {
	handler := handlers.PersonData{}
	r.POST("/", handler.BodyBind)
}
func SetToRedisRouter(r *gin.RouterGroup) {
	handler := handlers.Redis{}
	r.POST("/set", handler.SetToRedis)
}

func GetFromRedisRouter(r *gin.RouterGroup) {
	//http://192.168.59.133:5005/api/redis/get/testKey
	handler := handlers.RedisKey{}
	r.GET("/get/:key", handler.GetFromRedis)
}

func GetJWT(r *gin.RouterGroup){
	handler := handlers.JWT{}
	r.GET("/get/jwt", handler.Generate)
}
