package echo

import (
	"reflect"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

func TestNewLoggerWrapper(t *testing.T) {
	logs := logrus.New()

	type args struct {
		logger *logrus.Logger
	}
	tests := []struct {
		name string
		args args
		want *LoggerWrapper
	}{
		{
			name: "all success",
			args: args{
				logger: logs,
			},
			want: &LoggerWrapper{
				Logger: logs,
				prefix: "echo",
				level:  log.INFO,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLoggerWrapper(tt.args.logger)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoggerWrapper() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoggerMethods(t *testing.T) {
	logs := logrus.New()

	wrapper := NewLoggerWrapper(logs)
	if wrapper == nil {
		t.Errorf("NewLoggerWrapper() got = %v", nil)
	}

	if got := wrapper.Output(); got == nil {
		t.Errorf("TestLoggerMethods() - Output got = %v, want %v", got, nil)
	}
	if got := wrapper.Prefix(); got != "echo" {
		t.Errorf("TestLoggerMethods() - Prefix got = %v, want %v", got, "echo")
	}
	wrapper.SetPrefix("new-prefix")
	if got := wrapper.Prefix(); got != "new-prefix" {
		t.Errorf("TestLoggerMethods() - SetPrefix - Prefix got = %v, want %v", got, "new-prefix")
	}
	if got := wrapper.Level(); got != log.INFO {
		t.Errorf("TestLoggerMethods() - Level got = %v, want %v", got, log.INFO)
	}
	wrapper.SetLevel(log.WARN)
	if got := wrapper.Level(); got != log.WARN {
		t.Errorf("TestLoggerMethods() - SetLevel - Level got = %v, want %v", got, log.WARN)
	}
	wrapper.SetHeader("header")
	wrapper.Printj(log.JSON{})
	wrapper.Debugj(log.JSON{})
	wrapper.Infoj(log.JSON{})
	wrapper.Warnj(log.JSON{})
	wrapper.Errorj(log.JSON{})
	//wrapper.Fatalj(log.JSON{})
	//wrapper.Panicj(log.JSON{})
}
