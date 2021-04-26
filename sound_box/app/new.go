package app

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/sound_box"
	"github.com/zsy-cn/bms/sound_box/internal"
	"github.com/zsy-cn/bms/util/log"
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("anning")
	// 合法的环境变量只能包含下划线_, 不能包含中横线或点号
	// replacer用于将目标key转换成合法的环境变量字符串格式
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetDefault("mediacore-addr", conf.MediaCoreServiceAddr+conf.MediaCoreServicePort)
}

func NewSoundBoxService(
	l log.Logger,
	db *gorm.DB,
	cfg sound_box.SoundBoxConfig,
) (sound_box.SoundBoxService, error) {
	return internal.NewSoundBoxService(l, db, cfg)
}
