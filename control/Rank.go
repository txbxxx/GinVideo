/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/24 10:38
 * @File:  Rank
 * @Software: GoLand
 **/

package control

import (
	videoSvc "GliGliVideo/service/videoSvc"
	"github.com/gin-gonic/gin"
)

// DailyRank 每日排行
func DailyRank(c *gin.Context) {
	service := videoSvc.DailyRankService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
