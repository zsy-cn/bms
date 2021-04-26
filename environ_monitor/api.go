package environ_monitor

import (
	"github.com/zsy-cn/bms/protos"
)

type EnvironMonitorService interface {
	GetEnvironMonitorDevice(customerID uint64, deviceSN string) (device *protos.EnvironMonitorDevice, err error)
	GetEnvironMonitorDevices(req *protos.GetDevicesRequestForCustomer) (deviceList *protos.EnvironMonitorDeviceList, err error)
	GetEnvironMonitorDeviceStatus(customerID uint64, deviceSN string) (status uint64, err error)
	GetLastEnvironMoniterAverageInfo(customerID uint64, groupID uint64, datetime string) (envAverageData *protos.EnvironMonitorSectionAverageData, err error)
	GetEnvironMonitorSectionAverageData(customer uint64, groupID uint64, dateFrom string, dateTo string, delta uint64) (historyDataList []map[string]interface{}, err error)

	GetStandardWeatherInfo(url string) (msg *EnvironMonitor, err error)
}
