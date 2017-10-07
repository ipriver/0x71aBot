package config

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
	"io/ioutil"
)

const configFilePath string = "config.json"

type Configer interface {
	Load()
	Save() error
}

var Config *GlobalConfig

//main configuration data structure which is parsed from config.json
type GlobalConfig struct {
	HostAddr     string `json:"host"`
	Port         int    `json:"port"`
	LoginBotName string `json:"botName"`
	LogOath      string `json:"oath"`
	Rc           redis.Conn
	Db           *sql.DB
}

func (gc *GlobalConfig) ConnectToSQL() {
	const (
		DB_USER = "ipriver"
		DB_PASS = "root"
		DB_NAME = "twitch_bot"
	)

	gc.Db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASS, DB_NAME))
	if err != nil {
		panic(err)
	}
}

func (gc *GlobalConfig) ConnectToNOSQL() {
	Rc, err = redis.DialURL("redis://user:@localhost:6379/0")
	if err != nil {
		panic(err)
	}
}

//loads data from GlobalConfig.json into the structure
func (gc *GlobalConfig) Load() error {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &gc)
	if err != nil {
		return err
	}
	gc.ConnectToSQL()
	gc.ConnectToNOSQL()
	return err
}

//saves new data into GlobalConfig.json file
func (gc *GlobalConfig) Save() error {
	js, _ := json.Marshal(gc)
	err := ioutil.WriteFile("config.json", js, 0644)
	if err != nil {
		return err
	}
	return err
}

func init() {
	Config = new(GlobalConfig)
	err := Config.Load()
	if err != nil {
		panic(err)
	}
}
