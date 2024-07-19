/**
 * @Author tanchang
 * @Description //获取腾讯云cos临时密钥
 * @Date 2024/7/17 11:32
 * @File:  cosToken
 * @Software: GoLand
 **/

package utils

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"net/http"
	"net/url"
	"os"
	"time"
)

// CosToken 获取腾讯云cos临时密钥
func CosToken() (*sts.Credentials, error) {
	c := sts.NewClient(
		// 通过环境变量获取密钥, os.Getenv 方法表示获取环境变量
		os.Getenv("SECRETID"),
		os.Getenv("SECRETKEY"),
		nil,
	)
	// 策略概述 https://cloud.tencent.com/document/product/436/18023
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          "ap-hongkong",
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
						"name/cos:GetObject",
					},
					Effect: "allow",
					Resource: []string{
						//这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						//存储桶的命名格式为 BucketName-APPID，此处填写的 bucket 必须为此格式
						"qcs::cos: ap-hongkong:uid/" + os.Getenv("APPID") + ":" + os.Getenv("BUCKET") + "/upload/video/*",
					},
				},
			},
		},
	}
	res, err := c.GetCredential(opt)
	if err != nil {
		return nil, err
	}
	return res.Credentials, nil
}

// VideoURL 视频URL
func VideoURL(URL string) string {
	token, _ := CosToken()
	u, _ := url.Parse("https://govideo-1305907375.cos.ap-hongkong.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{})
	opt := &cos.PresignedURLOptions{
		Query:  &url.Values{},
		Header: &http.Header{},
	}
	//添加session
	opt.Query.Add("x-cos-security-token", token.SessionToken)
	getURL, _ := c.Object.GetPresignedURL(context.Background(), http.MethodGet, URL, token.TmpSecretID, token.TmpSecretKey, time.Hour, opt)
	return getURL.String()
}

// CoverURL 封面URL
func CoverURL(URL string) string {
	token, _ := CosToken()
	u, _ := url.Parse("https://govideo-1305907375.cos.ap-hongkong.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{})
	opt := &cos.PresignedURLOptions{
		Query:  &url.Values{},
		Header: &http.Header{},
	}
	//添加session
	opt.Query.Add("x-cos-security-token", token.SessionToken)
	getURL, _ := c.Object.GetPresignedURL(context.Background(), http.MethodGet, URL, token.TmpSecretID, token.TmpSecretKey, time.Hour, opt)
	return getURL.String()
}
