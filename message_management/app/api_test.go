package app_test

import (
	"os"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/jinzhu/gorm"

	"github.com/theplant/gofixtures"
	"github.com/theplant/testingutils"
	"github.com/zsy-cn/bms/message_management"
	"github.com/zsy-cn/bms/message_management/app"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/migration"
)

var thedb *gorm.DB
var logger = log.NewLogger(os.Stdout)
var ms message_management.MessageManagementService

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
	ms, err = app.NewMessageManagementService(*logger, thedb, message_management.MessageManagementConfig{})
	if err != nil {
		panic(err)
	}
}

func TestNotifications(t *testing.T) {
	data.TruncatePut(thedb)
	//////////////////////////////// GetNotifications
	expected1 := &protos.NotificationList{
		List: []*protos.NotificationInfo{
			{
				ID:                 2,
				DeviceName:         "井盖测试设备01",
				DeviceSerialNumber: "0000000000000001",
				DeviceType:         "井盖传感器",
				DeviceTypeID:       1,
				Group:              "井盖分组01",
				GroupID:            1,
				CreatedAt:          "2019-01-01 12:00:00",
				Content:            "井盖报警",
				Solved:             "false",
				SolvedAt:           "0001-01-01 08:00:00",
			},
			{
				ID:                 1,
				DeviceName:         "井盖测试设备02",
				DeviceSerialNumber: "0000000000000002",
				DeviceType:         "井盖传感器",
				DeviceTypeID:       1,
				Group:              "井盖分组01",
				GroupID:            1,
				CreatedAt:          "2019-01-01 06:00:00",
				Content:            "井盖报警",
				Solved:             "false",
				SolvedAt:           "0001-01-01 08:00:00",
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

	actual1, err := ms.GetNotifications(req)
	if err != nil {
		t.Error(err)
	}

	diff1 := testingutils.PrettyJsonDiff(expected1, actual1)
	if len(diff1) > 0 {
		t.Error(diff1)
	}

	//////////////////////////////// DiscardNotification
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fakeTimeStr := "2019-01-01 12:00:01"
	fakeNow, _ := time.ParseInLocation("2006-01-02 15:04:05", fakeTimeStr, loc)
	timenowPatcher := monkey.Patch(time.Now, func() time.Time {
		return fakeNow
	})
	defer timenowPatcher.Unpatch()

	var customerID uint64 = 1
	deviceSN := "0000000000000001"
	err = ms.DiscardNotification(customerID, deviceSN)
	if err != nil {
		t.Error(err)
	}
	expected2 := &protos.NotificationList{
		List: []*protos.NotificationInfo{
			{
				ID:                 2,
				DeviceName:         "井盖测试设备01",
				DeviceSerialNumber: "0000000000000001",
				DeviceType:         "井盖传感器",
				DeviceTypeID:       1,
				Group:              "井盖分组01",
				GroupID:            1,
				CreatedAt:          "2019-01-01 12:00:00",
				Content:            "井盖报警",
				Solved:             "true",
				SolvedAt:           fakeTimeStr,
			},
			{
				ID:                 1,
				DeviceName:         "井盖测试设备02",
				DeviceSerialNumber: "0000000000000002",
				DeviceType:         "井盖传感器",
				DeviceTypeID:       1,
				Group:              "井盖分组01",
				GroupID:            1,
				CreatedAt:          "2019-01-01 06:00:00",
				Content:            "井盖报警",
				Solved:             "false",
				SolvedAt:           "0001-01-01 08:00:00",
			},
		},
		Count:       2,
		CurrentPage: 1,
		PageSize:    10,
		TotalCount:  2,
	}

	actual2, err := ms.GetNotifications(req)
	if err != nil {
		t.Error(err)
	}

	diff2 := testingutils.PrettyJsonDiff(expected2, actual2)
	if len(diff2) > 0 {
		t.Error(diff2)
	}

}
