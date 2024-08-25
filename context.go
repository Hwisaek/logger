package slogger

import (
	"context"
	"github.com/google/uuid"
)

func NewContext(oldCtx ...context.Context) (newCtx context.Context) {
	newCtx = context.Background()
	if len(oldCtx) > 0 {
		newCtx = oldCtx[0]
	}

	newCtx = context.WithValue(newCtx, ContextKeyTraceId, uuid.NewString())
	newCtx = context.WithValue(newCtx, ContextKeySpanId, &[]int{-1}[0])

	return newCtx
}
