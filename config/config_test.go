package config

import (
	"reflect"
	"testing"
)

func TestGlobalConfig(t *testing.T) {
	c := GlobalConfig{}
	rc := reflect.TypeOf(c)
	if rc.Name() != "GlobalConfig" {
		t.Error("incorrect struct name")
	}
	host := "host"
	port := 8080
	login := "login"
	oath := "oath18"
	c = GlobalConfig{host, port, login, oath}
	if c.GetHost() != host || reflect.TypeOf(c.GetHost()).String() != "string" {
		t.Error("error in GetHost")
	}
	if c.GetPort() != port || reflect.TypeOf(c.GetPort()).String() != "int" {
		t.Error("error in GetPort")

	}
	if c.GetLoginBotName() != login || reflect.TypeOf(c.GetLoginBotName()).String() != "string" {
		t.Error("error in GetLoginBotName")

	}
	if c.GetOath() != oath || reflect.TypeOf(c.GetOath()).String() != "string" {
		t.Error("error in GetOath")

	}
}

func TestGlobalConfigThrowingErrorIfFileDoesntexist(t *testing.T) {
	//this should work
	gc := GlobalConfig{}
	configFile = "../config.json"
	err := gc.Load()
	if err != nil {
		t.Error("shouldn't return error")
	}
	//this should throw error
	configFile = "../randomFile.json"
	err = gc.Load()
	if err == nil {
		t.Error("should return error")
	}
}

func TestUserConfig(t *testing.T) {
	c := UserConfig{}
	rc := reflect.TypeOf(c)
	host := "host"
	port := 8080
	login := "login"
	oath := "oath18"
	channel := "test"
	c = UserConfig{GlobalConfig{host, port, login, oath}, channel}
	if rc.Name() != "UserConfig" {
		t.Error("incorrect struct name")
	}
	if c.GetHost() != host || reflect.TypeOf(c.GetHost()).String() != "string" {
		t.Error("error in GetHost")
	}
	if c.GetPort() != port || reflect.TypeOf(c.GetPort()).String() != "int" {
		t.Error("error in GetPort")

	}
	if c.GetLoginBotName() != login || reflect.TypeOf(c.GetLoginBotName()).String() != "string" {
		t.Error("error in GetLoginBotName")

	}
	if c.GetOath() != oath || reflect.TypeOf(c.GetOath()).String() != "string" {
		t.Error("error in GetOath")
	}
	if c.GetChannel() != channel || reflect.TypeOf(c.GetChannel()).String() != "string" {
		t.Error("error in GetChannel")
	}
}
