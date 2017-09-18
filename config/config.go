package config

import (
	"encoding/json"
	"io/ioutil"
)

type GlobalConfig struct {
	HostAddr     string `json:"host"`
	Port         int    `json:"port"`
	LoginBotName string `json:"botName"`
	LogOath      string `json:"oath"`
}

type UserConfig struct {
	GlobalConfig
	Channel string `json:"channel"`
}

func LoadUserConfig(userDecoder *json.Decoder) (*UserConfig, error) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	gc := GlobalConfig{}
	err = json.Unmarshal(file, &gc)
	uc := UserConfig{gc, ""}
	err = userDecoder.Decode(&uc)
	if err != nil {
		return nil, err
	}
	return &uc, nil
}
