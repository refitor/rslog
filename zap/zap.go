package zap

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	zaplog "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLog struct {
	logger *zaplog.Logger

	level    string
	depth    int
	sugarLog *zaplog.SugaredLogger
}

func New() *zapLog {
	p := &zapLog{
		level: "info",
		depth: 4,
	}

	p.logger = initZapLog(zapcore.InfoLevel)
	p.sugarLog = p.logger.Sugar()

	return p
}

func initZapLog(l zapcore.Level) *zaplog.Logger {
	core := zapcore.NewCore(
		// zapcore.NewJSONEncoder(ZapNewEncoderConfig()),
		zapcore.NewConsoleEncoder(ZapNewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(os.Stdout),
		l,
	)
	return zaplog.New(core, zaplog.AddCaller())
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

func (p *zapLog) SetDepth(depth int) int {
	oldDepth := p.depth
	p.depth = depth
	return oldDepth
}

func (p *zapLog) SetLevel(l string) {
	p.logger = initZapLog(getZapLevel(l))
	p.sugarLog = p.logger.Sugar()
}
func (p *zapLog) ResetLog(l interface{}) {
	if logger, ok := l.(*zaplog.Logger); ok {
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
func getZapLevel(l string) zapcore.Level {
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
		filename+".%Y%m%d", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(maxDay)),
		// rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

// 格式化时间
func ZapTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func ZapNewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",                        // json时时间键
		LevelKey:       "level",                       // json时日志等级键
		NameKey:        "name",                        // json时日志记录器键
		CallerKey:      "call",                        // json时日志文件信息键
		MessageKey:     "msg",                         // json时日志消息键
		StacktraceKey:  "stack",                       // json时堆栈键
		LineEnding:     zapcore.DefaultLineEnding,     // 友好日志换行符
		EncodeLevel:    zapcore.CapitalLevelEncoder,   // 友好日志等级名大小写（info INFO）
		EncodeTime:     ZapTimeEncoder,                // 友好日志时日期格式化
		EncodeDuration: zapcore.StringDurationEncoder, // 时间序列化
		EncodeCaller:   zapcore.ShortCallerEncoder,    // 日志文件信息（包/文件.go:行号）
	}
}
