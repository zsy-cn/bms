package parser

import (
	"github.com/zsy-cn/bms/protos"
)

// Parser 上行信息解析器
type Parser interface {
	Decode(uplinkMsg *protos.ParserHubUplinkMsg) (err error)
	HealthCheck(empty *protos.Empty) (resp *protos.ParserHealthCheckResponse, err error)
}
