package router

import (
	"log"
	"template/middleware/jwt"
	"template/router/api"
	v1 "template/router/api/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.Default())
	log.Println("router run ...")
	router.POST("/auth/login", api.GetAuth)
	router.POST("/auth/register", api.CreateUser)
	apiv1 := router.Group("v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/diary", v1.GetDiarys)
		apiv1.POST("/diary/count", v1.GetDiaryCount)
		apiv1.POST("/diary/create", v1.CreateDiary)
	}
	return router
}
