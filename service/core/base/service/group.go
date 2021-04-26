package service

import (
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
)

// GetGroupList ...
// 暂不做分页处理
func (core *Core) GetGroupList(req *protos.GetGroupsRequest) (groupList *protos.GroupList, err error) {
	core.logger.Debug("get group list in GetGroupList()")
	// 先创建List对象, 查询出错时返回空列表而不是错误
	groupList = &protos.GroupList{
		List:  []*protos.Group{},
		Count: 0,
	}

	query := core.db.Model(&model.Group{})
	// ...这里应是条件查询语句
	query = query.Where(&model.Group{CustomerID: req.CustomerID})

	// 首先得到count总量
	var count uint64
	err = query.Count(&count).Error
	if err != nil {
		core.logger.Errorf("find group count failed: %s", err.Error())
		return
	}

	groupModels := []*model.Group{}
	err = query.Find(&groupModels).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() != "record not found" {
			core.logger.Errorf("find groups failed: %s", err.Error())
		}
		// 查询出错返回空记录, 忽略错误
		err = nil
		return
	}

	_groupList := []*protos.Group{}
	for _, groupModel := range groupModels {
		groupPb := &protos.Group{}
		err = model2Pb(groupModel, groupPb)
		if err != nil {
			core.logger.Errorf("transform group: %s model to protobuf object failed: %s", groupModel.Name, err.Error())
			return groupList, nil
		}
		_groupList = append(_groupList, groupPb)
	}
	groupList.List = _groupList
	groupList.Count = count

	return
}

// AddGroup ...
func (core *Core) AddGroup(groupPb *protos.Group) (err error) {
	err = core.checkForeignKey(groupPb)
	if err != nil {
		// 日志在checkForeignKey函数中已经打过了
		return
	}
	groupModel := &model.Group{}
	err = pb2Model(groupPb, groupModel)
	if err != nil {
		core.logger.Errorf("transform group: %s protobuf object to model failed: %s", groupPb.Name, err.Error())
		return
	}
	err = core.db.Create(groupModel).Error
	if err != nil {
		core.logger.Errorf("insert group: %s into database failed: %s", groupPb.Name, err.Error())
		return
	}
	return
}

// UpdateGroup ...
// update操作目前只能修改分组名称及可用性, 其他的不能改, 所以不能修改外键字段
func (core *Core) UpdateGroup(groupPb *protos.Group) (err error) {
	record := &model.Group{}
	err = core.db.First(record, groupPb.ID).Error
	if err != nil {
		core.logger.Errorf("find group: %s failed in update: %s", groupPb.Name, err.Error())
		return
	}

	groupMap := map[string]interface{}{}
	err = pb2Map(groupPb, groupMap)
	if err != nil {
		core.logger.Errorf("transform group: %s protobuf object to map failed: %s", groupPb.Name, err.Error())
		return
	}
	err = core.db.Model(record).Update(groupMap).Error
	if err != nil {
		core.logger.Errorf("update group: %s failed: %s", groupPb.Name, err.Error())
		return
	}
	return
}

// DeleteGroup ...
// 删除分组的同时也要把分组下的设备删除
func (core *Core) DeleteGroup(req *protos.DeleteGroupRequest) (err error) {
	record := &model.Group{}
	err = core.db.First(record, req.ID).Error
	if err != nil {
		// 不管是不是record not found错误, 都需要返回
		core.logger.Errorf("find group: %d failed in delete: %s", req.ID, err.Error())
		return
	}
	deviceTypeID := record.DeviceTypeID
	deviceTypeRecord := &model.DeviceType{}
	err = core.db.First(deviceTypeRecord, deviceTypeID).Error
	if err != nil {
		core.logger.Errorf("find device type: %d failed: %s", record.DeviceTypeID, err.Error())
		return
	}

	tx := core.db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	err = tx.Delete(record).Error
	if err != nil {
		core.logger.Errorf("delete group: %d failed in delete: %s", req.ID, err.Error())
		return
	}

	// 这里需要调用device服务删除指定group的设备.
	return
}

func (core *Core) checkForeignKey(groupPb *protos.Group) (err error) {
	customerModel := &model.Customer{}
	err = core.db.First(customerModel, groupPb.CustomerID).Error
	if err != nil {
		core.logger.Errorf("get the customer: %d failed: %s", groupPb.CustomerID, err.Error())
		return
	}
	deviceTypeModel := &model.DeviceType{}
	err = core.db.First(deviceTypeModel, groupPb.DeviceTypeID).Error
	if err != nil {
		core.logger.Errorf("get the device type: %d failed: %s", groupPb.CustomerID, err.Error())
		return
	}
	return
}

func model2Pb(model *model.Group, pb *protos.Group) (err error) {
	pb.ID = model.ID
	pb.Name = model.Name
	pb.CustomerID = model.CustomerID
	pb.DeviceTypeID = model.DeviceTypeID
	return
}

// pb2Model 一般用于Create操作
func pb2Model(pb *protos.Group, model *model.Group) (err error) {
	model.ID = pb.ID
	model.Name = pb.Name
	model.CustomerID = pb.CustomerID
	model.DeviceTypeID = pb.DeviceTypeID
	// 创建操作中不需要写入Status字段
	// model.Status = pb.Status
	return
}

func pb2Map(pb *protos.Group, theMap map[string]interface{}) (err error) {
	theMap["name"] = pb.Name
	theMap["status"] = pb.Status
	return
}
