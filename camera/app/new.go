package app

import (
	"github.com/zsy-cn/bms/camera"
	"github.com/zsy-cn/bms/camera/internal"
	"github.com/zsy-cn/bms/util/log"
)

func NewCameraService(
	l log.Logger,
	cfg camera.CameraConfig,
) (camera.CameraService, error) {
	return internal.NewCameraService(l, cfg)
}
