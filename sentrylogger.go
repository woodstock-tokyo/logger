package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/woodstock-tokyo/go-aws-sdk/secretsmanager"
)

var (
	sentryDSN string
	once      sync.Once
)

func getSentryDSN() string {
	once.Do(func() {
		env := os.Getenv("WS_JOBS_ENVIRONMENT")
		secretID := fmt.Sprintf("woodstock-jobs-%s", env)

		svc := secretsmanager.NewService(
			os.Getenv("WS_SECRETS_MANAGER_AWS_ACCESS_KEY_ID"),
			os.Getenv("WS_SECRETS_MANAGER_AWS_SECRET_ACCESS_KEY"),
		)
		svc.SetRegion("ap-northeast-1")

		resp := svc.GetSecretValue(secretID)
		if resp.Error != nil {
			WithFields(Fields{
				"error": resp.Error,
				"env":   env,
			}).Error("Failed to get Sentry DSN from Secrets Manager")
			return
		}

		dsn, ok := resp.SecretValue["WS_SENTRY_DSN"]
		if !ok || dsn == "" {
			Error("SENTRY_DSN not found in secret")
			return
		}

		sentryDSN = dsn
	})

	return sentryDSN
}

func InitSentry() error {
	dsn := getSentryDSN()
	if dsn == "" {
		return fmt.Errorf("empty Sentry DSN")
	}

	return sentry.Init(sentry.ClientOptions{
		Dsn:         dsn,
		Environment: os.Getenv("WS_JOBS_ENVIRONMENT"),
		Debug:       os.Getenv("WS_JOBS_ENVIRONMENT") == "development",
	})
}

// LogToSentry sends message to Sentry with environment info
func LogToSentry(message string, level sentry.Level) {
	// add environment info to message
	env := os.Getenv("WS_JOBS_ENVIRONMENT")
	fullMessage := fmt.Sprintf("[%s](%s) %s", level, env, message)

	sentry.CaptureMessage(fullMessage)

	// flush to ensure message is sent
	sentry.Flush(2 * time.Second)
}

// LogError sends error to Sentry
func LogError(err error) {
	sentry.CaptureException(err)
	sentry.Flush(2 * time.Second)
}
