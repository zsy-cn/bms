package parserhub

import (
	"context"
	"time"

	"github.com/henrylee2cn/faygo"
	"github.com/mitchellh/mapstructure"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/server/controller"
)

// UplinkHandler ...
func UplinkHandler(ctx *faygo.Context) (err error) {
	logger.Debug("request ParserHub UplinkHandler controller")
	result := controller.NewResult()
	defer ctx.JSON(200, result, true)

	var jsonparams map[string]interface{}
	// 从Body中得到json形式的参数(因为没有`<in:body>`标签, 这里要直接从body中提取数据)
	ctx.BindJSON(&jsonparams)
	logger.Debugf("uplink message: %+v\n", jsonparams)

	uplinkMsg := &protos.ParserHubUplinkMsg{}
	err = mapstructure.Decode(jsonparams, uplinkMsg)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		logger.Errorf("decode uplink message failed: %s", err.Error())
		return
	}

	cli, exist := controller.ServiceMap.Get("parserhub_cli")
	if !exist {
		logger.Info("Trying to get parserhub service failed, grpc was not connected")
		return
	}

	_ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = cli.(protos.ParserHubServiceClient).ParseAndSave(_ctx, uplinkMsg)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		logger.Errorf("parser and save uplink message failed: %s", err.Error())
		return
	}

	return
}
