/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 15:28
 * @File:  route
 * @Software: GoLand
 **/

package router

import (
	"GliGliVideo/control"
	"GliGliVideo/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	httpServer := gin.Default()
	//跨域
	httpServer.Use(middleware.Cors())

	user := httpServer.Group("/user")
	{
		user.POST("/login", control.Login)
		user.POST("/register", control.Register)
	}

	video := httpServer.Group("/video")
	{
		video.POST("/create", control.CreateVideo)
		video.GET("/list", control.ListVideo)
		video.PUT("/:identity", control.ModifyVideo)
		video.GET("/:identity", control.DetailVideo)
		video.DELETE("/:identity", control.DeleteVideo)
		video.POST("/upload", control.UploadVideo)
	}
	return httpServer
}
