package Services

import (
	"GoRealTimeChat/DataModels/Internal/ConfigModels"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiNiuYunService struct {
}

// 上传文件到七牛云

func (QiNiuYunService) UploadFile(localFile string, fileName string) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: ConfigModels.ConfigData.QiNiu.Bucket,
	}

	mac := qbox.NewMac(ConfigModels.ConfigData.QiNiu.AccessKey, ConfigModels.ConfigData.QiNiu.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	config := storage.Config{
		Zone:          &storage.ZoneHuanan, // 设置存储区域为华南
		UseHTTPS:      false,               // 设置不使用 HTTPS
		UseCdnDomains: false,               // 设置不使用 CDN 域名
	}

	formUploader := storage.NewFormUploader(&config) // 创建表单上传器
	ret := storage.PutRet{}                          // 存储上传结果的结构体

	// 可选配置
	putExtra := storage.PutExtra{}                                                                   // 上传的额外配置
	err := formUploader.PutFile(context.Background(), &ret, upToken, fileName, localFile, &putExtra) // 使用表单上传器上传文件
	if err != nil {
		fmt.Println(err) // 输出错误信息
		return "", err   // 返回错误
	}

	return ConfigModels.ConfigData.QiNiu.Domain + "/" + fileName, nil // 返回文件的访问地址
}
