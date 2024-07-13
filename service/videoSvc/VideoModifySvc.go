/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/12 23:37
 * @File:  VideoModifySvc
 * @Software: GoLand
 **/

package videoSvc

import (
	"GliGliVideo/model"
	"GliGliVideo/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VideoModifySvc struct {
	Title string `json:"title" form:"title" binding:"required,min=5,max=50"`
	Info  string `json:"info" form:"info"`
}

func (service *VideoModifySvc) ModifyVideo(identity string) gin.H {
	//查询数据是否存在
	var video model.Video
	err := utils.DB.Take(&video, "identity = ?", identity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gin.H{
				"code": -1,
				"msg":  "视频不存在",
			}
		}
		return gin.H{
			"code": -1,
			"msg":  "数据库错误",
		}
	}

	//如果存在则修改
	modifyV := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}

	err = utils.DB.Where("identity = ?", identity).Updates(&modifyV).Error
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "数据库错误",
		}
	}
	return gin.H{
		"code": 200,
		"msg":  "修改成功",
	}
}
