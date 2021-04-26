package app

import (
	"errors"
	"os"
	"testing"

	"github.com/zsy-cn/bms/camera"
	"github.com/zsy-cn/bms/util/log"
)

var cameraService camera.CameraService
var _logger *log.Logger

func MustGetLogger() log.Logger {
	if _logger != nil {
		return *_logger
	}
	_logger := log.NewLogger(os.Stdout)
	_logger.Info("msg", "aigleconfig: Initialized logger")
	return *_logger
}

func init() {
	var err error
	cameraService, err = NewCameraService(
		MustGetLogger(),
		camera.CameraConfig{
			AppKey:     "",
			AppSecret:  "",
			YS7Addr:    "",
			ConfigPath: "",
		},
	)

	if err != nil {
		panic(err)
	}
}

func TestRequestAccessToken(t *testing.T) {
	accessToken, err := cameraService.RequestAccessToken()
	if err != nil {
		t.Error(err)
	}
	if len(accessToken) != 64 {
		msg := errors.New("access token length is not valid")
		t.Error(msg)
	}
	t.Logf("get access token: %s", accessToken)
}
