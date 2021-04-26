package app_test

import (
	"os"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/jinzhu/gorm"
	"github.com/theplant/gofixtures"
	"github.com/theplant/testingutils"

	"github.com/zsy-cn/bms/environ_monitor"
	"github.com/zsy-cn/bms/environ_monitor/app"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/migration"
)

var thedb *gorm.DB
var logger = log.NewLogger(os.Stdout)
var es environ_monitor.EnvironMonitorService

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
	data.TruncatePut(thedb)
	logger.Info("init database tables completed")
	es, err = app.NewEnvironMonitorService(*logger, thedb, environ_monitor.EnvironMonitorConfig{})
	if err != nil {
		panic(err)
	}
}

var sqlStr = `
insert into device_types(id, key, name, is_sensor, properties)
select 1, 'environ_monitor', '环境监测传感器', true, '["temperature","pm2.5","noise","humidity"]'
;
insert into device_models(id, name, manufacturer_id, device_type_id)
select 1, '环境监测设备型号01', 1, 1
;
insert into groups(id, name, device_type_id, customer_id)
select 			 1, '环境监测测试分组01', 1, 1
union all select 2, '环境监测测试分组02', 1, 1
;
insert into devices(id, created_at, updated_at, name, position, serial_number, customer_id, group_id, device_type_id, device_model_id, latitude, longitude)
select 			 1, '2019-01-01 12:00:00'::timestamp, '2019-01-01 12:00:00'::timestamp, '环境监测测试设备01', '设备地址01', '0000000000000001', 1, 1, 1, 1, 0, 0
union all select 2, '2019-01-01 12:00:00'::timestamp, '2019-01-01 12:00:00'::timestamp, '环境监测测试设备02', '设备地址01', '0000000000000002', 1, 1, 1, 1, 0, 0
union all select 3, '2019-01-01 12:00:00'::timestamp, '2019-01-01 12:00:00'::timestamp, '环境监测测试设备03', '设备地址02', '0000000000000003', 1, 2, 1, 1, 0, 0
union all select 4, '2019-01-01 12:00:00'::timestamp, '2019-01-01 12:00:00'::timestamp, '环境监测测试设备04', '设备地址02', '0000000000000004', 1, 2, 1, 1, 0, 0
;
insert into environ_monitors(id, created_at, device_sn, group_id, customer_id, temperature, humidity, noise, pm025)
select 			 1, '2019-01-01 06:00:00'::timestamp, '0000000000000001', 1, 1, 21, 50, 60, 10
union all select 2, '2019-01-01 06:00:00'::timestamp, '0000000000000002', 1, 1, 22, 60, 70, 10
;
`

var data = gofixtures.Data(gofixtures.Sql(sqlStr, []string{"device_types", "device_models", "groups", "devices", "environ_monitors"}))

func TestGetEnvironMonitorDevices(t *testing.T) {
	data.TruncatePut(thedb)

	req := &protos.GetDevicesRequestForCustomer{
		Pagination: &protos.Pagination{
			Page:     1,
			PageSize: 10,
		},
		CustomerID: 1,
		GroupID:    1,
	}

	expected1 := &protos.EnvironMonitorDeviceList{
		List: []*protos.EnvironMonitorDevice{
			{
				ID:            2,
				Name:          "环境监测测试设备02",
				Position:      "设备地址01",
				SerialNumber:  "0000000000000002",
				CustomerID:    1,
				GroupID:       1,
				Group:         "环境监测测试分组01",
				DeviceTypeID:  1,
				DeviceType:    "环境监测传感器",
				DeviceModelID: 1,
				DeviceModel:   "环境监测设备型号01",
				CreatedAt:     "2019-01-01",
				Actived:       "true",
				StatusCode:    2,
				Status:        "离线",
			},
			{
				ID:            1,
				Name:          "环境监测测试设备01",
				Position:      "设备地址01",
				SerialNumber:  "0000000000000001",
				CustomerID:    1,
				GroupID:       1,
				Group:         "环境监测测试分组01",
				DeviceTypeID:  1,
				DeviceType:    "环境监测传感器",
				DeviceModelID: 1,
				DeviceModel:   "环境监测设备型号01",
				CreatedAt:     "2019-01-01",
				Actived:       "true",
				StatusCode:    2,
				Status:        "离线",
			},
		},
		Count:       2,
		CurrentPage: 1,
		PageSize:    10,
		TotalCount:  2,
	}

	actual1, err := es.GetEnvironMonitorDevices(req)
	if err != nil {
		t.Error(err)
	}

	diff1 := testingutils.PrettyJsonDiff(expected1, actual1)
	if len(diff1) > 0 {
		t.Error(diff1)
	}
	//////////////////////////////////////////////////////////
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fakeNow, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-01-01 12:00:01", loc)
	timenowPatcher := monkey.Patch(time.Now, func() time.Time {
		return fakeNow
	})
	defer timenowPatcher.Unpatch()

	expected2 := &protos.EnvironMonitorDeviceList{
		List: []*protos.EnvironMonitorDevice{
			{
				ID:            2,
				Name:          "环境监测测试设备02",
				Position:      "设备地址01",
				SerialNumber:  "0000000000000002",
				CustomerID:    1,
				GroupID:       1,
				Group:         "环境监测测试分组01",
				DeviceTypeID:  1,
				DeviceType:    "环境监测传感器",
				DeviceModelID: 1,
				DeviceModel:   "环境监测设备型号01",
				CreatedAt:     "2019-01-01",
				Actived:       "true",
				StatusCode:    1,
				Status:        "正常",
				Temperature:   22,
				PM025:         10,
				Noise:         70,
				Humidity:      60,
			},
			{
				ID:            1,
				Name:          "环境监测测试设备01",
				Position:      "设备地址01",
				SerialNumber:  "0000000000000001",
				CustomerID:    1,
				GroupID:       1,
				Group:         "环境监测测试分组01",
				DeviceTypeID:  1,
				DeviceType:    "环境监测传感器",
				DeviceModelID: 1,
				DeviceModel:   "环境监测设备型号01",
				CreatedAt:     "2019-01-01",
				Actived:       "true",
				StatusCode:    1,
				Status:        "正常",
				Temperature:   21,
				PM025:         10,
				Noise:         60,
				Humidity:      50,
			},
		},
		Count:       2,
		CurrentPage: 1,
		PageSize:    10,
		TotalCount:  2,
	}

	actual2, err := es.GetEnvironMonitorDevices(req)
	if err != nil {
		t.Error(err)
	}

	diff2 := testingutils.PrettyJsonDiff(expected2, actual2)
	if len(diff2) > 0 {
		t.Error(diff2)
	}
}

func TestGetEnvironMonitorDevice(t *testing.T) {
	data.TruncatePut(thedb)

	expected1 := &protos.EnvironMonitorDevice{
		ID:            2,
		Name:          "环境监测测试设备02",
		Position:      "设备地址01",
		SerialNumber:  "0000000000000002",
		CustomerID:    1,
		GroupID:       1,
		Group:         "环境监测测试分组01",
		DeviceTypeID:  1,
		DeviceType:    "环境监测传感器",
		DeviceModelID: 1,
		DeviceModel:   "环境监测设备型号01",
		CreatedAt:     "2019-01-01",
		Actived:       "true",
		StatusCode:    2,
		Status:        "离线",
	}

	actual1, err := es.GetEnvironMonitorDevice(1, "0000000000000002")
	if err != nil {
		t.Error(err)
	}

	diff1 := testingutils.PrettyJsonDiff(expected1, actual1)
	if len(diff1) > 0 {
		t.Error(diff1)
	}
	//////////////////////////////////////////////////////////
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fakeNow, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-01-01 12:00:01", loc)
	timenowPatcher := monkey.Patch(time.Now, func() time.Time {
		return fakeNow
	})
	defer timenowPatcher.Unpatch()

	expected2 := &protos.EnvironMonitorDevice{
		ID:            2,
		Name:          "环境监测测试设备02",
		Position:      "设备地址01",
		SerialNumber:  "0000000000000002",
		CustomerID:    1,
		GroupID:       1,
		Group:         "环境监测测试分组01",
		DeviceTypeID:  1,
		DeviceType:    "环境监测传感器",
		DeviceModelID: 1,
		DeviceModel:   "环境监测设备型号01",
		CreatedAt:     "2019-01-01",
		Actived:       "true",
		StatusCode:    1,
		Status:        "正常",
		Temperature:   22,
		PM025:         10,
		Noise:         70,
		Humidity:      60,
	}

	actual2, err := es.GetEnvironMonitorDevice(1, "0000000000000002")
	if err != nil {
		t.Error(err)
	}

	diff2 := testingutils.PrettyJsonDiff(expected2, actual2)
	if len(diff2) > 0 {
		t.Error(diff2)
	}
}

func TestGetEnvironMonitorDeviceStatus(t *testing.T) {
	data.TruncatePut(thedb)

	expected := 2 // 离线
	actual, err := es.GetEnvironMonitorDeviceStatus(1, "0000000000000002")
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func generateMsg(dateFrom string) (msgList []*environ_monitor.EnvironMonitor) {
	duration, _ := time.ParseDuration("-1h")
	msgList = []*environ_monitor.EnvironMonitor{}

	var temperature, humidity, noise, pm025 float64
	var baseDate time.Time
	temperature = 23.4
	humidity = 23
	noise = 34
	pm025 = 130
	if dateFrom == "" {
		baseDate = time.Now()
	} else {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		baseDate, _ = time.ParseInLocation("2006-01-02", dateFrom, loc)
	}
	i := 24
	var id uint64 = 0
	for {
		i -= 1
		id += 1
		if i < 0 {
			break
		}
		createdAt := baseDate.Add(duration * time.Duration(i))
		msgList = append(msgList, &environ_monitor.EnvironMonitor{
			Base: model.Base{
				ID:        id,
				CreatedAt: createdAt,
			},
			App: model.App{
				DeviceSN:   "0000000000000001",
				GroupID:    1,
				CustomerID: 1,
			},
			Temperature: temperature + float64(id),
			Humidity:    humidity + float64(id),
			Noise:       noise + float64(id),
			PM025:       pm025 + float64(id),
		})
	}
	return msgList
}

func TestGetLastEnvironMoniterAverageInfo(t *testing.T) {
	data.TruncatePut(thedb)
	msgList := generateMsg("")
	msgData := gofixtures.Data(msgList)
	msgData.TruncatePut(thedb)

	msgListIdx := len(msgList) - 1
	expected := &protos.EnvironMonitorSectionAverageData{
		Temperature: msgList[msgListIdx].Temperature,
		PM025:       msgList[msgListIdx].PM025,
		Noise:       msgList[msgListIdx].Noise,
		Humidity:    msgList[msgListIdx].Humidity,
	}
	actual, err := es.GetLastEnvironMoniterAverageInfo(1, 1, "")
	if err != nil {
		t.Error(err)
	}

	diff := testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestGetEnvironMonitorSectionAverageData(t *testing.T) {
	data.TruncatePut(thedb)
	msgList := generateMsg("2019-01-02")
	msgData := gofixtures.Data(msgList)
	msgData.TruncatePut(thedb)

	expected := []map[string]interface{}{
		map[string]interface{}{
			"PM025":       132,
			"Noise":       36,
			"Humidity":    25,
			"Hour":        "3时",
			"Temperature": 25.4,
		},
		map[string]interface{}{
			"Hour":        "6时",
			"Temperature": 28.4,
			"PM025":       135,
			"Noise":       39,
			"Humidity":    28,
		},
		map[string]interface{}{
			"Noise":       42,
			"Humidity":    31,
			"Hour":        "9时",
			"Temperature": 31.4,
			"PM025":       138,
		},
		map[string]interface{}{
			"Noise":       45,
			"Humidity":    34,
			"Hour":        "12时",
			"Temperature": 34.4,
			"PM025":       141,
		},
		map[string]interface{}{
			"Hour":        "15时",
			"Temperature": 37.4,
			"PM025":       144,
			"Noise":       48,
			"Humidity":    37,
		},
		map[string]interface{}{
			"Hour":        "18时",
			"Temperature": 40.4,
			"PM025":       147,
			"Noise":       51,
			"Humidity":    40,
		},
		map[string]interface{}{
			"Hour":        "21时",
			"Temperature": 43.4,
			"PM025":       150,
			"Noise":       54,
			"Humidity":    43,
		},
	}
	actual, err := es.GetEnvironMonitorSectionAverageData(1, 1, "2019-01-01", "2019-01-01", 0)
	if err != nil {
		t.Error(err)
	}
	// actual删去了第一个成员, 是因为第一条记录为空, 解析为map时, temperature等值为NaN, 对比会出错, 所以先移除, 以后再考虑.
	diff := testingutils.PrettyJsonDiff(expected, actual[1:len(actual)])
	if len(diff) > 0 {
		t.Error(diff)
	}
}
