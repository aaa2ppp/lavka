package helper

import (
	"context"
	"log/slog"

	"lavka/internal/logger"
)

type Helper struct {
	ctx context.Context
	gr  string
	op  string
	log *slog.Logger
}

func New(ctx context.Context, gr, op string) *Helper {
	return &Helper{
		ctx: ctx,
		gr:  gr,
		op:  op,
	}
}

func (x *Helper) Log() *slog.Logger {
	if x.log == nil {
		x.log = logger.GetLoggerFromContextOrDefault(x.ctx).
			WithGroup(x.gr).With("op", x.op)
	}
	return x.log
}


