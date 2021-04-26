package service

import (
	"github.com/jinzhu/gorm"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

// Contact ...
type Contact struct {
	logger *log.Logger
	db     *gorm.DB
}

// New ...
func New(logger *log.Logger) (contact *Contact, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	contact = &Contact{
		db:     db,
		logger: logger,
	}
	return
}

// Get 获取指定ID的Contact对象
// 如果查询出错就返回空记录, 不要返回错误了
func (c *Contact) Get(req *protos.GetContactRequest) (contactPb *protos.Contact, err error) {
	contactModel := &model.Contact{}
	err = c.db.First(contactModel, req.ID).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() != "record not found" {
			c.logger.Errorf("find contact: %d failed: %s", req.ID, err.Error())
		}
		err = nil
		return
	}

	contactPb = &protos.Contact{}
	err = model2Pb(contactModel, contactPb)
	if err != nil {
		c.logger.Errorf("transform contact: %s model to protobuf object failed: %s", contactModel.Name, err.Error())
		return
	}
	return
}

// GetList 获取指定Customer客户名下的所有Contact联系人
func (c *Contact) GetList(req *protos.GetContactsRequest) (contactList *protos.ContactList, err error) {
	c.logger.Debugf("get contact list for customer: %d", req.CustomerID)
	// 先创建List对象, 查询出错时返回空列表而不是错误
	contactList = &protos.ContactList{
		List:  []*protos.Contact{},
		Count: 0,
	}

	query := c.db.Model(&model.Contact{})

	query = query.Where(&model.Contact{CustomerID: req.CustomerID})
	var count uint64
	err = query.Count(&count).Error
	if err != nil {
		c.logger.Errorf("find contact count failed: %s", err.Error())
		return
	}

	contactModels := []*model.Contact{}
	err = query.Find(&contactModels).Error
	if err != nil {
		if err.Error() != "record not found" {
			c.logger.Errorf("find contacts for customer: %d failed: %s", req.CustomerID, err.Error())
		}
		err = nil
		return
	}

	_contactList := []*protos.Contact{}
	for _, contactModel := range contactModels {
		contactPb := &protos.Contact{}
		err = model2Pb(contactModel, contactPb)
		if err != nil {
			c.logger.Errorf("transform contact: %s model to protobuf object failed: %s", contactModel.Name, err.Error())
			err = nil
			return
		}
		_contactList = append(_contactList, contactPb)
	}
	contactList.List = _contactList
	contactList.Count = count
	return
}

// Add ...
func (c *Contact) Add(contactPb *protos.Contact) (err error) {
	contactModel := &model.Contact{}
	err = pb2Model(contactPb, contactModel)
	if err != nil {
		c.logger.Errorf("transform contact: %s protobuf object to model failed: %s", contactPb.Name, err.Error())
		return
	}
	err = c.db.Create(contactModel).Error
	if err != nil {
		c.logger.Errorf("insert contact: %s into database failed: %s", contactPb.Name, err.Error())
		return
	}
	return
}

// Update ...
func (c *Contact) Update(contactPb *protos.Contact) (err error) {
	record := &model.Contact{}
	err = c.db.First(record, contactPb.ID).Error
	if err != nil {
		c.logger.Errorf("find contact: %s : %s", contactPb.Name, err.Error())
		return
	}

	contactMap := map[string]interface{}{}
	err = pb2Map(contactPb, contactMap)
	if err != nil {
		c.logger.Errorf("transform contact: %s protobuf object to map failed: %s", contactPb.Name, err.Error())
		return
	}
	err = c.db.Model(record).Update(contactMap).Error
	if err != nil {
		c.logger.Errorf("update contact: %s failed: %s", contactPb.Name, err.Error())
		return
	}
	return
}

// Delete ...
func (c *Contact) Delete(contactPb *protos.Contact) (err error) {
	record := &model.Contact{}
	err = c.db.First(record, contactPb.ID).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() == "record not found" {
			err = nil
		} else {
			c.logger.Errorf("find contact: %s failed: %s", contactPb.Name, err.Error())
		}
		return
	}
	err = c.db.Delete(record).Error
	if err != nil {
		c.logger.Errorf("delete contact: %s failed: %s", contactPb.Name, err.Error())
		return
	}
	return
}

// pb2Model 一般用于Create操作
func pb2Model(pb *protos.Contact, model *model.Contact) (err error) {
	model.ID = pb.ID
	model.CustomerID = pb.CustomerID
	model.Name = pb.Name
	model.Phone = pb.Phone
	model.Email = pb.Email
	return
}

// pb2Map 一般用于Update操作(gorm对结构体字段为默认值如0, false, 或""不做操作)
func pb2Map(pb *protos.Contact, theMap map[string]interface{}) (err error) {
	// theMap["id"] = pb.ID
	theMap["customer_id"] = pb.CustomerID
	theMap["name"] = pb.Name
	theMap["phone"] = pb.Phone
	theMap["email"] = pb.Email
	return
}

// model2Pb 一般用于Get(List)操作
func model2Pb(model *model.Contact, pb *protos.Contact) (err error) {
	pb.ID = model.ID
	pb.CustomerID = model.CustomerID
	pb.Name = model.Name
	pb.Phone = model.Phone
	pb.Email = model.Email
	return
}
