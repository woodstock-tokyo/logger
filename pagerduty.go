package logger

import (
	"context"
	"time"

	pd "github.com/PagerDuty/go-pagerduty"
	"github.com/sirupsen/logrus"
)

type hook struct {
	serviceKey string
	levels     []logrus.Level
}

func NewPagerDutyHook(serviceKey string) *hook {
	return &hook{
		serviceKey: serviceKey,
		levels: []logrus.Level{
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		},
	}
}

func (hook *hook) Fire(entry *logrus.Entry) error {
	severity := "info"
	if entry.Level == logrus.ErrorLevel {
		severity = "error"
	} else if entry.Level == logrus.FatalLevel {
		severity = "critical"
	} else if entry.Level == logrus.PanicLevel {
		severity = "critical"
	}

	source, ok := entry.Data["app"]
	if !ok || source == "" {
		source = "unknown"
	}

	// https://developer.pagerduty.com/docs/ZG9jOjExMDI5NTgx-send-an-alert-event
	event := pd.V2Event{
		RoutingKey: hook.serviceKey,
		Action:     "trigger",
		Payload: &pd.V2Payload{
			Source:    source.(string),
			Summary:   entry.Message,
			Details:   entry.Data,
			Timestamp: entry.Time.Format(time.RFC3339Nano),
			Severity:  severity,
			Group:     app,
		},
	}

	_, err := pd.ManageEventWithContext(context.Background(), event)
	if err != nil {
		entry.Warnf("failed to alert pagerduty: %+v", err)
	}

	return nil
}

func (hook *hook) Levels() []logrus.Level {
	return hook.levels
}
