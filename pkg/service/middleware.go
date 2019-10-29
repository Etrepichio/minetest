package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/minesweeper/pkg/models"
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

func (mw loggingMiddleware) NewGame(ctx context.Context, game *models.Game) (err error) {

	// defer logging to log response
	defer func() {
		mw.logger.Log("method", "NewGame",
			"name", game.Name,
			"error", err)
	}()

	// call Minesweepersvc to invoke
	// next middleware (or service)
	return mw.next.NewGame(ctx, game)

}

func (mw loggingMiddleware) LoadGame(ctx context.Context, name string) (res *models.Game, err error) {

	// defer logging to log response
	defer func() {
		mw.logger.Log("method", "LoadGame",
			"name", res.Name,
			"error", err)
	}()

	// call Minesweepersvc to invoke
	// next middleware (or service)
	return mw.next.LoadGame(ctx, name)
}

func (mw loggingMiddleware) SaveGame(ctx context.Context, game *models.Game) (err error) {

	// defer logging to log response
	defer func() {
		mw.logger.Log("method", "SaveGame",
			"name", game.Name,
			"error", err)
	}()

	// call Minesweepersvc to invoke
	// next middleware (or service)
	return mw.next.SaveGame(ctx, game)
}

func (mw loggingMiddleware) Click(ctx context.Context, name string, rowClick int, columnClick int) (res *models.Game, err error) {

	// defer logging to log response
	defer func() {
		mw.logger.Log("method", "Click",
			"name", res.Name,
			"error", err)
	}()

	// call Minesweepersvc to invoke
	// next middleware (or service)
	return mw.next.Click(ctx, name, rowClick, columnClick)

}
