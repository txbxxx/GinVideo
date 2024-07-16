/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/12 23:06
 * @File:  VideoListSvc
 * @Software: GoLand
 **/

package videoSvc

import (
	"GliGliVideo/model"
	"GliGliVideo/serializes"
	"GliGliVideo/utils"
	"github.com/gin-gonic/gin"
)

type VideoListSvc struct {
	Page int `json:"page" form:"page" binding:"required,min=1"`
	Size int `json:"size" form:"size" binding:"required,min=1"`
}

// ListVideo 获取视频列表
func (service *VideoListSvc) ListVideo() gin.H {
	videos := make([]model.Video, 0)
	//offset从0开始，所以，需要处理一下page
	service.Page = (service.Page - 1) * service.Size
	//查询列表
	err := utils.DB.Offset(service.Page).Limit(service.Size).Find(&videos).Error
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "获取视频列表失败",
		}
	}
	return gin.H{
		"code": 200,
		"page": service.Page,
		"size": service.Size,
		"msg":  "获取视频列表成功",
		"data": serializes.VideoSerializeList(videos),
	}
}
