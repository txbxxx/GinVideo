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
	"github.com/sirupsen/logrus"
)

// CreateVideo 创建视频
func CreateVideo(c *gin.Context) {
	var service videoSvc.VideoCreateSvc
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, service.CreateVideo())
}

// ListVideo 列出视频
func ListVideo(c *gin.Context) {
	var service videoSvc.VideoListSvc
	if err := c.ShouldBind(&service); err != nil {
		logrus.Println("列出视频错误", err.Error())
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, service.ListVideo())
}

// ModifyVideo 修改视频
func ModifyVideo(c *gin.Context) {
	var service videoSvc.VideoModifySvc
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, service.ModifyVideo(c.Param("identity")))
}

// DetailVideo 视频详情
func DetailVideo(c *gin.Context) {
	var service videoSvc.VideoDetailSvc
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, service.DetailVideo(c.Param("identity")))
}

// DeleteVideo 删除视频
func DeleteVideo(c *gin.Context) {
	var service videoSvc.VideoDeleteSvc
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, service.DeleteVideo(c.Param("identity")))
}

// UploadVideo 上传视频
func UploadVideo(c *gin.Context) {
	var service videoSvc.VideoUploadSvc
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, service.Post())
}
