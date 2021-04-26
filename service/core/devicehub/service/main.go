package service

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/pagination"
)

// Device 增删服务
type Device struct {
	logger          *log.Logger
	db              *gorm.DB
	deviceSensorCli protos.DeviceSensorServiceClient
}

// New ...
func New(logger *log.Logger) (device *Device, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	address := viper.GetString("device-sensor-addr")
	deviceSensorConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("connect devicesensor failed: %s", err.Error())
		return
	}
	deviceSensorCli := protos.NewDeviceSensorServiceClient(deviceSensorConn)

	device = &Device{
		logger:          logger,
		db:              db,
		deviceSensorCli: deviceSensorCli,
	}
	return
}

// GetList 用于条件查询, 且查询的主要是Device表中各种设备的公共属性
// 不需要到指定设备表中查询
func (d *Device) GetList(req *protos.GetDevicesRequest) (deviceList *protos.DeviceList, err error) {
	// 先创建List对象, 查询出错时返回空列表而不是错误
	deviceList = &protos.DeviceList{
		List:  []*protos.Device{},
		Count: 0,
	}

	query := d.db.Model(&model.Device{})
	// ...这里应是条件查询语句

	// 首先得到count总量
	var count uint64
	err = query.Count(&count).Error
	if err != nil {
		d.logger.Errorf("find device count failed: %s", err.Error())
		return
	}
	// 构建分页查询语句
	query = pagination.BuildPaginationQuery(query, req.Pagination)

	deviceModels := []*model.Device{}
	err = query.Find(&deviceModels).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() != "record not found" {
			d.logger.Errorf("find device failed: %s", err.Error())
		}
		// 查询出错返回空记录, 忽略错误
		err = nil
		return
	}

	_deviceList := []*protos.Device{}
	for _, deviceModel := range deviceModels {
		devicePb := &protos.Device{}
		err = model2Pb(deviceModel, devicePb)
		if err != nil {
			d.logger.Errorf("transform device: %s model to protobuf object failed: %s", deviceModel.Name, err.Error())
			return deviceList, nil
		}
		_deviceList = append(_deviceList, devicePb)
	}
	deviceList.List = _deviceList
	deviceList.Count = count
	return
}

// Add ...
func (d *Device) Add(devicePb *protos.Device) (err error) {
	err = d.checkForeignKey(devicePb)
	if err != nil {
		// 日志在checkForeignKey函数中已经打过了
		return
	}
	deviceModel := &model.Device{}
	err = pb2Model(devicePb, deviceModel)
	if err != nil {
		d.logger.Errorf("transform device: %s protobuf object to model failed: %s", devicePb.Name, err.Error())
		return
	}

	tx := d.db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	err = tx.Create(deviceModel).Error
	if err != nil {
		d.logger.Errorf("insert device: %s into database failed: %s", devicePb.Name, err.Error())
		return
	}

	// 这里需要调用不同设备的接口来创建新记录
	err = d.DispatchAdd(deviceModel, devicePb.ExtraInfo)
	if err != nil {
		// 日志在DispatchAdd函数体内部打印
		return
	}

	return
}

// Update ...
// Device表的Update操作目前只更新devices表中的公共属性,
// 像lora设备的devEUI, sound_box的sqr编码等, 不可以更新,
// 所以这里不需要调用grpc接口
func (d *Device) Update(devicePb *protos.Device) (err error) {
	err = d.checkForeignKey(devicePb)
	if err != nil {
		// 日志在checkForeignKey函数中已经打过了
		return
	}

	deviceRecord := &model.Device{}
	err = d.db.First(deviceRecord, devicePb.ID).Error
	if err != nil {
		d.logger.Errorf("find device: %s failed in update: %s", devicePb.Name, err.Error())
		return
	}

	deviceMap := map[string]interface{}{}
	err = pb2Map(devicePb, deviceMap)
	if err != nil {
		d.logger.Errorf("transform device: %s protobuf object to model failed: %s", devicePb.Name, err.Error())
		return
	}
	err = d.db.Model(deviceRecord).Update(deviceMap).Error
	if err != nil {
		d.logger.Errorf("update device: %s failed: %s", devicePb.Name, err.Error())
		return
	}
	return
}

// Delete ...
func (d *Device) Delete(req *protos.DeleteDeviceRequest) (err error) {
	record := &model.Device{}
	err = d.db.First(record, req.ID).Error
	if err != nil {
		// 不管是不是record not found错误, 都需要返回
		d.logger.Errorf("find device: %d failed in delete: %s", req.ID, err.Error())
		return
	}

	tx := d.db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	err = tx.Delete(record).Error
	if err != nil {
		d.logger.Errorf("delete device: %d failed in delete: %s", req.ID, err.Error())
		return
	}

	err = d.DispatchDelete(record)
	if err != nil {
		// 日志在DispatchDelete函数体内部打印
		return
	}

	return
}

// DeleteGroup ...
func (d *Device) DeleteGroup(req *protos.DeleteGroupDeviceRequest) (err error) {
	deviceModels := []*model.Device{}
	err = d.db.Where(&model.Device{GroupID: req.GroupID}).Find(&deviceModels).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() == "record not found" {
			err = nil
		} else {
			d.logger.Errorf("find devices of this group: %d failed: %s", req.GroupID, err.Error())
		}
		return
	}

	// 这里批量删除, 但是是在Delete方法内创建事务.
	// 因为要考虑到loraclient删除设备时批量删除也是单个设备删除的,
	// 所以这里一个一个地删.
	for _, deviceModel := range deviceModels {
		deleteDeviceReq := &protos.DeleteDeviceRequest{
			ID: deviceModel.ID,
		}
		err = d.Delete(deleteDeviceReq)
		if err != nil {
			// 日志在Delete函数体内部打印, 可以直接返回
			return
		}
	}

	return
}

func (d *Device) checkForeignKey(devicePb *protos.Device) (err error) {
	// 用户检查
	customerModel := &model.Customer{}
	err = d.db.First(customerModel, devicePb.CustomerID).Error
	if err != nil {
		d.logger.Errorf("get the customer: %d failed: %s", devicePb.CustomerID, err.Error())
		return
	}
	// 分组检查
	groupModel := &model.Group{}
	err = d.db.First(groupModel, devicePb.GroupID).Error
	if err != nil {
		d.logger.Errorf("get the group: %d failed: %s", devicePb.GroupID, err.Error())
		return
	}
	// 设备类型检查
	deviceTypeModel := &model.DeviceType{}
	err = d.db.First(deviceTypeModel, devicePb.DeviceTypeID).Error
	if err != nil {
		d.logger.Errorf("get the device type: %d failed: %s", devicePb.DeviceTypeID, err.Error())
		return
	}
	// 设备型号检查
	deviceModelModel := &model.DeviceModel{}
	err = d.db.First(deviceModelModel, devicePb.DeviceModelID).Error
	if err != nil {
		d.logger.Errorf("get the device type: %d failed: %s", devicePb.DeviceModelID, err.Error())
		return
	}
	return
}

func pb2Model(pb *protos.Device, model *model.Device) (err error) {
	model.Name = pb.Name
	model.SerialNumber = pb.SerialNumber
	model.Description = pb.Description
	model.CustomerID = pb.CustomerID
	model.GroupID = pb.GroupID
	model.DeviceTypeID = pb.DeviceTypeID
	model.DeviceModelID = pb.DeviceModelID

	model.Latitude = pb.Latitude
	model.Longitude = pb.Longitude
	return
}

func pb2Map(pb *protos.Device, theMap map[string]interface{}) (err error) {
	theMap["name"] = pb.Name
	theMap["serialNumber"] = pb.SerialNumber
	theMap["description"] = pb.Description
	theMap["customerID"] = pb.CustomerID
	theMap["groupID"] = pb.GroupID
	theMap["deviceTypeID"] = pb.DeviceTypeID
	theMap["deviceModelID"] = pb.DeviceModelID

	theMap["latitude"] = pb.Latitude
	theMap["longitude"] = pb.Longitude
	theMap["status"] = pb.Status
	return
}

func model2Pb(model *model.Device, pb *protos.Device) (err error) {
	pb.ID = model.ID
	pb.Name = model.Name
	pb.SerialNumber = model.SerialNumber
	pb.Description = model.Description
	pb.CustomerID = model.CustomerID
	pb.GroupID = model.GroupID
	pb.DeviceTypeID = model.DeviceTypeID
	pb.DeviceModelID = model.DeviceModelID

	pb.Latitude = model.Latitude
	pb.Longitude = model.Longitude
	return
}
