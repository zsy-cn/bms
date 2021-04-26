package internal

import (
	"fmt"
	"math"
	"time"

	"github.com/zsy-cn/bms/environ_monitor"
	"github.com/zsy-cn/bms/protos"

	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/timesection"
)

// getLastEnvironMonitorInfo 获取指定设备的最近一次的有效消息记录.
func getLastEnvironMonitorInfo(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (msg *environ_monitor.EnvironMonitor, err error) {
	msg = &environ_monitor.EnvironMonitor{}
	now := time.Now()
	d2, _ := time.ParseDuration(offlineTimeout)
	deadline := now.Add(d2)
	whereArgs := "device_sn = ? and created_at > ?"
	err = db.Where(whereArgs, deviceSN, deadline).Order("id desc").First(msg).Error
	return
}

// getLastEnvironMoniterInfos 查询指定条件下的设备最近一次的环境记录
// @param: datetime 表示只查询指定时间点之前的记录, 如果为空, 则查询当前时间之前的环境设备记录
func getLastEnvironMoniterInfos(db *gorm.DB, log log.Logger, customerID uint64, groupID uint64, datetime string) (envRecords []*environ_monitor.EnvironMonitor, err error) {
	// 查询所有设备最近的消息, 条件查询语句在后面拼接
	sqlStr1 := "select device_sn, max(created_at) as created_at from environ_monitors where 1 = 1 %s group by device_sn "
	sqlStr2 := "select main_tbl.* from environ_monitors as main_tbl inner join (%s) as b on main_tbl.device_sn = b.device_sn and main_tbl.created_at = b.created_at"
	whereStr := ""
	if customerID != 0 {
		whereStr = fmt.Sprintf(whereStr+" and customer_id = %d ", customerID)
	}
	if groupID != 0 {
		whereStr = fmt.Sprintf(whereStr+" and group_id = %d ", groupID)
	}
	if datetime != "" {
		whereStr = fmt.Sprintf(whereStr+" and created_at < timestamp '%s' ", datetime)
	}
	_sqlStr1 := fmt.Sprintf(sqlStr1, whereStr)
	sqlStr := fmt.Sprintf(sqlStr2, _sqlStr1)
	log.Debugf("execute sql in GetEnvironMonitorSectionAverageData(): %s", sqlStr)

	envRecords = []*environ_monitor.EnvironMonitor{}
	err = db.Raw(sqlStr).Scan(&envRecords).Error
	if err != nil {
		log.Errorf("query environ monitor message failed in getLastEnvironMoniterInfos(): %s", err.Error())
		return
	}
	return
}

// GetEnvironMonitorDevices 获取环境监测设备列表
// TOTO: 获取设备列表同时查询各设备最近一次的信息情况, 在地图上作为数据点展示.
func GetEnvironMonitorDevices(db *gorm.DB, log log.Logger, req *protos.GetDevicesRequestForCustomer) (envDeviceList *protos.EnvironMonitorDeviceList, err error) {
	envDeviceList = &protos.EnvironMonitorDeviceList{
		List: []*protos.EnvironMonitorDevice{},
	}
	req.DeviceTypeID = deviceTypeID
	deviceList, err := _deviceManagementService.GetDevices(req)
	if err != nil {
		log.Errorf("get device list by device management service failed in GetEnvironMonitorDevices(): %s", err.Error())
		return
	}
	for _, device := range deviceList.List {
		envDevice := &protos.EnvironMonitorDevice{}
		err = deviceToEnvironMonitorDevice(db, log, device, envDevice)
		if err != nil {
			log.Errorf("trans device to environ monitor failed in deviceToEnvironMonitorDevice(): %s", err.Error())
			return
		}
		envDeviceList.List = append(envDeviceList.List, envDevice)
	}

	envDeviceList.Count = deviceList.Count
	envDeviceList.TotalCount = deviceList.Count
	envDeviceList.CurrentPage = req.Pagination.Page
	envDeviceList.PageSize = req.Pagination.PageSize
	return
}

// GetEnvironMonitorDevice ...
func GetEnvironMonitorDevice(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (envDevice *protos.EnvironMonitorDevice, err error) {
	device, err := _deviceManagementService.GetDevice(deviceSN, customerID, deviceTypeID)
	if err != nil {
		return
	}
	envDevice = &protos.EnvironMonitorDevice{}
	err = deviceToEnvironMonitorDevice(db, log, device, envDevice)
	if err != nil {
		return
	}
	return
}

// GetEnvironMonitorDeviceStatus ...
// @return: status -1: 设备不存在, 1: 在线, 2: 离线, 3: 其他
func GetEnvironMonitorDeviceStatus(db *gorm.DB, log log.Logger, customerID uint64, deviceSN string) (status uint64, err error) {
	_, err = getLastEnvironMonitorInfo(db, log, customerID, deviceSN)
	if err != nil {
		err = nil
		status = 2 // 离线
		return
	}
	status = 1
	return
}

// GetLastEnvironMoniterAverageInfo 查询指定条件下的环境设备(最新的)均值数据
// @param: datetime 表示只查询指定时间点之前的最新的记录
func GetLastEnvironMoniterAverageInfo(db *gorm.DB, log log.Logger, customerID uint64, groupID uint64, datetime string) (envAverageData *protos.EnvironMonitorSectionAverageData, err error) {
	envRecords, err := getLastEnvironMoniterInfos(db, log, customerID, groupID, datetime)
	if err != nil {
		return
	}
	sum := len(envRecords)
	var temperature, pm025, noise, humidity float64
	for _, envRecord := range envRecords {
		temperature += envRecord.Temperature
		pm025 += envRecord.PM025
		noise += envRecord.Noise
		humidity += envRecord.Humidity
	}
	temperatureAverage := temperature / float64(sum)
	temperatureAverage = Round(temperatureAverage, 1)
	pm025Average := pm025 / float64(sum)
	pm025Average = Round(pm025Average, 1)
	noiseAverage := noise / float64(sum)
	noiseAverage = Round(noiseAverage, 1)
	humidityAverage := humidity / float64(sum)
	humidityAverage = Round(humidityAverage, 1)

	envAverageData = &protos.EnvironMonitorSectionAverageData{
		Temperature: temperatureAverage,
		Noise:       noiseAverage,
		PM025:       pm025Average,
		Humidity:    humidityAverage,
	}
	return
}

// HistoryMapInDate ...
type HistoryMapInDate struct {
	DateKey string // 日期字符串, 一般为12-31形式
	Datas   []*protos.EnvironMonitorSectionAverageData
}

// GetEnvironMonitorSectionAverageData 查询指定设备指定日期之间的环境均值数据历史.
// @param: dateFrom 起始时间点. 格式为2019-01-01
// @param: dateTo 截止时间点. 格式与dateFrom相同, 包括当天. 所以实际在调用sql查询时, 截止时间总是为dateTo + 1.
//         如果dateTo大于了发起请求当时的时间点, 则以0值补全.
// @param: delta 间隔时间1-24, 单位为小时, 默认为3
func GetEnvironMonitorSectionAverageData(db *gorm.DB, log log.Logger, customerID uint64, groupID uint64, dateFrom string, dateTo string, delta uint64) (historyDataList []map[string]interface{}, err error) {
	historyDataList = []map[string]interface{}{}
	// 每隔3小时查询当时环境值
	h1 := "1h"
	if delta == 0 {
		delta = 3
	}

	startDate, endDate, now, err := timesection.GetTimeSection(dateFrom, dateTo)
	if err != nil {
		log.Errorf("get time section failed in GetEnvironMonitorSectionAverageData(): %s", err.Error())
		return
	}
	log.Debugf("startDate: %+v, endDate: %+v", startDate, endDate)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	deltaTime, _ := time.ParseDuration(h1)
	deltaTime = deltaTime * time.Duration(delta)

	historyMapsInDate := []*HistoryMapInDate{}
	for {
		if startDate.After(endDate) || startDate.Equal(endDate) {
			break
		}
		dateKey := fmt.Sprintf("%s", startDate.In(loc).Format("01-02"))
		hourKey := fmt.Sprintf("%d时", startDate.In(loc).Hour())
		averageData := &protos.EnvironMonitorSectionAverageData{
			Section: hourKey,
		}
		// 如果构造的时间点已经超过了当前此刻, 则返回0, 不再查询.
		if startDate.Before(now) {
			dateStr := startDate.In(loc).Format("2006-01-02 15:04:05")
			_averageData, err := GetLastEnvironMoniterAverageInfo(db, log, customerID, groupID, dateStr)
			if err != nil {
				return nil, err
			}
			averageData.Temperature = _averageData.Temperature
			averageData.PM025 = _averageData.PM025
			averageData.Humidity = _averageData.Humidity
			averageData.Noise = _averageData.Noise

			// 将当前时间点平均值加入列表
			exist := false
			for _, mapInDate := range historyMapsInDate {
				if mapInDate.DateKey == dateKey {
					exist = true
					mapInDate.Datas = append(mapInDate.Datas, averageData)
					break
				}
			}
			if !exist {
				historyMapsInDate = append(historyMapsInDate, &HistoryMapInDate{
					DateKey: dateKey,
					Datas:   []*protos.EnvironMonitorSectionAverageData{averageData},
				})
			}
		}
		startDate = startDate.Add(deltaTime)
	}

	// 原型规定, 单日历史按小时展示. 超过单日, 则按天计算平均值展示
	if len(historyMapsInDate) > 1 {
		for _, historyMap := range historyMapsInDate {
			var temperature, pm025, noise, humidity float64
			sum := len(historyMap.Datas)
			for _, data := range historyMap.Datas {
				temperature += data.Temperature
				pm025 += data.PM025
				noise += data.Noise
				humidity += data.Humidity
			}
			historyItem := map[string]interface{}{
				"Day":         historyMap.DateKey,
				"Temperature": temperature / float64(sum),
				"PM025":       pm025 / float64(sum),
				"Noise":       noise / float64(sum),
				"Humidity":    humidity / float64(sum),
			}
			historyDataList = append(historyDataList, historyItem)
		}
	} else {
		for _, historyMap := range historyMapsInDate {
			for _, data := range historyMap.Datas {
				historyItem := map[string]interface{}{
					"Hour":        data.Section,
					"Temperature": data.Temperature,
					"PM025":       data.PM025,
					"Noise":       data.Noise,
					"Humidity":    data.Humidity,
				}
				historyDataList = append(historyDataList, historyItem)
			}
		}
	}

	return
}

func deviceToEnvironMonitorDevice(db *gorm.DB, log log.Logger, device *protos.Device, envDevice *protos.EnvironMonitorDevice) (err error) {
	envDevice.ID = device.ID
	envDevice.Name = device.Name
	envDevice.SerialNumber = device.SerialNumber
	envDevice.DeviceModel = device.DeviceModel
	envDevice.DeviceModelID = device.DeviceModelID
	envDevice.DeviceType = device.DeviceType
	envDevice.DeviceTypeID = device.DeviceTypeID
	envDevice.Group = device.Group
	envDevice.GroupID = device.GroupID
	envDevice.CustomerID = device.CustomerID
	envDevice.Position = device.Position
	envDevice.Latitude = device.Latitude
	envDevice.Longitude = device.Longitude
	envDevice.Actived = device.Actived
	envDevice.CreatedAt = device.CreatedAt

	msg, err := getLastEnvironMonitorInfo(db, log, device.CustomerID, device.SerialNumber)
	if err != nil {
		err = nil
		envDevice.StatusCode = 2
		envDevice.Status = "离线"
	} else {
		envDevice.StatusCode = 1
		envDevice.Status = "正常"
		envDevice.Temperature = msg.Temperature
		envDevice.Humidity = msg.Humidity
		envDevice.Noise = msg.Noise
		envDevice.PM025 = msg.PM025
	}
	return
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}
