package log

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"strings"
)

type contextHandler struct {
	slog.Handler
	keys []string
}

func (h contextHandler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(h.observe(ctx)...)
	return h.Handler.Handle(ctx, r)
}

func (h contextHandler) observe(ctx context.Context) (as []slog.Attr) {
	_, file, line, _ := runtime.Caller(4)
	codePath := strings.TrimPrefix(fmt.Sprintf("%s:%d", file, line), workingDirectory)

	as = append(as, slog.Attr{
		Key:   "source",
		Value: slog.StringValue(codePath),
	})

	for _, k := range h.keys {
		v := ctx.Value(k)

		switch k {
		case contextKeySpanId:
			if order, ok := v.(*int); ok {
				*order++
			} else {
				zero := 0
				v = &zero
			}
		}

		as = append(as, slog.Attr{
			Key:   k,
			Value: slog.AnyValue(v),
		})
	}
	return
}
