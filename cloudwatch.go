package logger

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	cloudwatchlogs "github.com/kdar/logrus-cloudwatchlogs"
)

func NewCloudWatchHook(region, accessKey, accessSecret, group, stream string) *cloudwatchlogs.Hook {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, accessSecret, ""),
	})

	if err != nil {
		panic("invalid aws session for cloudwatch")
	}

	hook, err := cloudwatchlogs.NewHook(group, stream, sess)
	if err != nil {
		panic(err)
	}

	return hook
}
