package sys

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type (
	t_LogLevel int
	t_LogType  int
)

const (
	c_LOG_FATAL = t_LogType(0x1)
	c_LOG_ERROR = t_LogType(0x2)
	c_LOG_WARN  = t_LogType(0x4)
	c_LOG_INFO  = t_LogType(0x8)
	c_LOG_DEBUG = t_LogType(0x10)
)

const (
	c_LOG_LEVEL_NONE  = t_LogLevel(0x0)
	c_LOG_LEVEL_FATAL = c_LOG_LEVEL_NONE | t_LogLevel(c_LOG_FATAL)
	c_LOG_LEVEL_ERROR = c_LOG_LEVEL_FATAL | t_LogLevel(c_LOG_ERROR)
	c_LOG_LEVEL_WARN  = c_LOG_LEVEL_ERROR | t_LogLevel(c_LOG_WARN)
	c_LOG_LEVEL_INFO  = c_LOG_LEVEL_WARN | t_LogLevel(c_LOG_INFO)
	c_LOG_LEVEL_DEBUG = c_LOG_LEVEL_INFO | t_LogLevel(c_LOG_DEBUG)
	c_LOG_LEVEL_ALL   = c_LOG_LEVEL_DEBUG
)

func stackTrace(depth int) string {
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		return ""
	}
	strFileLine := fmt.Sprintf("[%s %s:%v]", filepath.Base(filepath.Dir(file)), filepath.Base(file), line)
	return strFileLine
}

func getLogLevel(level string) t_LogLevel {
	switch level {
	case "info":
		return c_LOG_LEVEL_INFO
	case "debug":
		return c_LOG_LEVEL_DEBUG
	case "error":
		return c_LOG_LEVEL_ERROR
	case "warn":
		return c_LOG_LEVEL_WARN
	}
	return 0
}
