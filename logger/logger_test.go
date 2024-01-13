package logger

import "testing"

func TestLogger(t *testing.T) {
	Info("test")
	Infof("test %s", "test")
	Warning("test")
	Warningf("test %s", "test")
	Debug("test")
	Debugf("test %s", "test")
	Error("test")
	Errorf("test %s", "test")
}
