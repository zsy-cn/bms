package internal

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/zsy-cn/bms/environ_monitor"
)

// GetStandardWeatherInfo 从墨迹天气获取标准环境数据
func (ss *DefaultEnvironMonitorService) GetStandardWeatherInfo(url string) (msg *environ_monitor.EnvironMonitor, err error) {
	dom, err := goquery.NewDocument(url)
	if err != nil {
		ss.l.Errorf("generate dom failed in GetStandardWeatherInfo(): %s", err.Error())
		return
	}
	pm025URL, _ := dom.Find("div.wea_alert > ul > li > a").Attr("href")
	dom2, err := goquery.NewDocument(pm025URL)
	if err != nil {
		ss.l.Errorf("generate pm2.5 dom failed in GetStandardWeatherInfo(): %s", err.Error())
		return
	}
	pm025Str := dom2.Find("div.aqi_info_item > ul > li:nth-child(2) > span").Text()
	temperatureStr := dom.Find("div.wea_weather > em").Text()
	humidityInfoStr := dom.Find("div.wea_about > span").Text()
	humidityInfoStr = strings.Split(humidityInfoStr, " ")[1]
	humidityStr := humidityInfoStr[:len(humidityInfoStr)-1]

	pm025, err := strconv.ParseFloat(pm025Str, 64)
	if err != nil {
		ss.l.Errorf("parse pm2.5 value failed in GetStandardWeatherInfo(): %s", err.Error())
		return
	}
	temperature, err := strconv.ParseFloat(temperatureStr, 64)
	if err != nil {
		ss.l.Errorf("parse temperature value failed in GetStandardWeatherInfo(): %s", err.Error())
		return
	}
	humidity, err := strconv.ParseFloat(humidityStr, 64)
	if err != nil {
		ss.l.Errorf("parse humidity value failed in GetStandardWeatherInfo(): %s", err.Error())
		return
	}

	msg = &environ_monitor.EnvironMonitor{
		Temperature: temperature,
		Humidity:    humidity,
		PM025:       pm025,
	}
	return
}
