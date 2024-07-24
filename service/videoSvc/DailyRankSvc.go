/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/24 10:33
 * @File:  DailyRankSvc
 * @Software: GoLand
 **/

package videoSvc

import (
	"GliGliVideo/model"
	"GliGliVideo/serializes"
	"GliGliVideo/utils"
	"GliGliVideo/utils/Rank"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// DailyRankService 每日排行的服务
type DailyRankService struct {
}

// Get 获取排行
func (service *DailyRankService) Get() gin.H {
	var videos []model.Video

	// 从redis读取点击前十的视频
	vids, _ := utils.Cache.ZRevRange(context.Background(), Rank.DailyRankKey, 0, 9).Result()
	fmt.Println(len(vids))
	if len(vids) >= 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(vids, ","))
		fmt.Println("ooo ", order)
		err := utils.DB.Where("id in (?)", vids).Order(order).Find(&videos).Error
		if err != nil {
			return gin.H{
				"code": -1,
				"msg":  "获取排行榜失败",
			}
		}
	}

	return gin.H{
		"code": 200,
		"msg":  "获取排行榜成功",
		"data": serializes.VideoSerializeList(videos),
	}
}
