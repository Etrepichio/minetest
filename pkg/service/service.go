package service

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/minesweeper/pkg/models"
)

const (
	defaultRows    = 8
	defaultColumns = 8
	defaultMines   = 14
	maxRows        = 36
	maxColumns     = 36
)

// Minesweepersvc interface that defines the
// service.
type Minesweepersvc interface {
	GetMinesweeper(ctx context.Context) (res MinesweeperResponse, err error)
	NewGame(ctx context.Context, game *models.Game) (err error)
	LoadGame(ctx context.Context, name string) (res *models.Game, err error)
	SaveGame(ctx context.Context, game *models.Game) (err error)
	Click(ctx context.Context, game *models.Game, rowClick int, columnClick int) (res *models.Game, err error)
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

func (m minesweeper) NewGame(ctx context.Context, game *models.Game) (err error) {

	if game.Name == "" {
		return errors.New(models.ErrNoNameGame)
	}
	if game.Rows == 0 {
		game.Rows = defaultRows
	}
	if game.Columns == 0 {
		game.Columns = defaultColumns
	}
	if game.Mines == 0 {
		game.Mines = defaultMines
	}
	if game.Rows > maxRows {
		game.Rows = maxRows
	}
	if game.Columns > maxColumns {
		game.Columns = maxColumns
	}
	game.Status = "new"

	//Here we should save the game in order to load it in the future

	return nil

}

func (m minesweeper) LoadGame(ctx context.Context, name string) (res *models.Game, err error) {

	if name == "" {
		return nil, errors.New(models.ErrNoNameGame)
	}

	//Here we should load the game

	return nil, nil

}

func (m minesweeper) SaveGame(ctx context.Context, game *models.Game) (err error) {

	//Here we save the game

	return nil
}

func (m minesweeper) Click(ctx context.Context, game *models.Game, rowClick int, columnClick int) (res *models.Game, err error) {

	//Here we add the logic, calculation of game over or victory, etc.

	return nil, nil
}
