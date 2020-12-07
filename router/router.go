package router

import (
	"template/middleware/jwt"
	"template/pkg/setting"
	"template/router/api"
	v1 "template/router/api/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.Default())
	router.POST("/auth/login", api.GetAuth)
	router.POST("/auth/register", api.CreateUser)
	apiv1 := router.Group("v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/diary", v1.GetDiarys)
		apiv1.PUT("/diary/edit", v1.ModifyDiary)
		apiv1.POST("/diary/count", v1.GetDiaryCount)
		apiv1.POST("/diary/create", v1.CreateDiary)
	}
	return router
}
