package app

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/device_management"
	"github.com/zsy-cn/bms/device_management/internal"
	"github.com/zsy-cn/bms/util/log"
)

func NewDeviceManagementService(
	l log.Logger,
	db *gorm.DB,
	cfg device_management.DeviceManagementConfig,
) (device_management.DeviceManagementService, error) {
	return internal.NewDeviceManagementService(l, db, cfg)
}
