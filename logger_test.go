package logger

import "testing"

func TestLogger(t *testing.T) {
	SetAppName("hoge")
	Print("error")
	WithFields(Fields{
		"animal": "walrus",
	}).Print("error")
	Type("auth").Print("error")
	Type("auth").WithFields(Fields{
		"animal": "walrus",
	}).Print("error")
	WithFields(Fields{
		"animal": "walrus",
	}).Type("auth").Print("error")
	WithSecretFields(Fields{
		"user": "hogehoge",
	}).Print("error")
	WithSecretFields(Fields{
		"user": "hogehoge",
	}).Type("auth").Print("error")
	Type("auth").WithFields(Fields{
		"animal": "walrus",
	}).WithSecretFields(Fields{
		"user": "hogehoge",
	}).Print("error")
	Type("auth").WithSecretFields(Fields{
		"user": "hogehoge",
	}).WithFields(Fields{
		"animal": "walrus",
	}).Print("error")
}
