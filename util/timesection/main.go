package timesection

import (
	"time"

	"github.com/apex/log"
)

// GetTimeSection 根据dateFrom, dateTo两个参数返回实际的开始与结束时间
// dateFrom与dateTo格式都为2019-02-24的字符串
func GetTimeSection(dateFrom, dateTo string) (startDate time.Time, endDate time.Time, now time.Time, err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now = time.Now()
	todayStr := now.In(loc).Format("2006-01-02")
	oneDay, _ := time.ParseDuration("24h")

	if dateFrom == "" {
		dateFrom = todayStr
	}
	startDate, err = time.ParseInLocation("2006-01-02", dateFrom, loc)
	if err != nil {
		log.Errorf("parse in location failed in GetEnvironMonitorSectionAverageData(): %s", err.Error())
		return startDate, endDate, now, err
	}

	if dateTo == "" {
		dateTo = todayStr
	}
	// 实际的endDate应该是 < endDate+1, 否则endDate当天的记录不会出现在结果中, 需要注意一下
	endDate, err = time.ParseInLocation("2006-01-02", dateTo, loc)
	if err != nil {
		log.Errorf("parse in location failed in GetEnvironMonitorSectionAverageData(): %s", err.Error())
		return startDate, endDate, now, err
	}
	endDate = endDate.Add(oneDay)
	/*
		// 如果endDate选择的日期已经是在今日之后, 则截止日期只能是现在
		if endDate.After(now) {
			endDate = now
		}
	*/
	return
}

// GetTimeDeadline 获取指定时间前的deadline时间对象
func GetTimeDeadline(offlineTimeout string) (deadline time.Time) {
	now := time.Now()
	d2, _ := time.ParseDuration(offlineTimeout)
	deadline = now.Add(d2)
	return deadline
}
