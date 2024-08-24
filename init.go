package slogger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/google/uuid"
)

var (
	contextKeyTraceId = "trace-id"
	contextKeySpanId  = "span-id"
	timeFormat        = "2006-01-02T15:04:05.000-07:00"
	workingDirectory  = ""
	level             = slog.LevelDebug
)

const (
	ContextKeyTraceId = "trace-id"
	ContextKeySpanId  = "span-id"
)

type Option struct {
	ContextKeyTraceId *string
	ContextKeySpanId  *string
	TimeFormat        *string
	WorkingDirectory  *string
	Level             *slog.Level
}

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
		if opt.Level != nil {
			level = *opt.Level
		}
	}

	h := contextHandler{
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
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

func NewContext(oldCtx ...context.Context) (newCtx context.Context) {
	newCtx = context.Background()
	if len(oldCtx) > 0 {
		newCtx = oldCtx[0]
	}

	newCtx = context.WithValue(newCtx, ContextKeyTraceId, uuid.NewString())
	newCtx = context.WithValue(newCtx, ContextKeySpanId, &[]int{-1}[0])

	return newCtx
}
