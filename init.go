package log

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

var (
	contextKeyTraceId = "trace-id"
	contextKeySpanId  = "span-id"
	timeFormat        = "2006-01-02T15:04:05.000-07:00"
	workingDirectory  = ""
)

type Option struct {
	ContextKeyTraceId *string
	ContextKeySpanId  *string
	TimeFormat        *string
	WorkingDirectory  *string
}

func Init(option ...Option) error {

	dir, err := os.Getwd()
	if err == nil {
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
	}

	h := contextHandler{
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					// Convert the time to custom format
					a.Value = slog.StringValue(a.Value.Time().Format(timeFormat))
				}
				return a
			},
		}),
		[]string{
			contextKeyTraceId,
			contextKeySpanId,
		},
	}

	slog.SetDefault(slog.New(h))
	return nil
}
