package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/environ_monitor"
	"github.com/zsy-cn/bms/geomagnetic"
	"github.com/zsy-cn/bms/manhole_cover"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/smoke"
	"github.com/zsy-cn/bms/sos"
	"github.com/zsy-cn/bms/sound_box"
	"github.com/zsy-cn/bms/trashcan"
	"github.com/zsy-cn/bms/water_level"
)

// AllModels ...
var AllModels = []interface{}{
	&model.Manufacturer{}, &model.DeviceType{}, &model.DeviceModel{}, &model.Group{},
	&model.Device{}, &model.Sensor{},
	&model.Customer{}, &model.Contact{},
	&model.Rule{}, &model.Notification{},
	&model.Manager{}, &model.Role{}, &model.Permission{},
}

// InitTables ...
func InitTables(db *gorm.DB) {
	err := db.AutoMigrate(AllModels...).Error
	if err != nil {
		panic(err)
	}

	sound_box.Migrate(db)
	environ_monitor.Migrate(db)
	geomagnetic.Migrate(db)
	manhole_cover.Migrate(db)
	smoke.Migrate(db)
	sos.Migrate(db)
	trashcan.Migrate(db)
	water_level.Migrate(db)
}

// SetupBaseDatas ...
func SetupBaseDatas(db *gorm.DB) {
	initManufacturer(db)
	initDeviceType(db)
	initRole(db)
	initManager(db)
}

func hasRecords(db *gorm.DB, table string) bool {
	var count int
	db.Table(table).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// initManufacturer ...
func initManufacturer(db *gorm.DB) {
	if hasRecords(db, "manufacturers") {
		return
	}

	manufacturers := []*model.Manufacturer{
		&model.Manufacturer{Name: "唯传"},
		&model.Manufacturer{Name: "利尔达"},
		&model.Manufacturer{Name: "盈飞驰"},
		&model.Manufacturer{Name: "GTI光通"},
	}

	for _, m := range manufacturers {
		err := db.Create(m).Error
		if err != nil {
			logger.Errorf("Create manufacture records failed: " + err.Error())
		}
	}
}

// initDeviceType ...
func initDeviceType(db *gorm.DB) {
	if hasRecords(db, "device_types") {
		return
	}
	deviceTypes := []*model.DeviceType{
		&model.DeviceType{
			Key:      "environ_monitor",
			Name:     "环境监测传感器",
			IsSensor: true,
			Properties: &model.StringArray{
				"temperature", "pm2.5", "noise",
			},
		},
		&model.DeviceType{
			Key:      "geomagnetic",
			Name:     "地磁传感器",
			IsSensor: true,
			Properties: &model.StringArray{
				"value",
			},
		},
		&model.DeviceType{
			Key:      "trashcan",
			Name:     "垃圾箱监测传感器",
			IsSensor: true,
			Properties: &model.StringArray{
				"percent",
			},
		},
		&model.DeviceType{
			Key:      "sos",
			Name:     "SOS报警器",
			IsSensor: true,
			Properties: &model.StringArray{
				"value",
			},
		},
		&model.DeviceType{
			Key:      "water_level",
			Name:     "水位监测传感器",
			IsSensor: true,
			Properties: &model.StringArray{
				"value",
			},
		},
		&model.DeviceType{
			Key:      "smoke",
			Name:     "火灾报警器",
			IsSensor: true,
			Properties: &model.StringArray{
				"value",
			},
		},
		&model.DeviceType{
			Key:      "manhole_cover",
			Name:     "井盖监测器",
			IsSensor: true,
			Properties: &model.StringArray{
				"value",
			},
		},
		&model.DeviceType{
			Key:        "lamp",
			Name:       "单灯",
			Properties: &model.StringArray{},
		},
		&model.DeviceType{
			Key:        "sound_box",
			Name:       "IP音箱",
			Properties: &model.StringArray{},
		},
		&model.DeviceType{
			Key:        "router",
			Name:       "Wifi路由器",
			Properties: &model.StringArray{},
		},
		&model.DeviceType{
			Key:        "camera",
			Name:       "摄像头",
			Properties: &model.StringArray{},
		},
	}

	for _, d := range deviceTypes {
		err := db.Create(d).Error
		if err != nil {
			logger.Errorf("Create device type records failed: " + err.Error())
		}
	}
}

// initRole ...
func initRole(db *gorm.DB) {
	if hasRecords(db, "roles") {
		return
	}
	roles := []*model.Role{
		&model.Role{Name: "超级管理员", PermissionIDs: &model.Uint64Array{}},
		&model.Role{Name: "管理员", PermissionIDs: &model.Uint64Array{}},
		&model.Role{Name: "现场安装", PermissionIDs: &model.Uint64Array{}},
	}

	for _, r := range roles {
		err := db.Create(r).Error
		if err != nil {
			logger.Errorf("Create role records failed: " + err.Error())
		}
	}
}

// initManager ...
func initManager(db *gorm.DB) {
	if hasRecords(db, "managers") {
		return
	}
	managers := []*model.Manager{
		&model.Manager{
			Name:        "admin",
			Passwd:      "123456",
			DisplayName: "TaTa",
			Phone:       "12345678901",
			RoleID:      1, // 超级管理
		},
	}

	for _, r := range managers {
		err := db.Create(r).Error
		if err != nil {
			logger.Errorf("Create manager records failed: " + err.Error())
		}
	}
}
