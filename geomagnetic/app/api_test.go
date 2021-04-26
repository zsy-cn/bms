package app_test

import (
	"os"
	"testing"
	"time"

	"bou.ke/monkey"

	"github.com/jinzhu/gorm"
	"github.com/theplant/gofixtures"
	"github.com/theplant/testingutils"
	"github.com/zsy-cn/bms/geomagnetic"
	"github.com/zsy-cn/bms/geomagnetic/app"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/migration"
)

var thedb *gorm.DB
var logger = log.NewLogger(os.Stdout)
var gs geomagnetic.GeomagneticService

var sqlStr = `
insert into device_types(id, key, name, is_sensor, properties)
select 1, 'geomagnetic', '地磁传感器', true, '["value"]'
;
insert into device_models(id, name, manufacturer_id, device_type_id)
select 1, '地磁设备型号01', 1, 1
;
insert into groups(id, name, device_type_id, customer_id)
select 1, '停车场01', 1, 1
;
insert into devices(id, created_at, name, position, serial_number, customer_id, group_id, device_type_id, device_model_id, latitude, longitude)
select 			 1, '2019-01-01 12:00:00'::timestamp, '地磁测试设备01', '设备地址01', '0000000000000001', 1, 1, 1, 1, 10, 20
union all select 2, '2019-01-01 12:00:00'::timestamp, '地磁测试设备02', '设备地址01', '0000000000000002', 1, 1, 1, 1, 30, 40
;
insert into geomagnetics(id, created_at, device_sn, group_id, customer_id, value)
select 			 1, '2019-01-01 12:00:00'::timestamp, '0000000000000001', 1, 1, '1'
union all select 2, '2019-01-01 12:00:00'::timestamp, '0000000000000002', 1, 1, '0'
;
`

var msgSQLStr = `
insert into geomagnetics(id, created_at, device_sn, group_id, customer_id, value)
select 			 1, '2019-01-01 01:00:00'::timestamp, '0000000000000001', 1, 1, '1'
union all select 2, '2019-01-01 02:00:00'::timestamp, '0000000000000001', 1, 1, '1'
union all select 3, '2019-01-01 03:00:00'::timestamp, '0000000000000001', 1, 1, '1'
union all select 4, '2019-01-01 04:00:00'::timestamp, '0000000000000001', 1, 1, '1'
union all select 5, '2019-01-01 05:00:00'::timestamp, '0000000000000001', 1, 1, '1'
union all select 6, '2019-01-01 06:00:00'::timestamp, '0000000000000001', 1, 1, '1'
union all select 7, '2019-01-01 07:00:00'::timestamp, '0000000000000001', 1, 1, '0'
union all select 8, '2019-01-01 08:00:00'::timestamp, '0000000000000001', 1, 1, '0'
union all select 9, '2019-01-01 09:00:00'::timestamp, '0000000000000001', 1, 1, '0'
union all select 10, '2019-01-01 10:00:00'::timestamp, '0000000000000001', 1, 1, '0'
union all select 11, '2019-01-01 11:00:00'::timestamp, '0000000000000001', 1, 1, '0'
union all select 12, '2019-01-01 12:00:00'::timestamp, '0000000000000001', 1, 1, '0'
union all select 13, '2019-01-01 13:00:00'::timestamp, '0000000000000001', 1, 1, '0'
union all select 14, '2019-01-01 14:00:00'::timestamp, '0000000000000001', 1, 1, '1'
union all select 15, '2019-01-01 15:00:00'::timestamp, '0000000000000001', 1, 1, '0'
union all select 16, '2019-01-01 16:00:00'::timestamp, '0000000000000001', 1, 1, '1'
union all select 17, '2019-01-01 17:00:00'::timestamp, '0000000000000001', 1, 1, '0'
union all select 18, '2019-01-01 18:00:00'::timestamp, '0000000000000001', 1, 1, '1'
union all select 19, '2019-01-01 19:00:00'::timestamp, '0000000000000001', 1, 1, '0'
union all select 20, '2019-01-01 20:00:00'::timestamp, '0000000000000001', 1, 1, '1'
union all select 21, '2019-01-01 21:00:00'::timestamp, '0000000000000001', 1, 1, '0'
union all select 22, '2019-01-01 22:00:00'::timestamp, '0000000000000001', 1, 1, '1'
union all select 23, '2019-01-01 23:00:00'::timestamp, '0000000000000001', 1, 1, '0'
union all select 24, '2019-01-02 00:00:00'::timestamp, '0000000000000001', 1, 1, '1'

union all select 25, '2019-01-01 01:00:00'::timestamp, '0000000000000002', 1, 1, '1'
union all select 26, '2019-01-01 12:00:00'::timestamp, '0000000000000002', 1, 1, '0'
;
`

var data = gofixtures.Data(gofixtures.Sql(sqlStr, []string{"device_types", "device_models", "groups", "devices", "geomagnetics"}))
var msgdData = gofixtures.Data(gofixtures.Sql(msgSQLStr, []string{"geomagnetics"}))

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
	gs, err = app.NewGeomagneticService(*logger, thedb, geomagnetic.GeomagneticConfig{})
	if err != nil {
		panic(err)
	}
}

func TestGetGeomagneticDevices(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.GeomagneticDeviceList{
		List: []*protos.GeomagneticDevice{
			{
				ID:            2,
				Name:          "地磁测试设备02",
				SerialNumber:  "0000000000000002",
				Position:      "设备地址01",
				Group:         "停车场01",
				GroupID:       1,
				DeviceType:    "地磁传感器",
				DeviceTypeID:  1,
				DeviceModel:   "地磁设备型号01",
				DeviceModelID: 1,
				CustomerID:    1,
				Latitude:      30,
				Longitude:     40,
				StatusCode:    2,
				Status:        "离线",
			},
			{
				ID:            1,
				Name:          "地磁测试设备01",
				SerialNumber:  "0000000000000001",
				Position:      "设备地址01",
				Group:         "停车场01",
				GroupID:       1,
				DeviceType:    "地磁传感器",
				DeviceTypeID:  1,
				DeviceModel:   "地磁设备型号01",
				DeviceModelID: 1,
				CustomerID:    1,
				Latitude:      10,
				Longitude:     20,
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

	actual, err := gs.GetGeomagneticDevices(req)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetGeomagneticDevice(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.GeomagneticDevice{
		ID:            1,
		Name:          "地磁测试设备01",
		SerialNumber:  "0000000000000001",
		Position:      "设备地址01",
		Group:         "停车场01",
		GroupID:       1,
		DeviceType:    "地磁传感器",
		DeviceTypeID:  1,
		DeviceModel:   "地磁设备型号01",
		DeviceModelID: 1,
		CustomerID:    1,
		Latitude:      10,
		Longitude:     20,
		StatusCode:    2,
		Status:        "离线",
	}
	var customerID uint64 = 1
	deviceSN := "0000000000000001"
	actual, err := gs.GetGeomagneticDevice(customerID, deviceSN)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetParkingInfos(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.ParkingPlaceList{
		List: []*protos.ParkingPlace{
			{
				ID:        1,
				Name:      "停车场01",
				Amount:    2,
				Used:      1,
				Unused:    1,
				Longitude: 20,
				Latitude:  10,
			},
		},
		Count:       1,
		CurrentPage: 1,
		PageSize:    10,
		TotalCount:  1,
	}

	req := &protos.GetParkingPlacesRequest{
		Pagination: &protos.Pagination{
			Page:     1,
			PageSize: 10,
		},
		CustomerID: 1,
	}

	actual, err := gs.GetParkingInfos(req)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetParkingInfo(t *testing.T) {
	data.TruncatePut(thedb)

	expected := &protos.ParkingPlace{
		ID:        1,
		Name:      "停车场01",
		Amount:    2,
		Used:      1,
		Unused:    1,
		Longitude: 20,
		Latitude:  10,
	}
	var customerID, groupID uint64 = 1, 1
	actual, err := gs.GetParkingInfo(customerID, groupID)
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetParkingHistory(t *testing.T) {
	data.TruncatePut(thedb)
	msgdData.TruncatePut(thedb)
	//////////////////////////////////////////////////////////
	expected1 := []map[string]interface{}{
		map[string]interface{}{
			"hour":    "1时",
			"percent": 0,
		},
		map[string]interface{}{
			"hour":    "3时",
			"percent": 1,
		},
		map[string]interface{}{
			"hour":    "5时",
			"percent": 1,
		},
		map[string]interface{}{
			"hour":    "7时",
			"percent": 1,
		},
		map[string]interface{}{
			"hour":    "9时",
			"percent": 0.5,
		},
		map[string]interface{}{
			"hour":    "11时",
			"percent": 0.5,
		},
		map[string]interface{}{
			"hour":    "13时",
			"percent": 0,
		},
		map[string]interface{}{
			"hour":    "15时",
			"percent": 0.5,
		},
		map[string]interface{}{
			"hour":    "17时",
			"percent": 0.5,
		},
		map[string]interface{}{
			"hour":    "19时",
			"percent": 0.5,
		},
		map[string]interface{}{
			"hour":    "21时",
			"percent": 0.5,
		},
		map[string]interface{}{
			"hour":    "23时",
			"percent": 0.5,
		},
	}

	var customerID, groupID uint64 = 1, 1
	actual1, err := gs.GetParkingHistory(customerID, groupID, "2019-01-01")
	if err != nil {
		t.Error(err)
	}

	diff1 := testingutils.PrettyJsonDiff(expected1, actual1)
	if len(diff1) > 0 {
		t.Error(diff1)
	}
	//////////////////////////////////////////////////////////
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fakeNow, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-01-01 12:00:00", loc)
	timenowPatcher := monkey.Patch(time.Now, func() time.Time {
		return fakeNow
	})
	defer timenowPatcher.Unpatch()
	expected2 := []map[string]interface{}{
		map[string]interface{}{
			"hour":    "1时",
			"percent": 0,
		},
		map[string]interface{}{
			"hour":    "3时",
			"percent": 1,
		},
		map[string]interface{}{
			"hour":    "5时",
			"percent": 1,
		},
		map[string]interface{}{
			"hour":    "7时",
			"percent": 1,
		},
		map[string]interface{}{
			"hour":    "9时",
			"percent": 0.5,
		},
		map[string]interface{}{
			"hour":    "11时",
			"percent": 0.5,
		},
		map[string]interface{}{
			"hour":    "13时",
			"percent": 0,
		},
		map[string]interface{}{
			"hour":    "15时",
			"percent": 0,
		},
		map[string]interface{}{
			"hour":    "17时",
			"percent": 0,
		},
		map[string]interface{}{
			"hour":    "19时",
			"percent": 0,
		},
		map[string]interface{}{
			"hour":    "21时",
			"percent": 0,
		},
		map[string]interface{}{
			"hour":    "23时",
			"percent": 0,
		},
	}
	actual2, err := gs.GetParkingHistory(customerID, groupID, "")
	if err != nil {
		t.Error(err)
	}

	diff2 := testingutils.PrettyJsonDiff(expected2, actual2)
	if len(diff2) > 0 {
		t.Error(diff2)
	}
}
