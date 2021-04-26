package env_yingfeichi_01

import (
	"github.com/zsy-cn/bms/protos"
)

/*
以lora-app-server发来的两个包为例
[2 1 3 136 0 73 3 183 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255]
[2 2 255 255 255 255 3 42 255 255 255 255 255 255 255 255 255 255 255 255 4 46 255 255 255 255 255 255 255 255 255 255 255 255]

两个包的第一个字节表示包总量, 2表示这一条完成的消息包含2个包.
第二个字节表示当前包的序号, 一个为1, 一个为2.
第一个包的第3, 4个字节为设备号, 好像没什么用.
之后的就是数据字段了.
*/

// MsgHandler ...
func (device EnvYingfeichi01) MsgHandler() {
	envDataStore := map[string]map[string]interface{}{}
	var sum int
	var err error
	for uplinkMsg := range device.msgChan {
		device.logger.Debug("have receive Msg from msgChan")
		if _, exist := envDataStore[uplinkMsg.DevEUI]; !exist {
			envDataStore[uplinkMsg.DevEUI] = map[string]interface{}{
				"locked": false,
				"data":   map[int][]byte{},
			}
		}
		theMap := envDataStore[uplinkMsg.DevEUI]
		lockedField := theMap["locked"].(bool)
		dataField := theMap["data"].(map[int][]byte)
		// 由于环境监测传感器会将同一个包重复发好几次, 在得到完整的包后, 就忽略剩余的包并开始处理
		if lockedField {
			device.logger.Debug("drop this Msg from msgChan")
			continue
		}
		sum = int(uplinkMsg.FinalData[0])
		order := int(uplinkMsg.FinalData[1])
		if order == 1 {
			dataField[order] = uplinkMsg.FinalData[4:]
		} else {
			dataField[order] = uplinkMsg.FinalData[2:]
		}
		// 如果本次上行数据包已经全部接收完成, 就开始解析
		if len(dataField) == sum {
			// 首先锁住当前设备的map, 丢弃之后重复的包
			lockedField = true
			// 然后合并数据
			var fullData []byte
			for i := 1; i <= sum; i++ {
				// 切片合并, 注意最后的三个点
				fullData = append(fullData, dataField[i]...)
				delete(dataField, i) // 移除此设备在map映射中的数据(但不移除键, 只移除值)
			}
			fullMsg := &protos.ParserHubUplinkMsg{
				FPort:     uplinkMsg.FPort,
				DevEUI:    uplinkMsg.DevEUI,
				FinalData: fullData,
			}
			err = device.decode(fullMsg)
			if err != nil {
				device.logger.Errorf("full uplink msg decode failed: %s", err.Error())
			}
			// 解锁, 清零
			lockedField = false
			sum = 0
		}
		device.logger.Debug("handle receive Msg from msgChan finished")
	}
}
