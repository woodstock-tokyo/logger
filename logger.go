package logger

import (
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var app string

// Fields type, used to pass to `WithFields`.
type Fields map[string]any

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	//InitSentry()
}

// Level log level
type Level logrus.Level

const (
	// InfoLevel info
	InfoLevel Level = Level(logrus.InfoLevel)
	// DebugLevel debug
	DebugLevel = Level(logrus.DebugLevel)
	// ErrorLevel error
	ErrorLevel = Level(logrus.ErrorLevel)
	// WarnLevel warn
	WarnLevel = Level(logrus.WarnLevel)
	// FatalLevel fatal
	FatalLevel = Level(logrus.FatalLevel)
	// PanicLevel panic
	PanicLevel = Level(logrus.PanicLevel)
)

// SetLevel set log level
func SetLevel(level Level) {
	logrus.SetLevel(logrus.Level(level))
}

// SetAppName set application name
func SetAppName(name string) {
	app = name
}

// Exporter export entry
type Exporter struct {
	le *logrus.Entry
}

// WithFields with fields
func (e *Exporter) WithFields(fields Fields) *Exporter {
	copyfields := e.le.Data
	copyfields["data"] = fields
	e.le.Data = copyfields
	return e
}

// WithError with fields
func (e *Exporter) WithError(err error) *Exporter {
	return e.WithFields(Fields{"err": err.Error()})
}

// WithSecretFields with secret fields
func (e *Exporter) WithSecretFields(fields Fields) *Exporter {
	copyfields := e.le.Data
	copyfields["data_secret"] = fields
	e.le.Data = copyfields
	return e
}

// Type type
func (e *Exporter) Type(typeName string) *Exporter {
	copyfields := e.le.Data
	copyfields["type"] = typeName
	e.le.Data = copyfields
	return e
}

// Print print
func (e *Exporter) Print(args ...any) {
	e.decorateRuntime().le.Print(args...)
}

// Printf printf
func (e *Exporter) Printf(format string, args ...any) {
	e.decorateRuntime().le.Printf(format, args...)
}

// Println println
func (e *Exporter) Println(args ...any) {
	e.decorateRuntime().le.Println(args...)
}

// Debug debug
func (e *Exporter) Debug(args ...any) {
	e.decorateRuntime().le.Debug(args...)
}

// Debugf debugf
func (e *Exporter) Debugf(format string, args ...any) {
	e.decorateRuntime().le.Debugf(format, args...)
}

// Debugln debugln
func (e *Exporter) Debugln(args ...any) {
	e.decorateRuntime().le.Debugln(args...)
}

// Info info
func (e *Exporter) Info(args ...any) {
	e.decorateRuntime().le.Info(args...)
}

// Infof infof
func (e *Exporter) Infof(format string, args ...any) {
	e.decorateRuntime().le.Infof(format, args...)
}

// Infoln infoln
func (e *Exporter) Infoln(args ...any) {
	e.decorateRuntime().le.Infoln(args...)
}

// Warn warn
func (e *Exporter) Warn(args ...any) {
	e.decorateRuntime().le.Warn(args...)
}

// Warnf warnf
func (e *Exporter) Warnf(format string, args ...any) {
	e.decorateRuntime().le.Warnf(format, args...)
}

// Warnln warnln
func (e *Exporter) Warnln(args ...any) {
	e.decorateRuntime().le.Warnln(args...)
}

// Error error
func (e *Exporter) Error(args ...any) {
	e.decorateRuntime().le.Error(args...)
}

// Errorf errorf
func (e *Exporter) Errorf(format string, args ...any) {
	e.decorateRuntime().le.Errorf(format, args...)
}

// Errorln errorln
func (e *Exporter) Errorln(args ...any) {
	e.decorateRuntime().le.Errorln(args...)
}

// Fatal fatal
func (e *Exporter) Fatal(args ...any) {
	e.decorateRuntime().le.Fatal(args...)
}

// Fatalf fatalf
func (e *Exporter) Fatalf(format string, args ...any) {
	e.decorateRuntime().le.Fatalf(format, args...)
}

// Fatalln fatalln
func (e *Exporter) Fatalln(args ...any) {
	e.decorateRuntime().le.Fatalln(args...)
}

// Panic panic
func (e *Exporter) Panic(args ...any) {
	e.decorateRuntime().le.Panic(args...)
}

// Panicf panicf
func (e *Exporter) Panicf(format string, args ...any) {
	e.decorateRuntime().le.Panicf(format, args...)
}

// Panicln panicln
func (e *Exporter) Panicln(args ...any) {
	e.decorateRuntime().le.Panicln(args...)
}

// Type init logger entry with logger type
func Type(typeName string) *Exporter {
	loggerInstance := logrus.New()
	loggerInstance.SetFormatter(&logrus.JSONFormatter{})
	loggerInstance.SetOutput(os.Stdout)

	// route to pagerduty
	if typeName == "pagerduty" {
		loggerInstance.AddHook(NewPagerDutyHook(os.Getenv("PAGERDUTY_API_KEY")))
	}

	// route to cloudwatch
	if typeName == "cloudwatch" {
		region := os.Getenv("AWS_REGION")
		if region == "" {
			region = "ap-northeast-1"
		}
		loggerInstance.AddHook(NewCloudWatchHook(region,
			os.Getenv("WS_CLOUDWATCH_AWS_ACCESS_KEY_ID"),
			os.Getenv("WS_CLOUDWATCH_AWS_SECRET_ACCESS_KEY"),
			os.Getenv("WS_CLOUDWATCH_LOG_GROUP"),
			os.Getenv("WS_CLOUDWATCH_STREAM_NAME")),
		)
	}

	formatedFields := Fields{
		"app":  app,
		"type": typeName,
	}

	return &Exporter{
		le: loggerInstance.WithFields(logrus.Fields(formatedFields)),
	}
}

// WithFields init logger entry with fields
func WithFields(fields Fields) *Exporter {
	formatedFields := Fields{
		"app":  app,
		"data": fields,
	}
	return &Exporter{
		le: logrus.WithFields(logrus.Fields(formatedFields)),
	}
}

// WithSecretFields init logger with secret fields
func WithSecretFields(fields Fields) *Exporter {
	formatedFields := Fields{
		"app":         app,
		"data_secret": fields,
	}
	return &Exporter{
		le: logrus.WithFields(logrus.Fields(formatedFields)),
	}
}

// defaultExporter default exporter with "app" field
func defaultExporter() *Exporter {
	formatedFields := Fields{
		"app": app,
	}
	return &Exporter{
		le: logrus.WithFields(logrus.Fields(formatedFields)),
	}
}

// Print print
func Print(args ...any) {
	defaultExporter().Print(args...)
}

// Printf printf
func Printf(format string, args ...any) {
	defaultExporter().Printf(format, args...)
}

// Println println
func Println(args ...any) {
	defaultExporter().Println(args...)
}

// Debug debug
func Debug(args ...any) {
	defaultExporter().Debug(args...)
}

// Debugf debugf
func Debugf(format string, args ...any) {
	defaultExporter().Debugf(format, args...)
}

// Derbugln debugln
func Debugln(args ...any) {
	defaultExporter().Debugln(args...)
}

// Info info
func Info(args ...any) {
	defaultExporter().Info(args...)
}

// Infof infof
func Infof(format string, args ...any) {
	defaultExporter().Infof(format, args...)
}

// Infoln infoln
func Infoln(args ...any) {
	defaultExporter().Infoln(args...)
}

// Warn warn
func Warn(args ...any) {
	defaultExporter().Warn(args...)
}

// Warnf warnf
func Warnf(format string, args ...any) {
	defaultExporter().Warnf(format, args...)
}

// Warnln warnln
func Warnln(args ...any) {
	defaultExporter().Warnln(args...)
}

// Error error
func Error(args ...any) {
	defaultExporter().Error(args...)
}

// Errorf errorf
func Errorf(format string, args ...any) {
	defaultExporter().Errorf(format, args...)
}

// Errorln errorln
func Errorln(args ...any) {
	defaultExporter().Errorln(args...)
}

// Fatal fatal
func Fatal(args ...any) {
	defaultExporter().Fatal(args...)
}

// Fatalf fatalf
func Fatalf(format string, args ...any) {
	defaultExporter().Fatalf(format, args...)
}

// Fatalln Fatalln
func Fatalln(args ...any) {
	defaultExporter().Fatalln(args...)
}

// Panic panic
func Panic(args ...any) {
	defaultExporter().Panic(args...)
}

// Panicf panicf
func Panicf(format string, args ...any) {
	defaultExporter().Panicf(format, args...)
}

// Panicln panicln
func Panicln(args ...any) {
	defaultExporter().Panicln(args...)
}

func (e *Exporter) decorateRuntime() *Exporter {
	e.le = decorateRuntimeContext(e.le)
	return e
}

func decorateRuntimeContext(logger *logrus.Entry) *logrus.Entry {
	var currentFile string
	if _, file, _, ok := runtime.Caller(1); ok {
		currentFile = file
	} else {
		return logger
	}

	skip := 2
	for {
		if pc, file, line, ok := runtime.Caller(skip); ok {
			if file == currentFile {
				skip = skip + 1
				continue
			}
			fName := runtime.FuncForPC(pc).Name()
			slash := strings.LastIndex(file, "/")
			if slash >= 0 {
				file = file[slash+1:]
			}
			return logger.WithField("file", file).WithField("line", line).WithField("func", fName)
		}

		return logger
	}
}
