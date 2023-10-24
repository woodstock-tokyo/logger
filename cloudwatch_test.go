package logger

import (
	"testing"
)

func TestCloudWatchHook(t *testing.T) {
	SetAppName("test")
	Type("cloudwatch").WithFields(Fields{"app": "test"}).Error("test")
}
