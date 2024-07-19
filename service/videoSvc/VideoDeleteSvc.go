/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/13 14:42
 * @File:  VideoDeleteSvc
 * @Software: GoLand
 **/

package videoSvc

import (
	"GliGliVideo/model"
	"GliGliVideo/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VideoDeleteSvc struct {
}

func (service *VideoDeleteSvc) DeleteVideo(identity string) gin.H {
	//视频存在执行删除操作
	err := utils.DB.Delete(&model.Video{}, "identity = ?", identity).Error
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

	return gin.H{
		"code": 200,
		"msg":  "删除成功",
	}
}
