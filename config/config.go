package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var err error

type GlobalConfig struct {
	HostAddr     string `json:"host"`
	Port         int    `json:"port"`
	LoginBotName string `json:"botName"`
	LogOath      string `json:"oath"`
}

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

func (gc *GlobalConfig) Save() error {
	js, _ := json.Marshal(gc)
	err = ioutil.WriteFile("config.json", js, 0644)
	if err != nil {
		return err
	}
	return nil
}

type UserConfig struct {
	GlobalConfig
	Channel string `json:"channel"`
}

func LoadUserConfig(channel string) (*UserConfig, error) {
	//Update with Load function
	gc := GlobalConfig{}
	gc.Load()
	uc := UserConfig{gc, channel}
	if err != nil {
		return nil, err
	}
	return &uc, nil
}
