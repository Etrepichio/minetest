package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service
type Middleware func(Minesweepersvc) Minesweepersvc

// LoggingMiddleware initializes loggingMiddleware struct
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Minesweepersvc) Minesweepersvc {
		return loggingMiddleware{logger, next}
	}
}

// loggindMiddleware struct
type loggingMiddleware struct {
	logger log.Logger     // go-kit logger instance
	next   Minesweepersvc // reference to service
}

// GetMinesweeper implements loggingMiddleware for
// Minesweepersvc.
func (mw loggingMiddleware) GetMinesweeper(ctx context.Context) (res MinesweeperResponse, err error) {
	// defer logging to log response
	defer func() {
		mw.logger.Log("method", "GetMinesweeper",
			"name", res.Name,
			"error", err)
	}()

	// call Minesweepersvc to invoke
	// next middleware (or service)
	return mw.next.GetMinesweeper(ctx)
}
