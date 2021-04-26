package message_management

import (
	"github.com/zsy-cn/bms/protos"
)

type MessageManagementService interface {
	DiscardNotification(customerID uint64, deviceSN string) (err error)
	GetNotifications(req *protos.GetDevicesRequestForCustomer) (notifications *protos.NotificationList, err error)
}
