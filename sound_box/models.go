package sound_box

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/model"
)

var Tables = []interface{}{
	&SoundBox{},
	&SoundBoxMedia{},
	&SoundBoxCronjob{},
}

func Migrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(Tables...).Error
	if err != nil {
		panic(err)
	}

	return
}

// SoundBox 音箱表
type SoundBox struct {
	model.Base
	DeviceSN string
	GroupID  uint64
	SQRCode  string `gorm:"size:32;not null;column:sqr_code"`
	SeedCode string `gorm:"size:32;not null"`
	Volume   uint8  // 音量, 1-100
}

type SoundBoxInfo struct {
	Kind   int32  `json:"kind,omitempty"`
	Sn     string `json:"sn,omitempty"`
	Status int32  `json:"status,omitempty"`
	Vol    int32  `json:"vol,omitempty"`
	LIp    string `json:"l_ip,omitempty"`
}

type SoundBoxInfoList struct {
	List []*SoundBoxInfo `json:"list"`
	Res  bool            `json:"res"`
}

// SoundBoxMedia 媒体表
// 不同用户使用不同的目录前缀.
// 注意: media作为媒体时为不可数名词, 所以在数据库创建的表名为: sound_box_media, 没有s
type SoundBoxMedia struct {
	model.Base
	Name       string `gorm:"not null"`
	CustomerID uint64 `gorm:"not null"`
	Duration   float64
	Size       uint64
	Path       string
}

// SoundBoxCronjob 定时任务表
type SoundBoxCronjob struct {
	model.Base
	Name           string `gorm:"not null"`
	CustomerID     uint64 `gorm:"not null"`
	StartAt        string `gorm:"not null"` // 格式 12:00
	StopAt         string `gorm:"not null"` // 格式 12:00
	RepeatAt       string `gorm:"not null"` // 一周的哪几天播放
	InvolvedGroups string `gorm:"not null"`
	InvolvedMedias string `gorm:"not null"`
	Enable         bool   `gorm:"not null; default: true"`
}
