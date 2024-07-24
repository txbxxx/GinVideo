/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 16:31
 * @File:  Cache
 * @Software: GoLand
 **/

package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var Cache *redis.Client

// RedisUtils redis连接
func RedisUtils(RDBAddr, RDBPwd, RDBDefaultDB string) {
	// 将字符串转换成int
	RDB, err := strconv.Atoi(RDBDefaultDB)
	if err != nil {
		fmt.Println("将string转换成int失败！")
	}

	//连接redis
	client := redis.NewClient(&redis.Options{
		Addr:     RDBAddr,
		Password: RDBPwd,
		DB:       RDB,
	})
	ping := client.Ping(context.Background())
	fmt.Println("test: ", ping)
	if ping.Err() != nil {
		fmt.Println("redis连接失败！")
	}
	Cache = client
}
