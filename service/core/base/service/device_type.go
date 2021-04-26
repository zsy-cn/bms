package service

import (
	"errors"
	"fmt"

	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
)

// GetDeviceTypeList ...
// 暂不做分页处理
func (core *Core) GetDeviceTypeList(req *protos.GetDeviceTypesRequest) (deviceTypeList *protos.DeviceTypeList, err error) {
	core.logger.Debug("get device type list in GetDeviceTypeList()")
	// 先创建List对象, 查询出错时返回空列表而不是错误
	deviceTypeList = &protos.DeviceTypeList{
		List:  []*protos.DeviceType{},
		Count: 0,
	}

	query := core.db.Model(&model.DeviceType{})

	// 首先得到count总量
	var count uint64
	err = query.Count(&count).Error
	if err != nil {
		core.logger.Errorf("find device type count failed: %s", err.Error())
		return
	}

	deviceTypeRecords := []*model.DeviceType{}
	err = query.Find(&deviceTypeRecords).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() != "record not found" {
			core.logger.Errorf("find device types failed: %s", err.Error())
		}
		// 查询出错返回空记录, 忽略错误
		err = nil
		return
	}

	_deviceTypeList := []*protos.DeviceType{}
	for _, deviceTypeRecord := range deviceTypeRecords {
		deviceTypePb := &protos.DeviceType{}
		err = deviceTypeModel2Pb(deviceTypeRecord, deviceTypePb)
		if err != nil {
			core.logger.Errorf("transform device type: %s model to protobuf object failed: %s", deviceTypeRecord.Name, err.Error())
			return deviceTypeList, nil
		}
		_deviceTypeList = append(_deviceTypeList, deviceTypePb)
	}
	deviceTypeList.List = _deviceTypeList
	deviceTypeList.Count = count

	return
}

// AddDeviceType ...
func (core *Core) AddDeviceType(deviceTypePb *protos.DeviceType) (err error) {
	deviceTypeModel := &model.DeviceType{}
	err = deviceTypePb2Model(deviceTypePb, deviceTypeModel)
	if err != nil {
		core.logger.Errorf("transform device type: %s protobuf object to model failed: %s", deviceTypePb.Name, err.Error())
		return
	}
	err = core.db.Create(deviceTypeModel).Error
	if err != nil {
		core.logger.Errorf("insert device type: %s into database failed: %s", deviceTypePb.Name, err.Error())
		return
	}
	return
}

// UpdateDeviceType ...
// update操作目前只能修改分组名称及可用性, 其他的不能改, 所以不能修改外键字段
func (core *Core) UpdateDeviceType(deviceTypePb *protos.DeviceType) (err error) {
	record := &model.DeviceType{}
	err = core.db.First(record, deviceTypePb.ID).Error
	if err != nil {
		core.logger.Errorf("find device type: %s failed in update: %s", deviceTypePb.Name, err.Error())
		return
	}

	deviceTypeMap := map[string]interface{}{}
	err = deviceTypePb2Map(deviceTypePb, deviceTypeMap)
	if err != nil {
		core.logger.Errorf("transform device type: %s protobuf object to map failed: %s", deviceTypePb.Name, err.Error())
		return
	}
	err = core.db.Model(record).Update(deviceTypeMap).Error
	if err != nil {
		core.logger.Errorf("update device type: %s failed: %s", deviceTypePb.Name, err.Error())
		return
	}
	return
}

// DeleteDeviceType ...
// 如果一个设备类型下有设备或是型号的话, 不要删除, 更不要级联删除
func (core *Core) DeleteDeviceType(req *protos.DeleteRequest) (err error) {
	core.logger.Debugf("delete device type with id: %d in DeleteDeviceType()", req.ID)
	record := &model.DeviceType{}
	err = core.db.First(record, req.ID).Error
	if err != nil {
		// 不管是不是record not found错误, 都需要返回
		core.logger.Errorf("find device type: %d failed in delete: %s", req.ID, err.Error())
		return
	}
	// deviceCount := core.db.Model(record).Association("Manufacturer").Count()
	var deviceCount uint64
	err = core.db.Model(&model.Device{}).Where(&model.Device{DeviceTypeID: record.ID}).Count(&deviceCount).Error
	if err != nil {
		core.logger.Errorf("get device count of this type: %s failed: %s", record.Name, err.Error())
		return
	}
	core.logger.Debugf("device count of this type: %s is: %d", record.Name, deviceCount)
	if deviceCount != 0 {
		errStr := fmt.Sprintf("there are still some devices of this device type: %s, you can't delete it", record.Name)
		core.logger.Errorf(errStr)
		err = errors.New(errStr)
		return
	}

	err = core.db.Delete(record).Error
	if err != nil {
		core.logger.Errorf("delete device type: %d failed in delete: %s", req.ID, err.Error())
		return
	}

	return
}

func deviceTypeModel2Pb(mod *model.DeviceType, pb *protos.DeviceType) (err error) {
	pb.ID = mod.ID
	pb.Key = mod.Key
	pb.Name = mod.Name
	pb.IsSensor = mod.IsSensor
	pb.Properties = []string(*mod.Properties)
	return
}

// deviceTypePb2Model 一般用于Create操作
func deviceTypePb2Model(pb *protos.DeviceType, mod *model.DeviceType) (err error) {
	prop := model.StringArray(pb.Properties)
	mod.ID = pb.ID
	mod.Key = pb.Key
	mod.Name = pb.Name
	mod.IsSensor = pb.IsSensor
	mod.Properties = &prop

	return
}

func deviceTypePb2Map(pb *protos.DeviceType, theMap map[string]interface{}) (err error) {
	theMap["name"] = pb.Name
	theMap["key"] = pb.Key
	theMap["is_sensor"] = pb.IsSensor
	theMap["properties"] = pb.Properties
	return
}
