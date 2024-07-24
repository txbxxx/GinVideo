/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/13 14:22
 * @File:  VideoDetailSvc
 * @Software: GoLand
 **/

package videoSvc

import (
	"GliGliVideo/model"
	"GliGliVideo/serializes"
	"GliGliVideo/utils"
	"GliGliVideo/utils/Rank"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VideoDetailSvc struct {
}

func (service *VideoDetailSvc) DetailVideo(identity string) gin.H {
	//通过identity查找视频是否存在
	var video model.Video
	err := utils.DB.Where("identity = ?", identity).Take(&video).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gin.H{
				"code": -1,
				"msg":  "视频不存在",
			}
		}
		logrus.Println("数据库错误: ", err.Error())
		return gin.H{
			"code": -1,
			"msg":  "数据库错误",
		}
	}
	//处理视频被观看的一系问题
	Rank.AddView(video)

	//存在直接返回
	return gin.H{
		"code": 200,
		"msg":  "获取视频详情成功",
		"data": serializes.VideoSerializeSingle(video),
	}
}
