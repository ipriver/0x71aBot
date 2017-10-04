package config

import (
	"encoding/json"
	"io/ioutil"
)

import "fmt"

var configFile string = "config.json"
var err error

type Configer interface {
	Load()
	Save() error
}

//main configuration data structure which is parsed from config.json
type GlobalConfig struct {
	HostAddr     string `json:"host"`
	Port         int    `json:"port"`
	LoginBotName string `json:"botName"`
	LogOath      string `json:"oath"`
}

func (gc *GlobalConfig) GetHost() string {
	return gc.HostAddr
}

func (gc *GlobalConfig) GetPort() int {
	return gc.Port
}

func (gc *GlobalConfig) GetLoginBotName() string {
	return gc.LoginBotName
}

func (gc *GlobalConfig) GetOath() string {
	return gc.LogOath
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
	GlobalConfig
	Channel string `json:"channel"`
}

func (uc *UserConfig) GetChannel() string {
	return uc.Channel
}

func (uc *UserConfig) SetChannel(channel string) {
	uc.Channel = channel
}

//loads data into struct
func (uc *UserConfig) Load() {
	gc := GlobalConfig{}
	gc.Load()
	uc.GlobalConfig = gc
	fmt.Println(gc)
}
