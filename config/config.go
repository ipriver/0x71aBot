package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	HostAddr  string
	Port      int
	LoginName string
	LogOath   string
	Channel   string
}

func LoadConfig() (*Config, error) {
	file, err := ioutil.ReadFile("../config.json")
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(file, config)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}
	return config, nil
}
