/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/12 20:50
 * @File:  VideoSerialize
 * @Software: GoLand
 **/

package serializes

import "GliGliVideo/model"

type VideoSerialize struct {
	Identity  string `json:"identity"`
	Info      string `json:"info"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
}

// VideoSerializeList 视频列表序列化
func VideoSerializeList(videos []model.Video) []VideoSerialize {
	var videoSerializeList []VideoSerialize
	for _, video := range videos {
		videoSerializeList = append(videoSerializeList, VideoSerialize{
			Identity:  video.Identity,
			Info:      video.Info,
			Title:     video.Title,
			Url:       video.Url,
			CreatedAt: video.CreatedAt.Unix(),
		})
	}
	return videoSerializeList
}

func VideoSerializeSingle(video model.Video) VideoSerialize {
	return VideoSerialize{
		Identity:  video.Identity,
		Info:      video.Info,
		Title:     video.Title,
		Url:       video.Url,
		CreatedAt: video.CreatedAt.Unix(),
	}
}
