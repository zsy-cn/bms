package internal

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

func (cs *DefaultCameraService) readConfigFile() (cfgMap map[string]string, err error) {
	cfgData, err := ioutil.ReadFile(cs.cfgPath)
	if err != nil {
		cs.l.Errorf("read config file failed: %s", err.Error())
		return
	}

	cfgMap = make(map[string]string)
	err = yaml.Unmarshal(cfgData, &cfgMap)
	if err != nil {
		cs.l.Errorf("unmarshal config data failed: %s", err.Error())
		return
	}
	return cfgMap, nil
}

func (cs *DefaultCameraService) writeConfigFile(key, val string) (err error) {
	cfgMap, err := cs.readConfigFile()
	if err != nil {
		// 日志已经打印过
		return
	}
	cfgMap[key] = val

	newCfgData, err := yaml.Marshal(&cfgMap)
	if err != nil {
		cs.l.Errorf("marshal config data failed: %s", err.Error())
		return
	}

	err = ioutil.WriteFile(cs.cfgPath, newCfgData, 0644)
	if err != nil {
		cs.l.Errorf("write config file failed: %s", err.Error())
		return
	}
	return
}
