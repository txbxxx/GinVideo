/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/20 11:16
 * @File:  Rank
 * @Software: GoLand
 **/

package Rank

import (
	"GliGliVideo/model"
	"GliGliVideo/utils"
	"context"
	"fmt"
	"strconv"
)

const (
	// DailyRankKey 每日排行
	DailyRankKey = "rank:daily"
)

// AddView 添加点击数
func AddView(video model.Video) {
	// 增加总点击数
	utils.Cache.Incr(context.Background(), VideoViewKey(video.ID))
	// 增加排行点击数
	utils.Cache.ZIncrBy(context.Background(), DailyRankKey, 1, strconv.Itoa(int(video.ID)))
}

// VideoViewKey 视频点击数的key
// view:video:1 -> 100
// view:video:2 -> 150
func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}

// View 获取点击数
func View(video model.Video) uint64 {
	countStr, _ := utils.Cache.Get(context.Background(), VideoViewKey(video.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// RestartDailyRank 重启一天的排名
func RestartDailyRank() error {
	return utils.Cache.Del(context.Background(), "rank:daily").Err()
}
