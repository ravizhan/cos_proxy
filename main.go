package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"time"
)

func main() {
	type Config struct {
		BucketUrl string
		SecretId  string
		SecretKey string
		Suffix    string
		Port      string
	}
	gin.SetMode(gin.ReleaseMode)
	_, err := os.Stat("config.json")
	if os.IsNotExist(err) {
		fmt.Println("未检测到配置文件")
		data := Config{"", "", "", "", ""}
		marshalData, _ := json.MarshalIndent(data, "", "\t")
		err = os.WriteFile("config.json", marshalData, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println("生成配置文件config.json,请全部填写完毕后启动")
		return
	}
	jsonData, _ := os.ReadFile("config.json")
	var config Config
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		panic(err)
	}
	value := reflect.ValueOf(config)
	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).String() == "" {
			fmt.Println("配置文件填写不完全")
			return
		}
	}
	u, _ := url.Parse(config.BucketUrl)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.SecretId,
			SecretKey: config.SecretKey,
		},
	})
	ctx := context.Background()
	r := gin.Default()
	r.GET("/cos/:image", func(c *gin.Context) {
		imageName := c.Param("image")
		presignedURL, err := client.Object.GetPresignedURL(ctx, http.MethodGet, imageName+config.Suffix, config.SecretId, config.SecretKey, time.Hour, nil)
		if err != nil {
			panic(err)
		}
		resp, err := http.Get(presignedURL.String())
		if err != nil {
			panic(err)
		}
		if resp.StatusCode == 200 {
			img, _ := io.ReadAll(resp.Body)
			contentType := resp.Header.Get("Content-Type")
			c.Data(200, contentType, img)
		} else {
			c.String(404, "404 page not found")
		}
	})
	fmt.Println("正在监听 0.0.0.0:" + config.Port)
	r.Run(":" + config.Port)
}
