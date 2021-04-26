package app_test

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/theplant/gofixtures"
	"github.com/theplant/testingutils"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/sound_box"
	"github.com/zsy-cn/bms/sound_box/app"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/migration"
)

var thedb *gorm.DB
var logger = log.NewLogger(os.Stdout)
var ss sound_box.SoundBoxService

var sqlStr = `
insert into device_types(id, key, name, is_sensor, properties)
select 1, 'sound_box', 'IP音箱', true, '[]'
;
insert into device_models(id, name, manufacturer_id, device_type_id)
select 1, '音箱设备型号01', 1, 1
;
insert into groups(id, name, device_type_id, customer_id)
select 1, '音箱分组01', 1, 1
;
insert into devices(id, created_at, name, position, serial_number, customer_id, group_id, device_type_id, device_model_id, latitude, longitude)
select 			 1, '2019-01-01 12:00:00'::timestamp, '音箱测试设备01', '设备地址01', '0000000000000001', 1, 1, 1, 1, 10, 20
union all select 2, '2019-01-01 12:00:00'::timestamp, '音箱测试设备02', '设备地址01', '0000000000000002', 1, 1, 1, 1, 30, 40
;
`

var data = gofixtures.Data(gofixtures.Sql(sqlStr, []string{"device_types", "device_models", "groups", "devices", "sound_boxes"}))

func init() {
	var err error
	connectStr := "host=localhost port=7723 user=lora_backend dbname=lora_backend_test sslmode=disable password=123456"
	thedb, err = gorm.Open("postgres", connectStr)
	if err != nil {
		panic(err)
	}
	err = thedb.Exec("set time zone 'Asia/Shanghai';").Error
	if err != nil {
		panic(err)
	}
	logger.Info("start to init database tables")
	migration.InitTables(thedb)
	data.TruncatePut(thedb)
	logger.Info("init database tables completed")
	ss, err = app.NewSoundBoxService(*logger, thedb, sound_box.SoundBoxConfig{})
	if err != nil {
		panic(err)
	}
}

func TestGetDevices(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.DeviceList{
		List: []*protos.Device{
			{
				ID:            2,
				Name:          "音箱测试设备02",
				SerialNumber:  "0000000000000002",
				Position:      "设备地址01",
				Group:         "音箱分组01",
				GroupID:       1,
				DeviceType:    "IP音箱",
				DeviceTypeID:  1,
				DeviceModel:   "音箱设备型号01",
				DeviceModelID: 1,
				CustomerID:    1,
				Latitude:      30,
				Longitude:     40,
				Actived:       "true",
				CreatedAt:     "2019-01-01",
			},
			{
				ID:            1,
				Name:          "音箱测试设备01",
				SerialNumber:  "0000000000000001",
				Position:      "设备地址01",
				Group:         "音箱分组01",
				GroupID:       1,
				DeviceType:    "IP音箱",
				DeviceTypeID:  1,
				DeviceModel:   "音箱设备型号01",
				DeviceModelID: 1,
				CustomerID:    1,
				Latitude:      10,
				Longitude:     20,
				Actived:       "true",
				CreatedAt:     "2019-01-01",
			},
		},
		Count:       2,
		CurrentPage: 1,
		PageSize:    10,
		TotalCount:  2,
	}

	req := &protos.GetDevicesRequestForCustomer{
		Pagination: &protos.Pagination{
			Page:     1,
			PageSize: 10,
		},
		CustomerID: 1,
	}

	actual, err := ss.GetDevices(req)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}
func TestGetDevice(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.Device{
		ID:            1,
		Name:          "音箱测试设备01",
		SerialNumber:  "0000000000000001",
		Position:      "设备地址01",
		Group:         "音箱分组01",
		GroupID:       1,
		DeviceType:    "IP音箱",
		DeviceTypeID:  1,
		DeviceModel:   "音箱设备型号01",
		DeviceModelID: 1,
		CustomerID:    1,
		Latitude:      10,
		Longitude:     20,
		Actived:       "true",
		CreatedAt:     "2019-01-01",
	}

	actual, err := ss.GetDevice(1, "0000000000000001")
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetSoundBoxDeviceGroups(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.SoundBoxDeviceGroupList{
		List: []*protos.SoundBoxDeviceGroup{
			{
				ID:          1,
				Name:        "音箱分组01",
				CustomerID:  1,
				Status:      "关",
				DeviceTotal: 2,
			},
		},
		Count:       1,
		CurrentPage: 1,
		PageSize:    10,
		TotalCount:  1,
	}

	req := &protos.GetDevicesRequestForCustomer{
		Pagination: &protos.Pagination{
			Page:     1,
			PageSize: 10,
		},
		CustomerID: 1,
	}

	actual, err := ss.GetSoundBoxDeviceGroups(req)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}
