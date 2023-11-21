package core

import (
	"context"
	"log/slog"
)

type mdkCtxKey struct{}

var ctxKey = mdkCtxKey{}

func WithTraceID(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, ctxKey, traceId)
}

func WithWriteTraceIdHandler(parent slog.Handler) *WriteTraceIdHandler {
	return &WriteTraceIdHandler{
		parent: parent,
	}
}

type WriteTraceIdHandler struct {
	parent slog.Handler
}

func (h *WriteTraceIdHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.parent.Enabled(ctx, level)
}

func (h *WriteTraceIdHandler) Handle(ctx context.Context, record slog.Record) error {
	record.Add(slog.String("traceId", ctx.Value(ctxKey).(string)))
	return h.parent.Handle(ctx, record)
}

func (h *WriteTraceIdHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &WriteTraceIdHandler{h.parent.WithAttrs(attrs)}
}

func (h *WriteTraceIdHandler) WithGroup(name string) slog.Handler {
	return &WriteTraceIdHandler{h.parent.WithGroup(name)}
}

var _ slog.Handler = (*WriteTraceIdHandler)(nil)
