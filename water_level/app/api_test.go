package app_test

import (
	"os"
	"testing"
	"time"

	"bou.ke/monkey"

	"github.com/jinzhu/gorm"
	"github.com/theplant/gofixtures"
	"github.com/theplant/testingutils"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/migration"
	"github.com/zsy-cn/bms/water_level"
	"github.com/zsy-cn/bms/water_level/app"
)

var thedb *gorm.DB
var logger = log.NewLogger(os.Stdout)
var ws water_level.WaterLevelService

var sqlStr = `
insert into device_types(id, key, name, is_sensor, properties)
select 1, 'water_level', '水位传感器', true, '["status"]'
;
insert into device_models(id, name, manufacturer_id, device_type_id)
select 1, '水位设备型号01', 1, 1
;
insert into groups(id, name, device_type_id, customer_id)
select 1, '水位分组01', 1, 1
;
insert into devices(id, created_at, name, position, serial_number, customer_id, group_id, device_type_id, device_model_id, latitude, longitude)
select 			 1, '2019-01-01 12:00:00'::timestamp, '水位测试设备01', '设备地址01', '0000000000000001', 1, 1, 1, 1, 10, 20
union all select 2, '2019-01-01 12:00:00'::timestamp, '水位测试设备02', '设备地址01', '0000000000000002', 1, 1, 1, 1, 30, 40
;
insert into water_levels(id, created_at, device_sn, group_id, customer_id, status)
select 			 1, '2019-01-01 06:00:00'::timestamp, '0000000000000001', 1, 1, 0
union all select 2, '2019-01-01 06:00:00'::timestamp, '0000000000000002', 1, 1, 1
union all select 3, '2019-01-01 12:00:00'::timestamp, '0000000000000001', 1, 1, 1
union all select 4, '2019-01-01 12:00:00'::timestamp, '0000000000000002', 1, 1, 0
;
insert into rules(id, created_at, customer_id, device_type_id, device_sn, section)
select 			 1, '2019-01-01 00:00:00'::timestamp, 1, 1, '0000000000000001', '[5,10,60,80]'
union all select 2, '2019-01-01 00:00:00'::timestamp, 1, 1, '0000000000000002', '[5,10,60,80]'
;
insert into notifications(id, created_at, updated_at, device_sn, msg_id, group_id, customer_id, device_type_id, key, content, solved)
select 			 1, '2019-01-01 06:00:00'::timestamp, '2019-01-01 06:00:00'::timestamp, '0000000000000001', 2, 1, 1, 1, 'water_level_high_alert', '水位高于 80 cm', false
union all select 2, '2019-01-01 12:00:00'::timestamp, '2019-01-01 06:00:00'::timestamp, '0000000000000002', 3, 1, 1, 1, 'water_level_low_alert', '水位低于 5 cm', false
;
`

var data = gofixtures.Data(gofixtures.Sql(sqlStr, []string{"device_types", "device_models", "groups", "devices", "water_levels", "notifications", "rules"}))

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
	ws, err = app.NewWaterLevelService(*logger, thedb, water_level.WaterLevelConfig{})
	if err != nil {
		panic(err)
	}
}

func TestGetDevices(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.SafetyDeviceList{
		List: []*protos.SafetyDevice{
			{
				ID:            2,
				Name:          "水位测试设备02",
				SerialNumber:  "0000000000000002",
				Position:      "设备地址01",
				Group:         "水位分组01",
				GroupID:       1,
				DeviceType:    "水位传感器",
				DeviceTypeID:  1,
				DeviceModel:   "水位设备型号01",
				DeviceModelID: 1,
				CustomerID:    1,
				Latitude:      30,
				Longitude:     40,
				Actived:       "true",
				CreatedAt:     "2019-01-01",
				StatusCode:    2,
				Status:        "离线",
			},
			{
				ID:            1,
				Name:          "水位测试设备01",
				SerialNumber:  "0000000000000001",
				Position:      "设备地址01",
				Group:         "水位分组01",
				GroupID:       1,
				DeviceType:    "水位传感器",
				DeviceTypeID:  1,
				DeviceModel:   "水位设备型号01",
				DeviceModelID: 1,
				CustomerID:    1,
				Latitude:      10,
				Longitude:     20,
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
	}

	actual, err := ws.GetDevices(req)
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

	expected := &protos.SafetyDevice{
		ID:            1,
		Name:          "水位测试设备01",
		SerialNumber:  "0000000000000001",
		Position:      "设备地址01",
		Group:         "水位分组01",
		GroupID:       1,
		DeviceType:    "水位传感器",
		DeviceTypeID:  1,
		DeviceModel:   "水位设备型号01",
		DeviceModelID: 1,
		CustomerID:    1,
		Latitude:      10,
		Longitude:     20,
		Actived:       "true",
		CreatedAt:     "2019-01-01",
		StatusCode:    2,
		Status:        "离线",
	}

	actual, err := ws.GetDevice(1, "0000000000000001")
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
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fakeNow, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-01-01 12:00:01", loc)
	timenowPatcher := monkey.Patch(time.Now, func() time.Time {
		return fakeNow
	})
	defer timenowPatcher.Unpatch()

	expected1 := &protos.DeviceCountInfo{
		Amount: 2,
		Alive:  2,
		Alarm:  2,
	}

	req := &protos.GetDevicesRequestForCustomer{
		CustomerID: 1,
	}
	actual1, err := ws.GetDeviceInfos(req)
	if err != nil {
		t.Error(err)
	}

	diff1 := testingutils.PrettyJsonDiff(expected1, actual1)
	if len(diff1) > 0 {
		t.Error(diff1)
	}
	//////////////////////////////////////////////////////////////////
	expected2 := &protos.DeviceCountInfo{
		Amount: 2,
		Alive:  2,
		Alarm:  1,
	}
	err = thedb.Model(&model.Notification{}).Where("id = 1").Updates(map[string]interface{}{"solved": true}).Error
	if err != nil {
		t.Error(err)
	}

	actual2, err := ws.GetDeviceInfos(req)
	if err != nil {
		t.Error(err)
	}

	diff2 := testingutils.PrettyJsonDiff(expected2, actual2)
	if len(diff2) > 0 {
		t.Error(diff2)
	}
}

func TestGetDeviceStatus(t *testing.T) {
	data.TruncatePut(thedb)
	var customerID uint64 = 1
	deviceSN := "0000000000000001"

	////////////////////////////////////////////////////////// 离线
	expected1 := 2

	actual1, err := ws.GetDeviceStatus(customerID, deviceSN)
	if err != nil {
		t.Error(err)
	}

	diff1 := testingutils.PrettyJsonDiff(expected1, actual1)
	if len(diff1) > 0 {
		t.Error(diff1)
	}
	////////////////////////////////////////////////////////// 报警
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fakeNow, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-01-01 12:00:01", loc)
	timenowPatcher := monkey.Patch(time.Now, func() time.Time {
		return fakeNow
	})
	defer timenowPatcher.Unpatch()

	expected2 := 3
	actual2, err := ws.GetDeviceStatus(customerID, deviceSN)
	if err != nil {
		t.Error(err)
	}

	diff2 := testingutils.PrettyJsonDiff(expected2, actual2)
	if len(diff2) > 0 {
		t.Error(diff2)
	}
	////////////////////////////////////////////////////////// 正常
	expected3 := 1
	err = thedb.Model(&model.Notification{}).Where("device_sn = ?", deviceSN).Updates(map[string]interface{}{"solved": true}).Error
	if err != nil {
		t.Error(err)
	}

	actual3, err := ws.GetDeviceStatus(customerID, deviceSN)
	if err != nil {
		t.Error(err)
	}

	diff3 := testingutils.PrettyJsonDiff(expected3, actual3)
	if len(diff3) > 0 {
		t.Error(diff3)
	}
}

func TestGetDeviceAlarmThresholds(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.WaterLevelAlarmThresholdList{
		List: []*protos.WaterLevelDeviceWithAlarmThreshold{
			{
				ID:            2,
				Name:          "水位测试设备02",
				Position:      "设备地址01",
				SerialNumber:  "0000000000000002",
				GroupID:       1,
				Group:         "水位分组01",
				CreatedAt:     "2019-01-01 00:00:00",
				LowStageWarn:  5,
				LowStageInfo:  10,
				HighStageInfo: 60,
				HighStageWarn: 80,
			},
			{
				ID:            1,
				Name:          "水位测试设备01",
				Position:      "设备地址01",
				SerialNumber:  "0000000000000001",
				GroupID:       1,
				Group:         "水位分组01",
				CreatedAt:     "2019-01-01 00:00:00",
				LowStageWarn:  5,
				LowStageInfo:  10,
				HighStageInfo: 60,
				HighStageWarn: 80,
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

	actual, err := ws.GetDeviceAlarmThresholds(req)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestSetDeviceAlarmThresholds(t *testing.T) {
	data.TruncatePut(thedb)

	var customerID uint64 = 1
	deviceSN := "0000000000000001"
	alarmThresholds := []float64{10, 20, 90, 100}

	req1 := &protos.SetAlarmThresholdRequest{
		DeviceSN:        deviceSN,
		CustomerID:      customerID,
		AlarmThresholds: alarmThresholds,
	}
	err := ws.SetDeviceAlarmThreshold(req1)
	if err != nil {
		t.Error(err)
	}

	expected := &protos.WaterLevelAlarmThresholdList{
		List: []*protos.WaterLevelDeviceWithAlarmThreshold{
			{
				ID:            2,
				Name:          "水位测试设备02",
				Position:      "设备地址01",
				SerialNumber:  "0000000000000002",
				GroupID:       1,
				Group:         "水位分组01",
				CreatedAt:     "2019-01-01 00:00:00",
				LowStageWarn:  5,
				LowStageInfo:  10,
				HighStageInfo: 60,
				HighStageWarn: 80,
			},
			{
				ID:            1,
				Name:          "水位测试设备01",
				Position:      "设备地址01",
				SerialNumber:  "0000000000000001",
				GroupID:       1,
				Group:         "水位分组01",
				CreatedAt:     "2019-01-01 00:00:00",
				LowStageWarn:  10,
				LowStageInfo:  20,
				HighStageInfo: 90,
				HighStageWarn: 100,
			},
		},
		Count:       2,
		CurrentPage: 1,
		PageSize:    10,
		TotalCount:  2,
	}

	req2 := &protos.GetDevicesRequestForCustomer{
		Pagination: &protos.Pagination{
			Page:     1,
			PageSize: 10,
		},
		CustomerID: customerID,
	}

	actual, err := ws.GetDeviceAlarmThresholds(req2)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}
