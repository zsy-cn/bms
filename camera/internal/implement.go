package internal

import (
	"github.com/robfig/cron"

	"github.com/zsy-cn/bms/camera"
	"github.com/zsy-cn/bms/util/log"
)

type DefaultCameraService struct {
	l           log.Logger
	accessToken string
	cameraID    string
	cfgPath     string
	appKey      string
	appSecret   string
	ys7Addr     string
}

// NewCameraService ...
func NewCameraService(
	l log.Logger,
	cfg camera.CameraConfig,
) (iService camera.CameraService, err error) {
	service := &DefaultCameraService{
		l:         l,
		cfgPath:   cfg.ConfigPath,
		appKey:    cfg.AppKey,
		appSecret: cfg.AppSecret,
		ys7Addr:   cfg.YS7Addr,
	}
	cfgMap, err := service.readConfigFile()
	if err != nil {
		// 日志已经打印过
		return
	}
	service.accessToken = cfgMap["access-token"]
	service.cameraID = cfgMap["camera-id"]
	cronJob := cron.New()
	interval := "0 0 3 * * *"
	// interval := "* * * * * *"
	service.requestAccessToken()
	cronJob.AddFunc(interval, service.requestAccessToken)
	cronJob.Start()
	return service, nil
}

var _ camera.CameraService = (*DefaultCameraService)(nil)

// GetAccessToken ...
func (cs *DefaultCameraService) GetAccessToken() (accessToken string, err error) {
	return cs.accessToken, nil
}

// GetMainScreen ...
func (cs *DefaultCameraService) GetMainScreen() (cameraID string, err error) {
	return cs.cameraID, nil
}

// SetMainScreen ...
func (cs *DefaultCameraService) SetMainScreen(cameraID string) (err error) {
	err = cs.writeConfigFile("camera-id", cameraID)
	if err != nil {
		return
	}
	cs.cameraID = cameraID
	return
}
