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
	"github.com/zsy-cn/bms/trashcan"
	"github.com/zsy-cn/bms/trashcan/app"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/migration"
)

var thedb *gorm.DB
var logger = log.NewLogger(os.Stdout)
var ts trashcan.TrashcanService

var sqlStr = `
insert into device_types(id, key, name, is_sensor, properties)
select 1, 'trashcan', '垃圾箱传感器', true, '["status"]'
;
insert into device_models(id, name, manufacturer_id, device_type_id)
select 1, '垃圾箱设备型号01', 1, 1
;
insert into groups(id, name, device_type_id, customer_id)
select 1, '垃圾箱分组01', 1, 1
;
insert into devices(id, created_at, name, position, serial_number, customer_id, group_id, device_type_id, device_model_id, latitude, longitude)
select 			 1, '2019-01-01 12:00:00'::timestamp, '垃圾箱测试设备01', '设备地址01', '0000000000000001', 1, 1, 1, 1, 10, 20
union all select 2, '2019-01-01 12:00:00'::timestamp, '垃圾箱测试设备02', '设备地址01', '0000000000000002', 1, 1, 1, 1, 30, 40
;
insert into trashcans(id, created_at, device_sn, group_id, customer_id, msg_type, percent)
select 			 1, '2019-01-01 06:00:00'::timestamp, '0000000000000001', 1, 1, 2, 20
union all select 2, '2019-01-01 06:00:00'::timestamp, '0000000000000002', 1, 1, 2, 90
union all select 3, '2019-01-01 12:00:00'::timestamp, '0000000000000001', 1, 1, 2, 95
union all select 4, '2019-01-01 12:00:00'::timestamp, '0000000000000002', 1, 1, 2, 40
;
insert into rules(id, created_at, customer_id, device_type_id, device_sn, section)
select 			 1, '2019-01-01 00:00:00'::timestamp, 1, 1, '0000000000000001', '[50,60,80]'
union all select 2, '2019-01-01 00:00:00'::timestamp, 1, 1, '0000000000000002', '[50,60,80]'
;
insert into notifications(id, created_at, updated_at, device_sn, msg_id, rule_id, group_id, customer_id, device_type_id, key, content, solved)
select 			 1, '2019-01-01 06:00:00'::timestamp, '2019-01-01 06:00:00'::timestamp, '0000000000000001', 2, 1, 1, 1, 1, 'trashcan_usage', '垃圾箱使用率超过 80 %', false
union all select 2, '2019-01-01 12:00:00'::timestamp, '2019-01-01 06:00:00'::timestamp, '0000000000000002', 3, 2, 1, 1, 1, 'trashcan_usage', '垃圾箱使用率超过 80 %', false
;
`

var data = gofixtures.Data(gofixtures.Sql(sqlStr, []string{"device_types", "device_models", "groups", "devices", "trashcans", "notifications", "rules"}))

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
	ts, err = app.NewTrashcanService(*logger, thedb, trashcan.TrashcanConfig{})
	if err != nil {
		panic(err)
	}
}

func TestGetDevices(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.TrashcanDeviceList{
		List: []*protos.TrashcanDevice{
			{
				ID:            2,
				Name:          "垃圾箱测试设备02",
				SerialNumber:  "0000000000000002",
				Position:      "设备地址01",
				Group:         "垃圾箱分组01",
				GroupID:       1,
				DeviceType:    "垃圾箱传感器",
				DeviceTypeID:  1,
				DeviceModel:   "垃圾箱设备型号01",
				DeviceModelID: 1,
				CustomerID:    1,
				Latitude:      30,
				Longitude:     40,
				Percent:       68,
			},
			{
				ID:            1,
				Name:          "垃圾箱测试设备01",
				SerialNumber:  "0000000000000001",
				Position:      "设备地址01",
				Group:         "垃圾箱分组01",
				GroupID:       1,
				DeviceType:    "垃圾箱传感器",
				DeviceTypeID:  1,
				DeviceModel:   "垃圾箱设备型号01",
				DeviceModelID: 1,
				CustomerID:    1,
				Latitude:      10,
				Longitude:     20,
				Percent:       24,
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

	actual, err := ts.GetDevices(req)
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

	expected := &protos.TrashcanDevice{
		ID:            1,
		Name:          "垃圾箱测试设备01",
		SerialNumber:  "0000000000000001",
		Position:      "设备地址01",
		Group:         "垃圾箱分组01",
		GroupID:       1,
		DeviceType:    "垃圾箱传感器",
		DeviceTypeID:  1,
		DeviceModel:   "垃圾箱设备型号01",
		DeviceModelID: 1,
		CustomerID:    1,
		Latitude:      10,
		Longitude:     20,
		Percent:       24,
	}
	var customerID uint64 = 1
	deviceSN := "0000000000000001"
	actual, err := ts.GetDevice(customerID, deviceSN)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetDeviceGroups(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.GetTrashCanDeviceGroupsResponse{
		List: []*protos.TrashCanDeviceGroup{
			{
				ID:          1,
				Name:        "垃圾箱分组01",
				CustomerID:  1,
				DeviceTotal: 2,
				DeviceOff:   2,
			},
		},
		Count:       1,
		CurrentPage: 1,
		PageSize:    10,
		TotalCount:  1,
	}
	req := &protos.GetTrashCanDeviceGroupsRequest{
		Pagination: &protos.Pagination{
			Page:     1,
			PageSize: 10,
		},
		CustomerID: 1,
	}
	actual, err := ts.GetDeviceGroups(req)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetDeviceStatus(t *testing.T) {
	data.TruncatePut(thedb)
	var customerID uint64 = 1
	deviceSN := "0000000000000001"

	////////////////////////////////////////////////////////// 离线
	expected1 := 2

	actual1, err := ts.GetDeviceStatus(customerID, deviceSN)
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
	actual2, err := ts.GetDeviceStatus(customerID, deviceSN)
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

	actual3, err := ts.GetDeviceStatus(customerID, deviceSN)
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

	expected := &protos.TrashcanAlarmThresholdList{
		List: []*protos.TrashcanDeviceWithAlarmThreshold{
			{
				ID:           2,
				Name:         "垃圾箱测试设备02",
				Position:     "设备地址01",
				SerialNumber: "0000000000000002",
				GroupID:      1,
				Group:        "垃圾箱分组01",
				CreatedAt:    "2019-01-01 00:00:00",
				StageInfo:    50,
				StageWarn:    60,
				StageAlert:   80,
			},
			{
				ID:           1,
				Name:         "垃圾箱测试设备01",
				Position:     "设备地址01",
				SerialNumber: "0000000000000001",
				GroupID:      1,
				Group:        "垃圾箱分组01",
				CreatedAt:    "2019-01-01 00:00:00",
				StageInfo:    50,
				StageWarn:    60,
				StageAlert:   80,
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

	actual, err := ts.GetDeviceAlarmThresholds(req)
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
	alarmThresholds := []float64{10, 20, 30}

	req1 := &protos.SetAlarmThresholdRequest{
		DeviceSN:        deviceSN,
		CustomerID:      customerID,
		AlarmThresholds: alarmThresholds,
	}
	err := ts.SetDeviceAlarmThreshold(req1)
	if err != nil {
		t.Error(err)
	}

	expected := &protos.TrashcanAlarmThresholdList{
		List: []*protos.TrashcanDeviceWithAlarmThreshold{
			{
				ID:           2,
				Name:         "垃圾箱测试设备02",
				Position:     "设备地址01",
				SerialNumber: "0000000000000002",
				GroupID:      1,
				Group:        "垃圾箱分组01",
				CreatedAt:    "2019-01-01 00:00:00",
				StageInfo:    50,
				StageWarn:    60,
				StageAlert:   80,
			},
			{
				ID:           1,
				Name:         "垃圾箱测试设备01",
				Position:     "设备地址01",
				SerialNumber: "0000000000000001",
				GroupID:      1,
				Group:        "垃圾箱分组01",
				CreatedAt:    "2019-01-01 00:00:00",
				StageInfo:    10,
				StageWarn:    20,
				StageAlert:   30,
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

	actual, err := ts.GetDeviceAlarmThresholds(req2)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}
