package app_test

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/theplant/gofixtures"
	"github.com/theplant/testingutils"

	"github.com/zsy-cn/bms/device_management"
	"github.com/zsy-cn/bms/device_management/app"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/migration"
)

var thedb *gorm.DB
var logger = log.NewLogger(os.Stdout)
var dm device_management.DeviceManagementService

func init() {
	var err error
	connectStr := "host=postgres-serv port=5432 user=backend dbname=backend sslmode=disable password=backend"
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
	logger.Info("init database tables completed")
	dm, err = app.NewDeviceManagementService(*logger, thedb, device_management.DeviceManagementConfig{})
	if err != nil {
		panic(err)
	}
}

var sqlStr = `
insert into device_types(id, key, name, is_sensor, properties)
select 1, 'sos', 'SOS传感器', true, '[]'
union all select 2, 'water_level', '水位报警器', true, '["value"]'
;
insert into device_models(id, name, manufacturer_id, device_type_id)
select 1, 'SOS设备型号01', 1, 1
union all select 2, '水位设备型号01', 1, 2
;
insert into groups(id, name, device_type_id, customer_id)
select 1, 'SOS测试分组01', 1, 1
union all select 2, 'SOS测试分组02', 1, 1
union all select 3, '水位测试分组01', 2, 1
;
insert into rules(id, created_at, device_type_id, customer_id, device_sn, section)
select 1, '2019-01-01 12:00:00'::timestamp, 2, 1, '0000000000000005', '[1.2,2.3,3.4,4.5]'
;
insert into devices(id, created_at, updated_at, name, position, serial_number, customer_id, group_id, device_type_id, device_model_id, latitude, longitude)
select 1, '2019-01-01 12:00:00'::timestamp, '2019-01-01 12:00:00'::timestamp, 'SOS测试设备01', '设备地址01', '0000000000000001', 1, 1, 1, 1, 0, 0
union all select 2, '2019-01-01 12:00:00'::timestamp, '2019-01-01 12:00:00'::timestamp, 'SOS测试设备02', '设备地址01', '0000000000000002', 1, 1, 1, 1, 0, 0
union all select 3, '2019-01-01 12:00:00'::timestamp, '2019-01-01 12:00:00'::timestamp, 'SOS测试设备03', '设备地址02', '0000000000000003', 1, 2, 1, 1, 0, 0
union all select 4, '2019-01-01 12:00:00'::timestamp, '2019-01-01 12:00:00'::timestamp, 'SOS测试设备04', '设备地址02', '0000000000000004', 1, 2, 1, 1, 0, 0
union all select 5, '2019-01-01 12:00:00'::timestamp, '2019-01-01 12:00:00'::timestamp, '水位测试设备01', '设备地址11', '0000000000000005', 1, 3, 2, 2, 0, 0
;
`

var data = gofixtures.Data(gofixtures.Sql(sqlStr, []string{"device_types", "device_models", "groups", "rules", "devices"}))

func TestGetDevices(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.DeviceList{
		List: []*protos.Device{
			&protos.Device{
				ID:            2,
				Name:          "SOS测试设备02",
				Position:      "设备地址01",
				SerialNumber:  "0000000000000002",
				CustomerID:    1,
				GroupID:       1,
				Group:         "SOS测试分组01",
				DeviceTypeID:  1,
				DeviceType:    "SOS传感器",
				DeviceModelID: 1,
				DeviceModel:   "SOS设备型号01",
				CreatedAt:     "2019-01-01",
				Actived:       "true",
			},
			&protos.Device{
				ID:            1,
				Name:          "SOS测试设备01",
				Position:      "设备地址01",
				SerialNumber:  "0000000000000001",
				CustomerID:    1,
				GroupID:       1,
				Group:         "SOS测试分组01",
				DeviceTypeID:  1,
				DeviceType:    "SOS传感器",
				DeviceModelID: 1,
				DeviceModel:   "SOS设备型号01",
				CreatedAt:     "2019-01-01",
				Actived:       "true",
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
		GroupID:    1,
	}

	actual, err := dm.GetDevices(req)
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
		Name:          "SOS测试设备01",
		Position:      "设备地址01",
		SerialNumber:  "0000000000000001",
		CustomerID:    1,
		GroupID:       1,
		Group:         "SOS测试分组01",
		DeviceTypeID:  1,
		DeviceType:    "SOS传感器",
		DeviceModelID: 1,
		DeviceModel:   "SOS设备型号01",
		CreatedAt:     "2019-01-01",
		Actived:       "true",
	}
	actual, err := dm.GetDevice("0000000000000001", 1, 1)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetDeviceGroupsByType(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.GetDeviceGroupsByTypeResponse{
		Count: 2,
		List: []*protos.DeviceGroupByType{
			&protos.DeviceGroupByType{
				ID:    1,
				Name:  "SOS传感器",
				Key:   "sos",
				Count: 4,
			},
			&protos.DeviceGroupByType{
				ID:    2,
				Name:  "水位报警器",
				Key:   "water_level",
				Count: 1,
			},
		},
	}
	actual, err := dm.GetDeviceGroupsByType(1)
	if err != nil {
		t.Error(err)
	}
	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetSafetyDevices(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.SafetyDeviceList{
		List: []*protos.SafetyDevice{
			{
				ID:            2,
				Name:          "SOS测试设备02",
				SerialNumber:  "0000000000000002",
				Position:      "设备地址01",
				Group:         "SOS测试分组01",
				GroupID:       1,
				DeviceType:    "SOS传感器",
				DeviceTypeID:  1,
				DeviceModel:   "SOS设备型号01",
				DeviceModelID: 1,
				CustomerID:    1,
				Actived:       "true",
				CreatedAt:     "2019-01-01",
				StatusCode:    2,
				Status:        "离线",
			},
			{
				ID:            1,
				Name:          "SOS测试设备01",
				SerialNumber:  "0000000000000001",
				Position:      "设备地址01",
				Group:         "SOS测试分组01",
				GroupID:       1,
				DeviceType:    "SOS传感器",
				DeviceTypeID:  1,
				DeviceModel:   "SOS设备型号01",
				DeviceModelID: 1,
				CustomerID:    1,
				Actived:       "true",
				CreatedAt:     "2019-01-01",
				StatusCode:    2,
				Status:        "离线",
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
		GroupID:    1,
	}

	actual, err := dm.GetSafetyDevices(req, "sos", "-24h")
	if err != nil {
		t.Error(err)
	}
	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetDeviceInfos(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.DeviceCountInfo{
		Amount: 4,
		Alive:  0,
		Alarm:  0,
	}
	req := &protos.GetDevicesRequestForCustomer{
		CustomerID:   1,
		DeviceTypeID: 1,
	}

	actual, err := dm.GetDeviceInfos(req, "sos", "-24h")
	if err != nil {
		t.Error(err)
	}
	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetSafetyDeviceStatus(t *testing.T) {
	data.TruncatePut(thedb)

	actual, err := dm.GetSafetyDeviceStatus(1, "sos", "sos", "0000000000000001")
	if err != nil {
		t.Error(err)
	}
	expected := 2 // 离线
	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetDeviceAlarmThresholds(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.AlarmThresholdList{
		List: []*protos.DeviceWithAlarmThreshold{
			{
				ID:              5,
				Name:            "水位测试设备01",
				Position:        "设备地址11",
				SerialNumber:    "0000000000000005",
				GroupID:         3,
				Group:           "水位测试分组01",
				DeviceModelID:   2,
				CreatedAt:       "2019-01-01 12:00:00",
				AlarmThresholds: []float64{1.2, 2.3, 3.4, 4.5},
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
		CustomerID:   1,
		DeviceTypeID: 2,
	}

	actual, err := dm.GetDeviceAlarmThresholds(req)
	if err != nil {
		t.Error(err)
	}
	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestSetDeviceAlarmThreshold(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.AlarmThresholdList{
		List: []*protos.DeviceWithAlarmThreshold{
			{
				ID:              5,
				Name:            "水位测试设备01",
				Position:        "设备地址11",
				SerialNumber:    "0000000000000005",
				GroupID:         3,
				Group:           "水位测试分组01",
				DeviceModelID:   2,
				CreatedAt:       "2019-01-01 12:00:00",
				AlarmThresholds: []float64{10, 20, 30, 40},
			},
		},
		Count:       1,
		CurrentPage: 1,
		PageSize:    10,
		TotalCount:  1,
	}

	req1 := &protos.SetAlarmThresholdRequest{
		CustomerID:      1,
		DeviceTypeID:    2,
		DeviceSN:        "0000000000000005",
		AlarmThresholds: []float64{10, 20, 30, 40},
	}

	err := dm.SetDeviceAlarmThreshold(req1)
	if err != nil {
		t.Error(err)
	}
	req2 := &protos.GetDevicesRequestForCustomer{
		Pagination: &protos.Pagination{
			Page:     1,
			PageSize: 10,
		},
		CustomerID:   1,
		DeviceTypeID: 2,
	}
	actual, err := dm.GetDeviceAlarmThresholds(req2)
	if err != nil {
		t.Error(err)
	}
	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestDeactiveDevice(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.Device{
		ID:            1,
		Name:          "SOS测试设备01",
		Position:      "设备地址01",
		SerialNumber:  "0000000000000001",
		CustomerID:    1,
		GroupID:       1,
		Group:         "SOS测试分组01",
		DeviceTypeID:  1,
		DeviceType:    "SOS传感器",
		DeviceModelID: 1,
		DeviceModel:   "SOS设备型号01",
		CreatedAt:     "2019-01-01",
		Actived:       "false",
	}
	err := dm.DeactiveDevice("0000000000000001", 1)
	if err != nil {
		t.Error(err)
	}
	actual, err := dm.GetDevice("0000000000000001", 1, 1)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}

	// 然后验证激活接口
	err2 := dm.ActiveDevice("0000000000000001", 1)
	if err2 != nil {
		t.Error(err2)
	}
	actual2, err2 := dm.GetDevice("0000000000000001", 1, 1)
	if err2 != nil {
		t.Error(err2)
	}
	expected.Actived = "true"
	diff2 := testingutils.PrettyJsonDiff(expected, actual2)
	if len(diff2) > 0 {
		t.Error(diff2)
	}

}
