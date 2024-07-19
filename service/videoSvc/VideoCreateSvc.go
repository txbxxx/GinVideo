/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/12 16:20
 * @File:  VideoCreateSvc.go
 * @Software: GoLand
 **/

package videoSvc

import (
	"GliGliVideo/model"
	"GliGliVideo/serializes"
	"GliGliVideo/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type VideoCreateSvc struct {
	Title string `json:"title" form:"title" binding:"required,max=100"`
	Info  string `json:"info" form:"info"`
	Url   string `json:"url" form:"url" binding:"required"`
	Cover string `json:"cover" form:"cover" binding:"required"`
}

func (service *VideoCreateSvc) CreateVideo() gin.H {
	//创建视频对象
	video := model.Video{
		Identity: utils.GenerateUUID(),
		Title:    service.Title,
		Info:     service.Info,
		Url:      service.Url,
		Cover:    service.Cover,
	}

	//插入数据库
	err := utils.DB.Create(&video).Error
	if err != nil {
		logrus.Error("创建视频失败" + err.Error())
		return gin.H{
			"code": -1,
			"msg":  "创建失败",
		}
	}

	return gin.H{
		"code": 200,
		"msg":  "视频创建成功！",
		"data": serializes.VideoSerializeSingle(video),
	}
}
