package app_test

import (
	"os"
	"testing"
	"time"

	"bou.ke/monkey"

	"github.com/jinzhu/gorm"
	"github.com/theplant/gofixtures"
	"github.com/theplant/testingutils"
	"github.com/zsy-cn/bms/manhole_cover"
	"github.com/zsy-cn/bms/manhole_cover/app"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/migration"
)

var thedb *gorm.DB
var logger = log.NewLogger(os.Stdout)
var ms manhole_cover.ManholeCoverService

var sqlStr = `
insert into device_types(id, key, name, is_sensor, properties)
select 1, 'manhole_cover', '井盖传感器', true, '["status"]'
;
insert into device_models(id, name, manufacturer_id, device_type_id)
select 1, '井盖设备型号01', 1, 1
;
insert into groups(id, name, device_type_id, customer_id)
select 1, '井盖分组01', 1, 1
;
insert into devices(id, created_at, name, position, serial_number, customer_id, group_id, device_type_id, device_model_id, latitude, longitude)
select 			 1, '2019-01-01 12:00:00'::timestamp, '井盖测试设备01', '设备地址01', '0000000000000001', 1, 1, 1, 1, 10, 20
union all select 2, '2019-01-01 12:00:00'::timestamp, '井盖测试设备02', '设备地址01', '0000000000000002', 1, 1, 1, 1, 30, 40
;
insert into manhole_covers(id, created_at, device_sn, group_id, customer_id, status)
select 			 1, '2019-01-01 06:00:00'::timestamp, '0000000000000001', 1, 1, 0
union all select 2, '2019-01-01 06:00:00'::timestamp, '0000000000000002', 1, 1, 1
union all select 3, '2019-01-01 12:00:00'::timestamp, '0000000000000001', 1, 1, 1
union all select 4, '2019-01-01 12:00:00'::timestamp, '0000000000000002', 1, 1, 0
;
insert into notifications(id, created_at, updated_at, device_sn, msg_id, group_id, customer_id, device_type_id, key, content, solved)
select 			 1, '2019-01-01 06:00:00'::timestamp, '2019-01-01 06:00:00'::timestamp, '0000000000000002', 2, 1, 1, 1, 'manhole_cover_open_alert', '井盖报警', false
union all select 2, '2019-01-01 12:00:00'::timestamp, '2019-01-01 06:00:00'::timestamp, '0000000000000001', 3, 1, 1, 1, 'manhole_cover_open_alert', '井盖报警', false
`

var data = gofixtures.Data(gofixtures.Sql(sqlStr, []string{"device_types", "device_models", "groups", "devices", "manhole_covers", "notifications"}))

func init() {
	var err error
	connectStr := "host=localhost port=1234 user=backend dbname=backend sslmode=disable password=backend"
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
	ms, err = app.NewManholeCoverService(*logger, thedb, manhole_cover.ManholeCoverConfig{})
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
				Name:          "井盖测试设备02",
				SerialNumber:  "0000000000000002",
				Position:      "设备地址01",
				Group:         "井盖分组01",
				GroupID:       1,
				DeviceType:    "井盖传感器",
				DeviceTypeID:  1,
				DeviceModel:   "井盖设备型号01",
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
				Name:          "井盖测试设备01",
				SerialNumber:  "0000000000000001",
				Position:      "设备地址01",
				Group:         "井盖分组01",
				GroupID:       1,
				DeviceType:    "井盖传感器",
				DeviceTypeID:  1,
				DeviceModel:   "井盖设备型号01",
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

	actual, err := ms.GetDevices(req)
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
		Name:          "井盖测试设备01",
		SerialNumber:  "0000000000000001",
		Position:      "设备地址01",
		Group:         "井盖分组01",
		GroupID:       1,
		DeviceType:    "井盖传感器",
		DeviceTypeID:  1,
		DeviceModel:   "井盖设备型号01",
		DeviceModelID: 1,
		CustomerID:    1,
		Latitude:      10,
		Longitude:     20,
		Actived:       "true",
		CreatedAt:     "2019-01-01",
		StatusCode:    2,
		Status:        "离线",
	}

	actual, err := ms.GetDevice(1, "0000000000000001")
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
	actual1, err := ms.GetDeviceInfos(req)
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

	actual2, err := ms.GetDeviceInfos(req)
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

	actual1, err := ms.GetDeviceStatus(customerID, deviceSN)
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
	actual2, err := ms.GetDeviceStatus(customerID, deviceSN)
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

	actual3, err := ms.GetDeviceStatus(customerID, deviceSN)
	if err != nil {
		t.Error(err)
	}

	diff3 := testingutils.PrettyJsonDiff(expected3, actual3)
	if len(diff3) > 0 {
		t.Error(diff3)
	}
}
