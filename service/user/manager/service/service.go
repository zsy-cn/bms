package service

import (
	"github.com/jinzhu/gorm"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

// Manager ...
type Manager struct {
	logger *log.Logger
	db     *gorm.DB
}

// New 实例化Manager服务对象
func New(logger *log.Logger) (manager *Manager, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}

	manager = &Manager{
		logger: logger,
		db:     db,
	}
	return
}

// Login ...
func (c *Manager) Login(req *protos.ManagerLoginRequest) (resp *protos.ManagerLoginResponse, err error) {
	username := req.UserName
	passwd := req.Password

	managerWhereArgs := map[string]interface{}{
		"name":        username,
		"passwd":      passwd,
		"enable":      "true",
		"role_enable": "true",
	}
	managerModel := &model.Manager{}
	err = c.db.Where(managerWhereArgs).First(managerModel).Error
	if err != nil {
		// 不管是不是record not found, 都要返回err
		if err.Error() != "record not found" {
			c.logger.Errorf("find manager: %s failed: %s", req.UserName, err.Error())
		}
		return
	}

	resp = &protos.ManagerLoginResponse{
		ID:     managerModel.RoleID,
		Name:   managerModel.Name,
		RoleID: managerModel.RoleID,
	}

	return
}
