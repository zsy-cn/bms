package sql

import (
	"fmt"

	"github.com/zsy-cn/bms/protos"
)

// MakeWhereStr ...
// 注意main_tbl为devices表, 因为其他表没有serial_number列
func MakeWhereStr(req *protos.GetDevicesRequestForCustomer) (whereStr string) {
	whereStr = " where 1=1 "
	if req.SerialNumber != "" {
		whereStr = fmt.Sprintf(whereStr+" and main_tbl.serial_number = '%s' ", req.SerialNumber)
	}
	if req.DeviceTypeID != 0 {
		whereStr = fmt.Sprintf(whereStr+" and main_tbl.device_type_id = %d ", req.DeviceTypeID)
	}
	if req.CustomerID != 0 {
		whereStr = fmt.Sprintf(whereStr+" and main_tbl.customer_id = %d ", req.CustomerID)
	}
	if req.GroupID != 0 {
		whereStr = fmt.Sprintf(whereStr+" and main_tbl.group_id = %d ", req.GroupID)
	}
	return
}
