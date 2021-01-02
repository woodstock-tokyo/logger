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
type Fields map[string]interface{}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
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
func (e *Exporter) Print(args ...interface{}) {
	e.decorateRuntime().le.Print(args...)
}

// Printf printf
func (e *Exporter) Printf(format string, args ...interface{}) {
	e.decorateRuntime().le.Printf(format, args...)
}

// Println println
func (e *Exporter) Println(args ...interface{}) {
	e.decorateRuntime().le.Println(args...)
}

// Debug debug
func (e *Exporter) Debug(args ...interface{}) {
	e.decorateRuntime().le.Debug(args...)
}

// Debugf debugf
func (e *Exporter) Debugf(format string, args ...interface{}) {
	e.decorateRuntime().le.Debugf(format, args...)
}

// Debugln debugln
func (e *Exporter) Debugln(args ...interface{}) {
	e.decorateRuntime().le.Debugln(args...)
}

// Info info
func (e *Exporter) Info(args ...interface{}) {
	e.decorateRuntime().le.Info(args...)
}

// Infof infof
func (e *Exporter) Infof(format string, args ...interface{}) {
	e.decorateRuntime().le.Infof(format, args...)
}

// Infoln infoln
func (e *Exporter) Infoln(args ...interface{}) {
	e.decorateRuntime().le.Infoln(args...)
}

// Warn warn
func (e *Exporter) Warn(args ...interface{}) {
	e.decorateRuntime().le.Warn(args...)
}

// Warnf warnf
func (e *Exporter) Warnf(format string, args ...interface{}) {
	e.decorateRuntime().le.Warnf(format, args...)
}

// Warnln warnln
func (e *Exporter) Warnln(args ...interface{}) {
	e.decorateRuntime().le.Warnln(args...)
}

// Error error
func (e *Exporter) Error(args ...interface{}) {
	e.decorateRuntime().le.Error(args...)
}

// Errorf errorf
func (e *Exporter) Errorf(format string, args ...interface{}) {
	e.decorateRuntime().le.Errorf(format, args...)
}

// Errorln errorln
func (e *Exporter) Errorln(args ...interface{}) {
	e.decorateRuntime().le.Errorln(args...)
}

// Fatal fatal
func (e *Exporter) Fatal(args ...interface{}) {
	e.decorateRuntime().le.Fatal(args...)
}

// Fatalf fatalf
func (e *Exporter) Fatalf(format string, args ...interface{}) {
	e.decorateRuntime().le.Fatalf(format, args...)
}

// Fatalln fatalln
func (e *Exporter) Fatalln(args ...interface{}) {
	e.decorateRuntime().le.Fatalln(args...)
}

// Panic panic
func (e *Exporter) Panic(args ...interface{}) {
	e.decorateRuntime().le.Panic(args...)
}

// Panicf panicf
func (e *Exporter) Panicf(format string, args ...interface{}) {
	e.decorateRuntime().le.Panicf(format, args...)
}

// Panicln panicln
func (e *Exporter) Panicln(args ...interface{}) {
	e.decorateRuntime().le.Panicln(args...)
}

// Type init logger entry with logger type
func Type(typeName string) *Exporter {
	formatedFields := Fields{
		"app":  app,
		"type": typeName,
	}
	return &Exporter{
		le: logrus.WithFields(logrus.Fields(formatedFields)),
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
func Print(args ...interface{}) {
	defaultExporter().Print(args...)
}

// Printf printf
func Printf(format string, args ...interface{}) {
	defaultExporter().Printf(format, args...)
}

// Println println
func Println(args ...interface{}) {
	defaultExporter().Println(args...)
}

// Debug debug
func Debug(args ...interface{}) {
	defaultExporter().Debug(args...)
}

// Debugf debugf
func Debugf(format string, args ...interface{}) {
	defaultExporter().Debugf(format, args...)
}

// Derbugln debugln
func Debugln(args ...interface{}) {
	defaultExporter().Debugln(args...)
}

// Info info
func Info(args ...interface{}) {
	defaultExporter().Info(args...)
}

// Infof infof
func Infof(format string, args ...interface{}) {
	defaultExporter().Infof(format, args...)
}

// Infoln infoln
func Infoln(args ...interface{}) {
	defaultExporter().Infoln(args...)
}

// Warn warn
func Warn(args ...interface{}) {
	defaultExporter().Warn(args...)
}

// Warnf warnf
func Warnf(format string, args ...interface{}) {
	defaultExporter().Warnf(format, args...)
}

// Warnln warnln
func Warnln(args ...interface{}) {
	defaultExporter().Warnln(args...)
}

// Error error
func Error(args ...interface{}) {
	defaultExporter().Error(args...)
}

// Errorf errorf
func Errorf(format string, args ...interface{}) {
	defaultExporter().Errorf(format, args...)
}

// Errorln errorln
func Errorln(args ...interface{}) {
	defaultExporter().Errorln(args...)
}

// Fatal fatal
func Fatal(args ...interface{}) {
	defaultExporter().Fatal(args...)
}

// Fatalf fatalf
func Fatalf(format string, args ...interface{}) {
	defaultExporter().Fatalf(format, args...)
}

// Fatalln Fatalln
func Fatalln(args ...interface{}) {
	defaultExporter().Fatalln(args...)
}

// Panic panic
func Panic(args ...interface{}) {
	defaultExporter().Panic(args...)
}

// Panicf panicf
func Panicf(format string, args ...interface{}) {
	defaultExporter().Panicf(format, args...)
}

// Panicln panicln
func Panicln(args ...interface{}) {
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
