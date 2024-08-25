package slogger

import "log/slog"

var (
	contextKeyTraceId = "trace-id"
	contextKeySpanId  = "span-id"
	timeFormat        = "2006-01-02T15:04:05.000-07:00"
	workingDirectory  = ""
	logLevel          = slog.LevelDebug
	addSource         = false
)

type Option struct {
	ContextKeyTraceId *string
	ContextKeySpanId  *string
	TimeFormat        *string
	WorkingDirectory  *string
	LogLevel          *slog.Level
	AddSource         bool
}

func GetLogLevel() slog.Level {
	return logLevel
}
