package logging

import (
	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"github.com/rs/zerolog"
)

type zerologger struct {
	cfg    *config.Config
	logger *zerolog.Logger
}

var zeroFinLogger *zerologger

var zerologMap = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error":  zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

func (l *zerologger) getLevel(*zerologMap map[string]zerolog.Level) zerolog.Level {
	level, ok := zerologMap[l.cfg.Logger.Level]
	if !ok {
		return zerologMap.DebugLevel
	}
	return level
}

func newZeroLogger(cfg zerolog.Config) zerologger{
	zero := &zerologger{cfg: cfg}
	zero.Init()
	return zero
}

func (zerologger *zerologger) Init() {

}

func Debug(message string) {}

func Debugf(message string) {}

func Info(message string) {}

func Infof(message string) {}

func Fatal(message string) {}

func Fatalf(message string) {}

func Warn(message string) {}

func Warnf(message string) {}
