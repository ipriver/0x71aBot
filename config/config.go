package config

import (
	"encoding/json"
	"io/ioutil"
)

var err error

//main configuration data structure which is parsed from config.json
type GlobalConfig struct {
	HostAddr     string `json:"host"`
	Port         int    `json:"port"`
	LoginBotName string `json:"botName"`
	LogOath      string `json:"oath"`
}

//loads data from config.json into the structure
func (gc *GlobalConfig) Load() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &gc)
	if err != nil {
		panic(err)
	}
}

//saves new data into config.json file
func (gc *GlobalConfig) Save() error {
	js, _ := json.Marshal(gc)
	err = ioutil.WriteFile("config.json", js, 0644)
	if err != nil {
		return err
	}
	return nil
}

//structure is used by Bot for creating new Bots stucts
type UserConfig struct {
	GlobalConfig
	Channel string `json:"channel"`
}

//creates UserConfig struct and returns it
func LoadUserConfig(channel string) (*UserConfig, error) {
	gc := GlobalConfig{}
	gc.Load()
	uc := UserConfig{gc, channel}
	if err != nil {
		return nil, err
	}
	return &uc, nil
}
