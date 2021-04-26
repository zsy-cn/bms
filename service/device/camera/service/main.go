package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/robfig/cron"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

// ICamera ...
type ICamera interface {
	GetAccessToken() (resp *protos.GetAccessTokenResponse, err error)
	SetMainScreen(req *protos.SetMainScreenRequest) (err error)
	GetMainScreen(req *protos.GetMainScreenRequest) (resp *protos.GetMainScreenResponse, err error)
}

// Camera 增删服务
type Camera struct {
	logger   *log.Logger
	redisCli *redis.Client
	cronJob  *cron.Cron
}

// New ...
func New(logger *log.Logger) (camera ICamera, err error) {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr + ":" + conf.RedisPort,
		Password: conf.RedisPassword,
		DB:       conf.RedisCameraServiceDB,
	})
	cronJob := cron.New()

	cameraServ := &Camera{
		logger:   logger,
		redisCli: redisCli,
		cronJob:  cronJob,
	}
	interval := "0 0 3 * * *"
	// interval := "* * * * * *"
	cameraServ.cronJob.AddFunc(interval, cameraServ.requestToken)
	cameraServ.cronJob.Start()
	return cameraServ, err
}

// GetAccessToken 从redis中查询access token的值
func (c *Camera) GetAccessToken() (resp *protos.GetAccessTokenResponse, err error) {
	accessToken, err := c.redisCli.Get("camera:ys7:accessToken").Result()
	if err != nil {
		c.logger.Errorf("get access token from redis failed: %s", err.Error())
		return
	}
	resp = &protos.GetAccessTokenResponse{
		AccessToken: accessToken,
	}
	return
}

// requestToken 向萤石请求access token, 并存入redis
func (c *Camera) requestToken() {
	c.logger.Debug("request access token in requestToken()")
	var err error
	contentType := "application/x-www-form-urlencoded"

	requestBody := map[string]string{
		"appKey":    conf.CameraServiceYS7AppKey,
		"appSecret": conf.CameraServiceYS7Secret,
	}
	byteData, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}
	jsonData := bytes.NewBuffer(byteData)

	resp, err := http.Post(conf.CameraServiceYS7Addr, contentType, jsonData)
	if err != nil {
		c.logger.Errorf("request ys to get access token failed: %s", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logger.Errorf("request ys to get access token failed: %s", err.Error())
		return
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, result)
	if err != nil {
		c.logger.Errorf("unmarshal access token response failed: %s", err.Error())
		return
	}

	if code, ok := result["code"]; !ok || code.(string) != "200" {
		c.logger.Errorf("we didn't get a valid access token: %s", result["msg"])
		return
	}

	resultDataField := result["data"].(map[string]string)
	accessToken := resultDataField["accessToken"]

	c.redisCli.Set("camera:ys7:accessToken", accessToken, 0)
	return
}

// SetMainScreen 设置
func (c *Camera) SetMainScreen(req *protos.SetMainScreenRequest) (err error) {
	var sessionID string
	sessionID = req.SessionID
	if sessionID == "" {
		sessionID = "mainscreen"
	}
	keyStr := fmt.Sprintf("camera:%s:%s", "mainscreen", sessionID)
	c.redisCli.Set(keyStr, req.CameraID, 0)
	return
}

// GetMainScreen ...
func (c *Camera) GetMainScreen(req *protos.GetMainScreenRequest) (resp *protos.GetMainScreenResponse, err error) {
	var sessionID string
	sessionID = req.SessionID
	if sessionID == "" {
		sessionID = "mainscreen"
	}
	keyStr := fmt.Sprintf("camera:%s:%s", "mainscreen", sessionID)
	cameraID, err := c.redisCli.Get(keyStr).Result()
	if err != nil {
		c.logger.Errorf("get main screen for session: %s failed: %s", sessionID, err.Error())
		return
	}
	resp = &protos.GetMainScreenResponse{
		CameraID: cameraID,
	}
	return
}
