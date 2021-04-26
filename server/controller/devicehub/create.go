package devicehub

import (
	"context"
	"encoding/json"
	"time"

	"github.com/henrylee2cn/faygo"
	"github.com/mitchellh/mapstructure"

	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/server/controller"
)

// CreateParams ...
type CreateParams struct {
	// <in:body>只能有一个标签, faygo会将请求体格式化为map类型
	Body map[string]interface{} `param:"<in:body> <required>"`
}

// Serve ...
func (p *CreateParams) Serve(ctx *faygo.Context) (err error) {
	logger.Debug("request device create controller")
	result := controller.NewResult()
	defer ctx.JSON(200, result, true)

	devicePb := &protos.Device{}
	body := p.Body

	err = mapstructure.Decode(body, devicePb)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		return
	}

	// 没有找到直接从faygo得到body的字节对象, 只能用marshal序列化.
	// extraInfo由device服务根据设备类型分别处理
	extraInfo, err := json.Marshal(body)
	if err != nil {
		logger.Errorf("marshal body failed: %s", err.Error())
		result.Code = -1
		result.Msg = err.Error()
		return
	}
	devicePb.ExtraInfo = &protos.ExtraDeviceInfo{
		Info: string(extraInfo),
	}
	cli, exist := controller.ServiceMap.Get("device_cli")
	if !exist {
		logger.Warn("Trying to create device failed, grpc was not connected")
		result.Code = -1
		result.Msg = ErrDeviceHubServiceDisconnected.Error()
		return
	}

	_ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = cli.(protos.DeviceServiceClient).Add(_ctx, devicePb)
	if err != nil {
		logger.Errorf("add device by devicehub service failed: %s", err.Error())
		result.Code = -1
		result.Msg = err.Error()
		return
	}
	return
}
