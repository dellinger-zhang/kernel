package logger

import (
	"testing"
)

func TestLoggerNormal(t *testing.T) {

	Info("Test 123")
	Debug("Test 123")
	Warning("Test 123")
	Error("Test 123")
	// Fatal("Test 123")
}

func TestLoggerFormat(t *testing.T) {

	Infof("This is a %s", "man")
	Debugf("He is %d years old.", 5)
	Warningf("He is a %04d kg", 75)
	Errorf("Hi, Mr.%s", "Dellinger")
}
