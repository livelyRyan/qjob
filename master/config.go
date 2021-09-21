package master

import (
	"encoding/json"
	"io/ioutil"
)

type GlobalConfig struct {
	ApiPort         int      `json:"apiPort"`
	ApiReadTimeout  int      `json:"apiReadTimeout"`
	ApiWriteTimeout int      `json:"apiWriteTimeout"`
	EtcdEndpoints   []string `json:"etcdEndpoints"`
	EtcdDialTimeout int      `json:"etcdDialTimeout"`
	StaticDir       string   `json:"staticDir"`
}

var Config *GlobalConfig

func InitGlobalConfig(configPath string) error {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	var config GlobalConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}
	Config = &config
	return nil
}
