package main

import (
	"os"

	"github.com/zsy-cn/bms/loraclient/service"
	"github.com/zsy-cn/bms/loraclient/transport"
	"github.com/zsy-cn/bms/util/log"
)

var logger = log.NewLogger(os.Stdout)

func main() {
	logger.SetLevel("debug")

	logger.Info("starting loraclient service")
	var loraclientServ service.ILoraclient
	loraclientServ, err := service.New(logger)
	if err != nil {
		logger.Errorf("starting loraclient service failed: %s", err.Error())
	}
	logger.Info("starting loraclient grpc server")
	transport.StartGrpcTransport(loraclientServ)
}
