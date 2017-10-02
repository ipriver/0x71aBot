package config

import (
	"encoding/json"
	"io/ioutil"
)

var configFile string = "config.json"
var err error

type Configer interface {
	Load()
	Save() error
}

//main configuration data structure which is parsed from config.json
type GlobalConfig struct {
	hostAddr     string `json:"host"`
	port         int    `json:"port"`
	loginBotName string `json:"botName"`
	logOath      string `json:"oath"`
}

func (gc *GlobalConfig) GetHost() string {
	return gc.hostAddr
}

func (gc *GlobalConfig) GetPort() int {
	return gc.port
}

func (gc *GlobalConfig) GetLoginBotName() string {
	return gc.loginBotName
}

func (gc *GlobalConfig) GetOath() string {
	return gc.logOath
}

//loads data from config.json into the structure
func (gc *GlobalConfig) Load() error {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &gc)
	if err != nil {
		return err
	}
	return nil
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
	channel string `json:"channel"`
	GlobalConfig
}

func (uc *UserConfig) GetChannel() string {
	return uc.channel
}

//loads data into struct
func (uc *UserConfig) Load() error {
	gc := GlobalConfig{}
	gc.Load()
	uc.GlobalConfig = gc
}
