package slogger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func Init(option ...Option) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	workingDirectory = strings.ReplaceAll(fmt.Sprintf("%s/", dir), `\`, `/`)

	if len(option) > 0 {
		opt := option[0]
		if opt.ContextKeyTraceId != nil {
			contextKeyTraceId = *opt.ContextKeyTraceId
		}
		if opt.ContextKeySpanId != nil {
			contextKeySpanId = *opt.ContextKeySpanId
		}
		if opt.TimeFormat != nil {
			timeFormat = *opt.TimeFormat
		}
		if opt.WorkingDirectory != nil {
			workingDirectory = *opt.WorkingDirectory
		}
		if opt.LogLevel != nil {
			logLevel = *opt.LogLevel
		}
		addSource = opt.AddSource
	}

	var h contextHandler
	h.Handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				// Convert the time to custom format
				a.Value = slog.StringValue(a.Value.Time().Format(timeFormat))
			}
			return a
		},
	})
	h.keyList = []string{
		contextKeyTraceId,
		contextKeySpanId,
	}
	h.addSource = addSource

	slog.SetDefault(slog.New(h))
	return nil
}
