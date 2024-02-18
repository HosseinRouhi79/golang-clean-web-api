package logging

import "github.com/HosseinRouhi79/golang-clean-web-api/src/config"


type Logger interface {
	Init()
	
	Debug(cat Category, subCat SubCategory,  msg string, extra map[ExtraKey]interface{})
	Debugf(tempplate string, args ...interface{})

	Info(cat Category, subCat SubCategory,  msg string, extra map[ExtraKey]interface{})
	Infof(tempplate string, args ...interface{})

	Fatal(cat Category, subCat SubCategory,  msg string, extra map[ExtraKey]interface{})
	Fatalf(tempplate string, args ...interface{})

	Warn(cat Category, subCat SubCategory,  msg string, extra map[ExtraKey]interface{})
	Warnf(tempplate string, args ...interface{})
}

func NewLogger(cfg *config.Config) Logger {
	if cfg.Logger.Logger == "zap" {
		return newZapLogger()
	}else if cfg.Logger.Logger == "zerolog"{
		return newZeroLogger()
	}
}