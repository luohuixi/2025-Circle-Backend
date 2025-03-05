package service

import (
	"context"

	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/spf13/viper"

	//"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	//"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	//"os"
	"time"

	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
	//"gopkg.in/yaml.v3"
)
type QiNiuYunConfig struct {
    AccessKey string `yaml:"access_key"` 
    SecretKey string `yaml:"secret_key"`
    Bucket    string `yaml:"bucket"`   
    Domain    string `yaml:"domain"`   
}

func ReadConfig(filename string) (*QiNiuYunConfig, error) {
    viper.SetConfigName("muxiconfig")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".") // 设置配置文件所在路径
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }
    var config QiNiuYunConfig
    config.AccessKey = viper.GetString("access_key")
    config.SecretKey = viper.GetString("secret_key")
    config.Bucket = viper.GetString("bucket")
    config.Domain = viper.GetString("domain")
    // file, err := os.Open(filename)  //打开文件
    // if err != nil {
    //    return nil, err
    // }
    // defer file.Close()

    // var config QiNiuYunConfig
    // decoder := yaml.NewDecoder(file) //读取yaml格式的文件
    // err = decoder.Decode(&config)  //Decode()方法将文件内容解析到结构体中
    // if err != nil {
    //    return nil, err
    // }

     return &config, nil
}

func GetToken(qiniuyun *QiNiuYunConfig) (string, error) {
    accessKey := qiniuyun.AccessKey
    secretKey := qiniuyun.SecretKey
    bucket := qiniuyun.Bucket
    // mac是一个身份验证的工具
    mac := credentials.NewCredentials(accessKey, secretKey)

    // 生成一个创建策略，包含有效时间以及存储对象
    putPolicy, err := uptoken.NewPutPolicy(bucket, time.Now().Add(2*time.Hour))
    if err != nil {
       return "", err
    }
    // 获取上传凭证
    upToken, err := uptoken.NewSigner(putPolicy, mac).GetUpToken(context.Background())
    if err != nil {
       return upToken, err
    }
    return upToken, nil
}