package service

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
)

var (
	// Known error types
	NoSweeper = errors.New("Failed to found a Game")
)

// Minesweepersvc interface that defines the
// service.
type Minesweepersvc interface {
	GetMinesweeper(ctx context.Context) (res MinesweeperResponse, err error)
}

// MinesweeperResponse is returned from the
// GetMinesweeper service method.
type MinesweeperResponse struct {
	Name string `json:"name,omit_empty"` // Name of the
}

// minesweeper implements the
// Minesweepersvc interface
type minesweeper struct {
	logger      log.Logger
	minesweeper MinesweeperResponse
}

// NewBasicService returns an instance of
// the Minesweepersvc.
func NewBasicService(logger log.Logger) Minesweepersvc {

	// retrieve a string
	sweeper, err := retrieveSweeper()
	if err != nil {
		logger.Log("method", "NewBasicService", "error", err)
	}

	return minesweeper{
		logger:      logger,
		minesweeper: sweeper,
	}
}

// New returns a fully initialized instance of
// the Minesweepersvc with middlewares
func New(logger log.Logger) Minesweepersvc {
	var svc Minesweepersvc
	{
		svc = NewBasicService(logger)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

// minesweeper implements Minesweepersvc
// Retrieve a Minesweeper Response (only a string for now)
func (m minesweeper) GetMinesweeper(ctx context.Context) (res MinesweeperResponse, err error) {

	return m.minesweeper, nil
}

// retrieveSweeper returns only a Response containing just a string

func retrieveSweeper() (mr MinesweeperResponse, err error) {

	var w MinesweeperResponse

	w.Name = "Todas las hojas son del viento"

	return w, nil
}
