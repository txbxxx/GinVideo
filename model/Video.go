/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/12 16:04
 * @File:  Video
 * @Software: GoLand
 **/

package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Identity string `gorm:"colum:identity;type:varchar(36);" json:"identity"`
	Title    string `gorm:"colum:title;type:varchar(50);" json:"title"`
	Info     string `gorm:"colum:info;type:text;" json:"info"`
	Url      string `gorm:"colum:url;type:varchar(50);" json:"url"`
}
