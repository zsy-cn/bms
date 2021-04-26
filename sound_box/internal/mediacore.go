package internal

import (
	"encoding/json"
	"io"
	"net"
	"strings"

	"github.com/spf13/viper"
	"github.com/zsy-cn/bms/sound_box"
	"github.com/zsy-cn/bms/util/log"
)

type SoundBoxDeviceStatus struct {
	StatusMap map[string]int8
	Online    int
	Offline   int
	Playing   int
	Unknown   int
}

func getMediaCoreStatus(log log.Logger) (soundboxInfoList *sound_box.SoundBoxInfoList, err error) {
	soundboxInfoList = &sound_box.SoundBoxInfoList{
		List: []*sound_box.SoundBoxInfo{},
		Res:  false,
	}
	conn, err := net.Dial("tcp", viper.GetString("mediacore-addr"))
	if err != nil {
		log.Errorf("connect to mediacore server failed in getSoundBoxStatus(): %s", err.Error())
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(`{"mode":1001}`))
	if err != nil {
		log.Errorf("send request to mediacore server failed in getSoundBoxStatus(): %s", err.Error())
		return
	}
	allbuf := make([]byte, 0)
	buf := make([]byte, 1024)
	length := 0

	for {
		n, err := conn.Read(buf)
		if n > 0 {
			length += n
			allbuf = append(allbuf, buf[:n]...)
		}
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			} else {
				log.Debug("read complete")
			}
			break
		}
	}
	jsonStr := string(allbuf[:length])
	payload := strings.Split(jsonStr, "\n")[1]

	err = json.Unmarshal([]byte(payload), soundboxInfoList)
	if err != nil {
		log.Errorf("unmarshal mediacore server response failed in getSoundBoxStatus(): %s", err.Error())
		return
	}
	return
}

// 获取指定seed码的设备在线情况
func getSoundBoxDevicesStatus(log log.Logger, seedCodes []string) (status *SoundBoxDeviceStatus, err error) {
	soundboxInfoList, err := getMediaCoreStatus(log)
	if err != nil {
		log.Errorf("get media core status failed in getSoundBoxStatus(): %s", err.Error())
		return
	}
	status = &SoundBoxDeviceStatus{
		StatusMap: map[string]int8{},
	}
	for _, seedCode := range seedCodes {
		for _, soundboxInfo := range soundboxInfoList.List {
			if soundboxInfo.Sn == seedCode {
				if soundboxInfo.Status == 0 { // 离线
					status.StatusMap[seedCode] = 2
				} else if soundboxInfo.Status == 1 { // 在线
					status.StatusMap[seedCode] = 1
				} else if soundboxInfo.Status == 2 { // 播放中
					status.StatusMap[seedCode] = 3
				}
				break
			}
		}
		if _, ok := status.StatusMap[seedCode]; !ok {
			status.StatusMap[seedCode] = -1
		}
	}
	return
}

func getSoundBoxStatus(log log.Logger, seedCode string) (code int8, err error) {
	status, err := getSoundBoxDevicesStatus(log, []string{seedCode})
	code = status.StatusMap[seedCode]
	return
}
