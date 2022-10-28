package logger

import (
	"testing"
)

func TestPagerDutyHook(t *testing.T) {
	SetAppName("test")
	Type("pagerduty").WithFields(Fields{"app": "test"}).Error("test")
}
