package rslog

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLog struct {
	logger *zap.Logger

	level    string
	sugarLog *zap.SugaredLogger
}

func newZapLog() *zapLog {
	p := &zapLog{}

	p.logger = initZapLog(zapcore.InfoLevel)
	p.sugarLog = p.logger.Sugar()

	return p
}

func initZapLog(l zapcore.Level) *zap.Logger {
	core := zapcore.NewCore(
		// zapcore.NewJSONEncoder(ZapNewEncoderConfig()),
		zapcore.NewConsoleEncoder(ZapNewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(os.Stdout),
		l,
	)
	return zap.New(core, zap.AddCaller())
}

func (p *zapLog) Info(datas ...interface{}) {
	p.sugarLog.Info(datas...)
}

func (p *zapLog) Infof(format string, datas ...interface{}) {
	p.sugarLog.Infof(format, datas...)
}

func (p *zapLog) Debug(datas ...interface{}) {
	p.sugarLog.Debug(datas...)
}

func (p *zapLog) Debugf(format string, datas ...interface{}) {
	p.sugarLog.Debugf(format, datas...)
}

func (p *zapLog) Warn(datas ...interface{}) {
	p.sugarLog.Warn(datas...)
}

func (p *zapLog) Warnf(format string, datas ...interface{}) {
	p.sugarLog.Warnf(format, datas...)
}

func (p *zapLog) Error(datas ...interface{}) {
	p.sugarLog.Error(datas...)
}

func (p *zapLog) Errorf(format string, datas ...interface{}) {
	p.sugarLog.Errorf(format, datas...)
}

func (p *zapLog) SetLevel(l string) {
	p.logger = initZapLog(GetZapLevel(l))
	p.sugarLog = p.logger.Sugar()
}
func (p *zapLog) ResetLog(l interface{}) {
	if logger, ok := l.(*zap.Logger); ok {
		p.logger = logger
		p.sugarLog = p.logger.Sugar()
	}
}

// "debug": zapcore.DebugLevel,
// "info": zapcore.InfoLevel,
// "warn": zapcore.WarnLevel,
// "error": zapcore.ErrorLevel,
// "dpanic": zapcore.DPanicLevel,
// "panic": zapcore.PanicLevel,
// "fatal": zapcore.FatalLevel,
func GetZapLevel(l string) zapcore.Level {
	switch l {
	case "info":
		return zapcore.InfoLevel
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	case "dpanic":
		return zapcore.DPanicLevel
	}
	return zapcore.InfoLevel
}

func GetZapWriter(filename string, maxDay int) io.Writer {
	hook, err := rotatelogs.New(
		filename+".%Y%m%d", // ????????????go??????????????????format??????
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(maxDay)),
		// rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

// ???????????????
func ZapTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func ZapNewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",                        // json????????????
		LevelKey:       "level",                       // json??????????????????
		NameKey:        "name",                        // json?????????????????????
		CallerKey:      "call",                        // json????????????????????????
		MessageKey:     "msg",                         // json??????????????????
		StacktraceKey:  "stack",                       // json????????????
		LineEnding:     zapcore.DefaultLineEnding,     // ?????????????????????
		EncodeLevel:    zapcore.CapitalLevelEncoder,   // ?????????????????????????????????info INFO???
		EncodeTime:     ZapTimeEncoder,                // ??????????????????????????????
		EncodeDuration: zapcore.StringDurationEncoder, // ???????????????
		EncodeCaller:   zapcore.ShortCallerEncoder,    // ????????????????????????/??????.go:?????????
	}
}
