package service

import (
	"errors"
	"fmt"

	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
)

// GetManufacturerList ...
// 暂不做分页处理
func (core *Core) GetManufacturerList(req *protos.GetManufacturersRequest) (manufacturerList *protos.ManufacturerList, err error) {
	core.logger.Debug("get manufacture list in GetManufacturerList()")
	// 先创建List对象, 查询出错时返回空列表而不是错误
	manufacturerList = &protos.ManufacturerList{
		List:  []*protos.Manufacturer{},
		Count: 0,
	}

	query := core.db.Model(&model.Manufacturer{})

	// 首先得到count总量
	var count uint64
	err = query.Count(&count).Error
	if err != nil {
		core.logger.Errorf("find manufacture count failed: %s", err.Error())
		return
	}

	manufacturerRecords := []*model.Manufacturer{}
	err = query.Find(&manufacturerRecords).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() != "record not found" {
			core.logger.Errorf("find manufactures failed: %s", err.Error())
		}
		// 查询出错返回空记录, 忽略错误
		err = nil
		return
	}

	_manufacturerList := []*protos.Manufacturer{}
	for _, manufacturerRecord := range manufacturerRecords {
		manufacturerPb := &protos.Manufacturer{}
		err = manufacturerModel2Pb(manufacturerRecord, manufacturerPb)
		if err != nil {
			core.logger.Errorf("transform manufacture: %s model to protobuf object failed: %s", manufacturerRecord.Name, err.Error())
			return manufacturerList, nil
		}
		_manufacturerList = append(_manufacturerList, manufacturerPb)
	}
	manufacturerList.List = _manufacturerList
	manufacturerList.Count = count

	return
}

// AddManufacturer ...
func (core *Core) AddManufacturer(manufacturerPb *protos.Manufacturer) (err error) {
	manufacturerModel := &model.Manufacturer{}
	err = manufacturerPb2Model(manufacturerPb, manufacturerModel)
	if err != nil {
		core.logger.Errorf("transform manufacture: %s protobuf object to model failed: %s", manufacturerPb.Name, err.Error())
		return
	}
	err = core.db.Create(manufacturerModel).Error
	if err != nil {
		core.logger.Errorf("insert manufacture: %s into database failed: %s", manufacturerPb.Name, err.Error())
		return
	}
	return
}

// UpdateManufacturer ...
// update操作目前只能修改分组名称及可用性, 其他的不能改, 所以不能修改外键字段
func (core *Core) UpdateManufacturer(manufacturerPb *protos.Manufacturer) (err error) {
	record := &model.Manufacturer{}
	err = core.db.First(record, manufacturerPb.ID).Error
	if err != nil {
		core.logger.Errorf("find manufacture: %s failed in update: %s", manufacturerPb.Name, err.Error())
		return
	}

	manufacturerMap := map[string]interface{}{}
	err = manufacturerPb2Map(manufacturerPb, manufacturerMap)
	if err != nil {
		core.logger.Errorf("transform manufacture: %s protobuf object to map failed: %s", manufacturerPb.Name, err.Error())
		return
	}
	err = core.db.Model(record).Update(manufacturerMap).Error
	if err != nil {
		core.logger.Errorf("update manufacture: %s failed: %s", manufacturerPb.Name, err.Error())
		return
	}
	return
}

// DeleteManufacturer ...
// 如果一个设备厂商下有型号记录的话, 不要删除, 更不要级联删除
func (core *Core) DeleteManufacturer(req *protos.DeleteRequest) (err error) {
	core.logger.Debugf("delete manufacture with id: %d in DeleteManufacturer()", req.ID)
	record := &model.Manufacturer{}
	err = core.db.First(record, req.ID).Error
	if err != nil {
		// 不管是不是record not found错误, 都需要返回
		core.logger.Errorf("find manufacture: %d failed in delete: %s", req.ID, err.Error())
		return
	}
	var deviceModelCount uint64
	err = core.db.Model(&model.DeviceModel{}).Where(&model.DeviceModel{ManufacturerID: record.ID}).Count(&deviceModelCount).Error
	if err != nil {
		core.logger.Errorf("get device model count of this manufacturer: %s failed: %s", record.Name, err.Error())
		return
	}
	core.logger.Debugf("device count of this manufacturer: %s is: %d", record.Name, deviceModelCount)
	if deviceModelCount != 0 {
		errStr := fmt.Sprintf("there are still some models of this manufacturer: %s, you can't delete it", record.Name)
		core.logger.Errorf(errStr)
		err = errors.New(errStr)
		return
	}

	err = core.db.Delete(record).Error
	if err != nil {
		core.logger.Errorf("delete manufacture: %d failed in delete: %s", req.ID, err.Error())
		return
	}

	return
}

func manufacturerModel2Pb(mod *model.Manufacturer, pb *protos.Manufacturer) (err error) {
	pb.ID = mod.ID
	pb.Name = mod.Name
	return
}

// manufacturerPb2Model 一般用于Create操作
func manufacturerPb2Model(pb *protos.Manufacturer, mod *model.Manufacturer) (err error) {
	mod.ID = pb.ID
	mod.Name = pb.Name
	return
}

func manufacturerPb2Map(pb *protos.Manufacturer, theMap map[string]interface{}) (err error) {
	theMap["name"] = pb.Name
	return
}
