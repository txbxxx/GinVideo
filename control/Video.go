/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/12 16:17
 * @File:  Video
 * @Software: GoLand
 **/

package control

import (
	"GliGliVideo/service/videoSvc"
	"github.com/gin-gonic/gin"
)

// CreateVideo 创建视频
func CreateVideo(c *gin.Context) {
	var service videoSvc.VideoCreateSvc
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, service.CreateVideo())
}

// ListVideo 列出视频
func ListVideo(c *gin.Context) {
	var service videoSvc.VideoListSvc
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, service.ListVideo())
}

func ModifyVideo(c *gin.Context) {
	var service videoSvc.VideoModifySvc
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, service.ModifyVideo(c.Param("identity")))
}
