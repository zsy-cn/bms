package service

import (
	"errors"
	"fmt"

	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
)

// GetDeviceModelList ...
// 暂不做分页处理
func (core *Core) GetDeviceModelList(req *protos.GetDeviceModelsRequest) (deviceModelList *protos.DeviceModelList, err error) {
	core.logger.Debug("get device model list in GetDeviceModelList()")
	// 先创建List对象, 查询出错时返回空列表而不是错误
	deviceModelList = &protos.DeviceModelList{
		List:  []*protos.DeviceModel{},
		Count: 0,
	}

	query := core.db.Model(&model.DeviceModel{})

	// 首先得到count总量
	var count uint64
	err = query.Count(&count).Error
	if err != nil {
		core.logger.Errorf("find device model count failed: %s", err.Error())
		return
	}

	deviceModelRecords := []*model.DeviceModel{}
	err = query.Find(&deviceModelRecords).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() != "record not found" {
			core.logger.Errorf("find device models failed: %s", err.Error())
		}
		// 查询出错返回空记录, 忽略错误
		err = nil
		return
	}

	_deviceModelList := []*protos.DeviceModel{}
	for _, deviceModelRecord := range deviceModelRecords {
		deviceModelPb := &protos.DeviceModel{}
		err = deviceModelModel2Pb(deviceModelRecord, deviceModelPb)
		if err != nil {
			core.logger.Errorf("transform device model: %s model to protobuf object failed: %s", deviceModelRecord.Name, err.Error())
			return deviceModelList, nil
		}
		_deviceModelList = append(_deviceModelList, deviceModelPb)
	}
	deviceModelList.List = _deviceModelList
	deviceModelList.Count = count

	return
}

// AddDeviceModel ...
func (core *Core) AddDeviceModel(deviceModelPb *protos.DeviceModel) (err error) {
	deviceModelModel := &model.DeviceModel{}
	err = deviceModelPb2Model(deviceModelPb, deviceModelModel)
	if err != nil {
		core.logger.Errorf("transform device model: %s protobuf object to model failed: %s", deviceModelPb.Name, err.Error())
		return
	}
	err = core.db.Create(deviceModelModel).Error
	if err != nil {
		core.logger.Errorf("insert device model: %s into database failed: %s", deviceModelPb.Name, err.Error())
		return
	}
	return
}

// UpdateDeviceModel ...
// update操作目前只能修改分组名称及可用性, 其他的不能改, 所以不能修改外键字段
func (core *Core) UpdateDeviceModel(deviceModelPb *protos.DeviceModel) (err error) {
	record := &model.DeviceModel{}
	err = core.db.First(record, deviceModelPb.ID).Error
	if err != nil {
		core.logger.Errorf("find device model: %s failed in update: %s", deviceModelPb.Name, err.Error())
		return
	}

	deviceModelMap := map[string]interface{}{}
	err = deviceModelPb2Map(deviceModelPb, deviceModelMap)
	if err != nil {
		core.logger.Errorf("transform device model: %s protobuf object to map failed: %s", deviceModelPb.Name, err.Error())
		return
	}
	err = core.db.Model(record).Update(deviceModelMap).Error
	if err != nil {
		core.logger.Errorf("update device model: %s failed: %s", deviceModelPb.Name, err.Error())
		return
	}
	return
}

// DeleteDeviceModel ...
// 如果一个型号下还有设备的话, 不要删除, 更不要级联删除
func (core *Core) DeleteDeviceModel(req *protos.DeleteRequest) (err error) {
	core.logger.Debugf("delete device model with id: %d in DeleteDeviceModel()", req.ID)
	record := &model.DeviceModel{}
	err = core.db.First(record, req.ID).Error
	if err != nil {
		// 不管是不是record not found错误, 都需要返回
		core.logger.Errorf("find device model: %d failed in delete: %s", req.ID, err.Error())
		return
	}
	// deviceCount := core.db.Model(record).Association("Manufacturer").Count()
	var deviceCount uint64
	err = core.db.Model(&model.Device{}).Where(&model.Device{DeviceModelID: record.ID}).Count(&deviceCount).Error
	if err != nil {
		core.logger.Errorf("get device count of this model: %s failed: %s", record.Name, err.Error())
		return
	}
	core.logger.Debugf("device count of this model: %s is: %d", record.Name, deviceCount)
	if deviceCount != 0 {
		errStr := fmt.Sprintf("there are still some devices of this device model: %s, you can't delete it", record.Name)
		core.logger.Errorf(errStr)
		err = errors.New(errStr)
		return
	}

	err = core.db.Delete(record).Error
	if err != nil {
		core.logger.Errorf("delete device model: %d failed in delete: %s", req.ID, err.Error())
		return
	}

	return
}

func deviceModelModel2Pb(mod *model.DeviceModel, pb *protos.DeviceModel) (err error) {
	pb.ID = mod.ID
	pb.Name = mod.Name
	pb.Description = mod.Description
	pb.ManufacturerID = mod.ManufacturerID
	pb.DeviceTypeID = mod.DeviceTypeID
	return
}

// deviceModelPb2Model 一般用于Create操作
func deviceModelPb2Model(pb *protos.DeviceModel, mod *model.DeviceModel) (err error) {
	mod.ID = pb.ID
	mod.Name = pb.Name
	mod.Description = pb.Description
	mod.ManufacturerID = pb.ManufacturerID
	mod.DeviceTypeID = pb.DeviceTypeID
	return
}

func deviceModelPb2Map(pb *protos.DeviceModel, theMap map[string]interface{}) (err error) {
	theMap["name"] = pb.Name
	theMap["description"] = pb.Description
	theMap["manufacturer_id"] = pb.ManufacturerID
	theMap["device_type_id"] = pb.DeviceTypeID
	return
}
