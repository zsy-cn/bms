package media

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/zsy-cn/bms/util/log"
)

// GetDuration 获取指定音频文件的时长
func GetDuration(logger log.Logger, audioPath string) (duration float64, err error) {
	var stdout bytes.Buffer

	cmd := exec.Command("ffprobe", "-loglevel", "quiet", "-print_format", "json", "-show_format", audioPath)
	cmd.Stdout = &stdout

	err = cmd.Run()
	if err != nil {
		logger.Errorf("execute ffprobe for %s error in getDuration(): %s", audioPath, err.Error())
		return
	}

	buf := make([]byte, 10240)
	len, err := stdout.Read(buf)
	if err != nil {
		logger.Errorf("read ffprobe stdout for %s error in getDuration(): %s", audioPath, err.Error())
		return
	}
	_buf := buf[:len]
	infoMap := map[string]interface{}{}
	err = json.Unmarshal(_buf, &infoMap)
	if err != nil {
		logger.Errorf("unmarshal ffprobe stdout for %s error in getDuration(): %s", audioPath, err.Error())
		return
	}

	format := infoMap["format"].(map[string]interface{})
	durationStr := format["duration"].(string)
	duration, err = strconv.ParseFloat(durationStr, 64)
	if err != nil {
		logger.Errorf("parse duration into float for %s error in getDuration(): %s", audioPath, err.Error())
		return
	}
	return
}

// FormatDuration 格式化浮点类型的duration值为00:03:45形式, 或是3分45秒形式
func FormatDuration(value float64) (durationStr string) {
	rawDuration := fmt.Sprintf("%ds", int(value))
	durationObj, _ := time.ParseDuration(rawDuration)

	durationStr = durationObj.String()

	// 返回结果可有两种格式
	durationStr = strings.Replace(durationStr, "h", "小时", -1)
	durationStr = strings.Replace(durationStr, "m", "分", -1)
	durationStr = strings.Replace(durationStr, "s", "秒", -1)
	fmt.Println(durationStr)

	/*
		date, _ := time.Parse("2006-01-02", "2000-01-01")
		newDate := date.Add(durationObj)
		durationStr = newDate.Format("15:04:05")
		fmt.Println(durationStr)
	*/
	return
}
