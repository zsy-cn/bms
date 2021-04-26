package env_yingfeichi_01

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/environ_monitor"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/parser"
	"github.com/zsy-cn/bms/util/log"
)

// EnvYingfeichi01 ...
type EnvYingfeichi01 struct {
	parser.ConsulParserService
	logger  *log.Logger
	db      *gorm.DB
	msgChan chan *protos.ParserHubUplinkMsg
}

// New ...
func New(logger *log.Logger) (device *EnvYingfeichi01, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	// 这里暂时将分块的上行信息存储到channel, 之后可以考虑存到redis中.
	msgChan := make(chan *protos.ParserHubUplinkMsg, 10000)
	device = &EnvYingfeichi01{
		logger:  logger,
		db:      db,
		msgChan: msgChan,
	}
	// 循环处理
	go device.MsgHandler()
	return
}

// Decode ...
// 要实现一个interface, receiver不能是指针???(在这里是EnvYingfeichi01实现api.go->Parser接口)
// uplinkMsg中FinalData已经经过base64解码, 无需重复操作, 直接解析即可
func (device EnvYingfeichi01) Decode(uplinkMsg *protos.ParserHubUplinkMsg) (err error) {
	device.logger.Debugf("receive uplink mesage in Decode(): %+v", uplinkMsg)
	device.logger.Debug("uplink mesage payload bytes in Decode(): ", uplinkMsg.FinalData)
	device.msgChan <- uplinkMsg
	device.logger.Debug("have send Msg to device msgChan")
	return
}

func (device EnvYingfeichi01) decode(uplinkMsg *protos.ParserHubUplinkMsg) (err error) {
	device.logger.Debugf("start decode full mesage in decode(): %+v", uplinkMsg)
	uplinkMsgPayload := &model.UplinkMsgPayload{
		Attributes: model.UplinkMsgAttribute{},
	}
	var intData []int
	// 把字节数组中的值全转换成int类型...因为byte类型不能进行数值计算, 不像js
	for _, b := range uplinkMsg.FinalData {
		intData = append(intData, int(b))
	}
	// for i := 0; i < len(intData)/2; i++ {
	for i := 0; i < 29; i++ {
		taskList4WeichuanENV01[i](intData, uplinkMsgPayload)
	}
	// 解析完成后入库
	baseAppModel := &model.App{}
	err = parser.FillBaseAppModel(device.db, uplinkMsg, uplinkMsgPayload, baseAppModel)
	if err != nil {
		// 错误日志在FillBaseAppModel函数内部打印过了
		return
	}
	temperature := uplinkMsgPayload.Attributes["Temperature"]
	pm025 := uplinkMsgPayload.Attributes["PM025"]
	noise := uplinkMsgPayload.Attributes["Noise"]
	humidity := uplinkMsgPayload.Attributes["Humidity"]
	// map类型, 不一定有此字段
	// windSpeed := uplinkMsgPayload.Attributes["WindSpeed"]
	// airPressure := uplinkMsgPayload.Attributes["AirPressure"]
	envModel := &environ_monitor.EnvironMonitor{
		App:         *baseAppModel, // 引用类型
		Temperature: temperature.Value,
		PM025:       pm025.Value,
		Noise:       noise.Value,
		Humidity:    humidity.Value,
		// AirSpeed:    windSpeed.Value,
		// Pressure:    airPressure.Value,
	}
	err = device.db.Create(envModel).Error
	if err != nil {
		device.logger.Errorf("create environ monitor app message record failed in Decode(): %s", err.Error())
		return
	}
	return
}
