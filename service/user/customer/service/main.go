package service

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
	"github.com/zsy-cn/bms/util/pagination"
)

// Customer ...
type Customer struct {
	logger        *log.Logger
	db            *gorm.DB
	loraclientCli protos.LoraclientServiceClient
	contactCli    protos.ContactServiceClient
}

// New 实例化Customer服务对象
func New(logger *log.Logger) (customer *Customer, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	loraclientAddress := viper.GetString("loraclient-addr")
	loraclientConn, err := grpc.Dial(loraclientAddress, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("connect loraclient failed: %s", err.Error())
		return
	}
	loraclientCli := protos.NewLoraclientServiceClient(loraclientConn)
	contactAddress := viper.GetString("contact-addr")
	contactConn, err := grpc.Dial(contactAddress, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("connect contact failed: %s", err.Error())
		return
	}
	contactCli := protos.NewContactServiceClient(contactConn)
	customer = &Customer{
		logger:        logger,
		db:            db,
		loraclientCli: loraclientCli,
		contactCli:    contactCli,
	}
	return
}

// Get ...
func (c *Customer) Get(req *protos.GetCustomerRequest) (customerPb *protos.Customer, err error) {
	customerModel := &model.Customer{}
	err = c.db.First(customerModel, req.ID).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() != "record not found" {
			c.logger.Errorf("find customer: %d failed: %s", req.ID, err.Error())
		}
		err = nil
		return
	}

	customerPb = &protos.Customer{}
	contactList, err := c.getMyContacts(req.ID)
	if err != nil {
		err = nil
		return
	}
	err = model2Pb(customerModel, customerPb)
	if err != nil {
		c.logger.Errorf("transform customer: %s model to protobuf object failed in get: %s", customerModel.Name, err.Error())
		err = nil
		return
	}
	customerPb.Contacts = contactList.List
	return
}

// GetList ...
func (c *Customer) GetList(req *protos.GetCustomersRequest) (customerList *protos.CustomerList, err error) {
	// 先创建List对象, 查询出错时返回空列表而不是错误
	customerList = &protos.CustomerList{
		List:  []*protos.Customer{},
		Count: 0,
	}

	query := c.db.Model(&model.Customer{})
	// ...这里应是条件查询语句

	// 首先得到count总量
	var count uint64
	err = query.Count(&count).Error
	if err != nil {
		c.logger.Errorf("find customer count failed: %s", err.Error())
		return
	}
	// 构建分页查询语句
	query = pagination.BuildPaginationQuery(query, req.Pagination)

	customerModels := []*model.Customer{}
	err = query.Find(&customerModels).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() != "record not found" {
			c.logger.Errorf("find customers failed: %s", err.Error())
		}
		// 查询出错返回空记录, 忽略错误
		err = nil
		return
	}

	_customerList := []*protos.Customer{}
	for _, customerModel := range customerModels {
		customerPb := &protos.Customer{}
		contactList, err := c.getMyContacts(customerModel.ID)
		if err != nil {
			return customerList, nil
		}
		c.logger.Debugf("get customer's contacts: %+v\n", contactList)
		err = model2Pb(customerModel, customerPb)
		if err != nil {
			c.logger.Errorf("transform customer: %s model to protobuf object failed: %s", customerModel.Name, err.Error())
			return customerList, nil
		}
		customerPb.Contacts = contactList.List
		_customerList = append(_customerList, customerPb)
	}
	customerList.List = _customerList
	customerList.Count = count
	return
}

// Add ...
func (c *Customer) Add(customerPb *protos.Customer) (err error) {
	c.logger.Debug("add customer in Add()")
	customerModel := &model.Customer{}
	err = pb2Model(customerPb, customerModel)
	if err != nil {
		c.logger.Errorf("transform customer: %s protobuf object to model failed: %s", customerPb.Name, err.Error())
		return
	}
	tx := c.db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	// 创建customer记录
	err = tx.Create(customerModel).Error
	if err != nil {
		c.logger.Errorf("insert customer: %s into database failed: %s", customerPb.Name, err.Error())
		return
	}
	// 创建contacts记录, 这里为保证使用loraclient创建org与数据库中customer的一致性, 不调用contact服务进行创建.
	for _, contactPb := range customerPb.Contacts {
		contactModel := &model.Contact{
			CustomerID: customerModel.ID,
			Name:       contactPb.Name,
			Phone:      contactPb.Phone,
			Email:      contactPb.Email,
		}
		err = tx.Create(contactModel).Error
		if err != nil {
			c.logger.Errorf("create contact: %s for customer: %s failed: %s", contactPb.Name, customerModel.Name, err.Error())
			return
		}
	}

	loraclientCustomer := &protos.LoraclientCustormer{
		OrgName:        customerPb.Path,
		OrgDisplayName: customerPb.Title,
		UserName:       customerPb.Path,
		Passwd:         customerPb.Passwd2,
	}
	// 由于通过loraclient创建新客户需要同时创建org, user, deviceprofile, serviceprofile, 还初始化各种app, 所以耗时较长
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	resp, err := c.loraclientCli.AddCustmoer(ctx, loraclientCustomer)
	if err != nil {
		c.logger.Errorf("create customer: %s by loraclient failed: %s", customerPb.Name, err.Error())
		return
	}
	c.logger.Debug("add customer by loraclient success")
	customerModel.LoraOrgID = resp.OrgID
	customerModel.LoraUserID = resp.UserID
	err = tx.Save(customerModel).Error
	if err != nil {
		c.logger.Errorf("save lora org id: %d and user id: %d for customer: %s failed: %s", resp.OrgID, resp.UserID, customerPb.Name, err.Error())
		return
	}
	return
}

// Update ...
// 由于目前原型设计的问题, update操作需要先将customer的contact记录删除再重新创建
func (c *Customer) Update(customerPb *protos.Customer) (err error) {
	record := &model.Customer{}
	err = c.db.First(record, customerPb.ID).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回
		if err.Error() == "record not found" {
			err = nil
		} else {
			c.logger.Errorf("find customer: %s : %s", customerPb.Name, err.Error())
			return
		}
	}

	customerMap := map[string]interface{}{}
	err = pb2Map(customerPb, customerMap)
	if err != nil {
		c.logger.Errorf("transform customer: %s protobuf object to map failed: %s", customerPb.Name, err.Error())
		return
	}

	tx := c.db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = tx.Model(record).Update(customerMap).Error
	if err != nil {
		c.logger.Errorf("update customer: %s failed: %s", customerPb.Name, err.Error())
		return
	}
	// 移除原来的contacts
	err = tx.Where("customer_id = ?", customerPb.ID).Delete(&model.Contact{}).Error
	if err != nil {
		c.logger.Errorf("delete customer: %s's contacts failed: %s", record.Name, err.Error())
		return
	}

	// 创建新的contacts, 这里为保证使用loraclient创建org与数据库中customer的一致性, 不调用contact服务进行创建.
	for _, contactPb := range customerPb.Contacts {
		contactModel := &model.Contact{
			CustomerID: record.ID,
			Name:       contactPb.Name,
			Phone:      contactPb.Phone,
			Email:      contactPb.Email,
		}
		err = tx.Create(contactModel).Error
		if err != nil {
			c.logger.Errorf("create contact: %s for customer: %s failed: %s", contactPb.Name, record.Name, err.Error())
			return
		}
	}

	// 调用loraclient更新org及user
	req := &protos.LoraclientUpdateCustormerRequest{
		OrgID:          record.LoraOrgID,
		OrgName:        record.Path,
		OrgDisplayName: record.Title,
		UserID:         record.LoraUserID,
		UserName:       record.Path,
		Passwd:         record.Passwd2,
	}
	c.logger.Debugf("update customer by loraclient request: %+v", req)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err = c.loraclientCli.UpdateCustomer(ctx, req)
	if err != nil {
		c.logger.Errorf("update customer: %s by loraclient failed: %s", customerPb.Name, err.Error())
		return
	}

	return
}

// Delete ...
// 同时删除客户及其联系人记录, 这里不调用contact rpc接口, 因为不需要model与protobuf的类型转换, 可以直接完成...更方便
func (c *Customer) Delete(customerPb *protos.Customer) (err error) {
	c.logger.Debugf("delete customer: %d", customerPb.ID)

	record := &model.Customer{}
	err = c.db.First(record, customerPb.ID).Error
	if err != nil {
		// 注意: gorm把`record not found`当作错误返回,
		if err.Error() == "record not found" {
			c.logger.Errorf("record not found for customer: %d", customerPb.ID)
			err = nil
		} else {
			c.logger.Errorf("find customer: %d failed: %s", customerPb.ID, err.Error())
		}
		return
	}

	// 同时删除联系人和客户记录
	tx := c.db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	// 注意: 批量删除不能使用Find找到目标记录再调用Delete()删除记录列表, 只能用Where先过滤.
	// 而且不会出现record not found错误
	err = tx.Where("customer_id = ?", customerPb.ID).Delete(&model.Contact{}).Error
	if err != nil {
		c.logger.Errorf("delete customer: %s's contacts failed: %s", record.Name, err.Error())
		return
	}
	err = tx.Delete(record).Error
	if err != nil {
		c.logger.Errorf("delete customer: %s failed: %s", record.Name, err.Error())
		return
	}

	// 然后调用loraclient删除对应的org和user, 如果record不存在, 则这一步就无意义
	req := &protos.LoraclientDeleteCustomerRequest{
		OrgID:  record.LoraOrgID,
		UserID: record.LoraUserID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err = c.loraclientCli.DeleteCustmoer(ctx, req)
	if err != nil {
		c.logger.Errorf("delete customer: %s by loraclient failed: %s", record.Name, err.Error())
	}
	return
}

// getMyContacts 调用contact相关函数获得指定客户下的所有联系人
// 其实使用gorm的关联查询也可以得到, 但那是model对象, 还需要将其转化成protobuf对象才能一起返回给调用者使用, 不如直接调用contact服务获得.
func (c *Customer) getMyContacts(customerID uint64) (contactList *protos.ContactList, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	getContactsRequest := &protos.GetContactsRequest{
		CustomerID: customerID,
	}
	contactList, err = c.contactCli.GetList(ctx, getContactsRequest)
	if err != nil {
		c.logger.Errorf("find customer: %d's contact list failed: %s", customerID, err.Error())
		return
	}
	return
}

func pb2Model(pb *protos.Customer, model *model.Customer) (err error) {
	model.ID = pb.ID
	model.Name = pb.Name
	model.Title = pb.Title
	model.Address = pb.Address
	model.Path = pb.Path
	model.Passwd1 = pb.Passwd1
	model.Passwd2 = pb.Passwd2

	return
}

func pb2Map(pb *protos.Customer, theMap map[string]interface{}) (err error) {
	theMap["id"] = pb.ID
	theMap["name"] = pb.Name
	theMap["title"] = pb.Title
	theMap["address"] = pb.Address
	theMap["path"] = pb.Path
	theMap["enable"] = pb.Enable

	return
}

// model2Pb 一般用于Get(List)操作
func model2Pb(model *model.Customer, pb *protos.Customer) (err error) {
	pb.ID = model.ID
	pb.Name = model.Name
	pb.Title = model.Title
	pb.Address = model.Address
	pb.Path = model.Path
	pb.Enable = model.Enable

	return
}
