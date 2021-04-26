package devicehub

import (
	"errors"
	"os"

	"github.com/zsy-cn/bms/util/log"
)

var logger = log.NewLogger(os.Stdout)

// ErrDeviceHubServiceDisconnected devicehub服务未连接
var ErrDeviceHubServiceDisconnected = errors.New("devicehub service is disconnected")
