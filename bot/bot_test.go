package bot

import (
	"reflect"
	"testing"
)

func TestBot(t *testing.T) {
	b := Bot{}
	bc := reflect.TypeOf(b)
	if bc.Name() != "Bot" {
		t.Error("incorrect struct name")
	}
	b = Bot{}
	if b.GetId()
}

func TestBotConstructor(t *testing.T) {
	b := Bot{}
	b.Constructor(22, "test")
	if b.GetId()
}
