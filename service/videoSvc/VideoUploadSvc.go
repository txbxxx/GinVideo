/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/17 13:42
 * @File:  VideoUploadSvc
 * @Software: GoLand
 **/

package videoSvc

import (
	"GliGliVideo/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"time"
)

type VideoUploadSvc struct {
	FileName string `json:"file_name" form:"file_name"`
}

func (service *VideoUploadSvc) Post() gin.H {
	token, err := utils.CosToken()
	if err != nil {
		logrus.Error("获取临时密钥失败", err.Error())
		return gin.H{
			"code": -1,
			"msg":  "获取临时密钥失败",
		}
	}
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse("https://govideo-1305907375.cos.ap-hongkong.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{})

	//获取文件后缀
	ext := filepath.Ext(service.FileName)
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	//key := "upload/video/" + service.FileName
	key := "upload/video/" + utils.GenerateUUID() + ext
	opt := &cos.PresignedURLOptions{
		Query: &url.Values{},
		Header: &http.Header{
			//设置上传文件类型
			"Content-Type": []string{mime.TypeByExtension(ext)},
		},
	}

	//添加session
	opt.Query.Add("x-cos-security-token", token.SessionToken)
	//获取上传文件预签名Url
	putURL, err := c.Object.GetPresignedURL(context.Background(), http.MethodPut, key, token.TmpSecretID, token.TmpSecretKey, time.Hour, opt)
	if err != nil {
		logrus.Error("上传失败", err.Error())
		return gin.H{
			"code": -1,
			"msg":  "获取上传预签名失败",
		}
	}
	opt2 := &cos.PresignedURLOptions{
		Query:  &url.Values{},
		Header: &http.Header{},
	}

	//添加session
	opt2.Query.Add("x-cos-security-token", token.SessionToken)
	//获取查看文件预签名Url
	getURL, errGet := c.Object.GetPresignedURL(context.Background(), http.MethodGet, key, token.TmpSecretID, token.TmpSecretKey, time.Hour, opt2)
	if errGet != nil {
		logrus.Error("获取失败", errGet.Error())
		return gin.H{
			"code": -1,
			"msg":  "获取下载预签名失败",
		}
	}

	return gin.H{
		"code": 200,
		"msg":  "success",
		"key":  key,
		"get":  getURL.String(),
		"put":  putURL.String(),
	}
}
