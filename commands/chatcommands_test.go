package commands

import (
	"reflect"
	"testing"
	"time"
)

func createChatObj() ChatCommand {
	c := ChatCommand{}
	c.Constructor("test", func() {}, "random message", false, 200)
	return c
}

func TestChatCommand(t *testing.T) {
	c := createChatObj()
	rc := reflect.TypeOf(c)
	if rc.Name() != "ChatCommand" {
		t.Error("incorrect struct name")
	}

}

func TestChatGetName(t *testing.T) {
	c := createChatObj()
	if c.GetName() != "test" {
		t.Error("incorrect value")
	}
	if reflect.TypeOf(c.GetName()).String() != "string" {
		t.Error("incorrect return type")
	}
}

func TestChatIsOnCooldown(t *testing.T) {
	c := createChatObj()
	if reflect.TypeOf(c.IsOnCooldown()).String() != "bool" {
		t.Error("incorrect return type")
	}
	if c.IsOnCooldown() != false {
		t.Error("incorrect result")
	}
}

func TestChatIsPrivate(t *testing.T) {
	c := createChatObj()
	if reflect.TypeOf(c.IsPrivate()).String() != "bool" {
		t.Error("incorrect return type")
	}
	if c.IsPrivate() != false {
		t.Error("incorrect result")
	}
}

func TestChatGetMessage(t *testing.T) {
	c := createChatObj()
	if reflect.TypeOf(c.GetMessage()).String() != "string" {
		t.Error("incorrect return type")
	}
	if c.GetMessage() != c.message {
		t.Error("incorrect result")
	}
}

func TestChatCall(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	c := createChatObj()
	c.Call()
}

func TestChatCallWithIncorrectInterface(t *testing.T) {
	defer func() {
		//Check that we are catching panic
		if err := recover(); err == nil {
			t.Error(err)
		}
	}()
	c := ChatCommand{}
	c.Constructor("random", 1111, "", false, 0)
	c.Call()
}

func TestSetCoolDown(t *testing.T) {
	c := createChatObj()
	if c.IsOnCooldown() == true {
		t.Error("should not be true")
	}
	c.setCooldown()
	if c.IsOnCooldown() == false {
		t.Error("should be on Cooldown")
	}
	time.Sleep(time.Duration(c.cdTime-100) * time.Millisecond)
	if c.IsOnCooldown() == false {
		t.Error("should be on Cooldown")
	}
	time.Sleep(time.Duration(c.cdTime) * time.Millisecond)
	if c.IsOnCooldown() != false {
		t.Error("should NOT be on Cooldown")
	}
}
