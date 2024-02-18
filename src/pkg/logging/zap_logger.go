package logging

import (
	"fmt"
	"sync"
	"time"

	"github.com/HosseinRouhi79/golang-clean-web-api/src/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

var zapOnce sync.Once

var zapLogLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.WarnLevel,
	"fatal": zap.FatalLevel,
}

var zapFinLogger *zap.SugaredLogger

func (l *zapLogger) getLevel(zapMap map[string]zapcore.Level) zapcore.Level {
	level, ok := zapMap[l.cfg.Logger.Level]
	if !ok {
		return zap.DebugLevel
	}
	return level
}

func (z *zapLogger) Init() {

	zapOnce.Do(func() {
		fileName := fmt.Sprintf("%s%s-%s.%s", z.cfg.Logger.FilePath, time.Now().Format("2006-01-02"))
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    1,
			MaxAge:     5,
			MaxBackups: 10,
			Compress:   true,
			LocalTime:  true,
		})
		zEncoder := zap.NewProductionEncoderConfig()
		zEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(zapcore.NewJSONEncoder(zEncoder), w, z.getLevel(zapLogLevelMap))

		logger := zap.New(core, zap.AddStacktrace(zap.ErrorLevel)).Sugar()

		zapFinLogger = logger.With("AppName", "myApp", "LoggerName", "Zap")
	})

	z.logger = zapFinLogger

}

func newZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger
}
