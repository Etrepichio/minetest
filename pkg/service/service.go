package service

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/minesweeper/pkg/db"
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
	Click(ctx context.Context, req models.ClickRequest) (res *models.Game, err error)
}

// MinesweeperResponse is returned from the
// GetMinesweeper service method.
type MinesweeperResponse struct {
	Name string `json:"name,omit_empty"`
}

// minesweeper implements the Minesweepersvc interface
// it also contains a logger and a db to store games.
type minesweeper struct {
	logger      log.Logger
	minesweeper MinesweeperResponse
	db          db.MineDBManager
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
		db:          db.New(),
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

//NewGame sets default values in the new game. Also creates the board and populates it with numbers and mines.
//Once the game is created, its status is "new", and is inserted in the db for future use
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

	if err := newBoard(game); err != nil {
		return err
	}
	game.Status = "new"
	//Here we should save the game in order to load it in the future
	if err := m.db.InsertGame(game); err != nil {
		return err
	}
	return nil

}

//LoadGame grabs and return a game by its name from the db
func (m minesweeper) LoadGame(ctx context.Context, name string) (res *models.Game, err error) {

	if name == "" {
		return nil, errors.New(models.ErrNoNameGame)
	}

	//Here we should load the game
	game, err := m.db.GetGame(name)
	if err != nil {
		return &models.Game{}, err
	}
	return game, nil

}

//SaveGame update the current status of the game into the db. Used when the user makes moves (Click)
func (m minesweeper) SaveGame(ctx context.Context, game *models.Game) (err error) {

	if err := m.db.UpdateGame(game); err != nil {
		return err
	}

	return nil
}

//Click updates the game according to the cell the user has clicked.
func (m minesweeper) Click(ctx context.Context, req models.ClickRequest) (res *models.Game, err error) {

	//First load the game from the db
	game, err := m.LoadGame(ctx, req.Name)

	if err != nil {
		return &models.Game{}, err
	}
	//click the specific cell
	if err := clickCell(game, req.Row, req.Column); err != nil {

		return &models.Game{}, err
	}
	//update the game to its new state
	if err := m.SaveGame(ctx, game); err != nil {
		return &models.Game{}, err
	}

	return game, nil
}

func clickCell(game *models.Game, row int, column int) error {
	//Check that row and column arent out of bounds
	if row > game.Rows-1 || row < 0 {
		return errors.New("invalid row")
	}
	if column > game.Columns-1 || column < 0 {
		return errors.New("invalid column")
	}
	//Checking if Cell was already clicked
	if game.Board[row][column].Clicked == true {
		return errors.New("Already clicked")
	}
	//Check for game loss
	if game.Board[row][column].Mine == true {
		game.Status = "game_over"
		return nil
	}
	game.Board[row][column].Clicked = true
	//Discovered tracks how many cells have been clicked (used for win condition)
	game.Discovered++
	//If the Cell has 0 mines in its proximity, we click all around it, recursively
	//This means we do the clickCell() function 9 times inside itself (all surrounding cells)
	//Complexity of this would be O(9^n), but the conditions of out of bounds and already clicked
	//mitigate the impact this would have in performance.

	if game.Board[row][column].Number == 0 {
		for x := row - 1; x < row+2; x++ {
			for y := column - 1; y < column+2; y++ {
				clickCell(game, x, y)
			}
		}
	}

	//Check for game win
	if game.Discovered+game.Mines == game.Rows*game.Columns {
		game.Status = "victory"
		return nil
	}
	return nil

}

//newBoard populates the board with mines and numbers
func newBoard(game *models.Game) error {

	//First we populate the board with mines
	numCells := game.Rows * game.Columns
	cells := make(models.CellRow, numCells)
	i := 0
	for i < game.Mines {
		//Get random spot for mines
		rand.Seed(time.Now().UnixNano())
		spot := rand.Intn(numCells)
		if cells[spot].Mine == false {
			cells[spot].Mine = true
			i++
		}
	}

	game.Board = make([]models.CellRow, game.Rows)
	//Fit in the final board each of the cell rows
	for c := range game.Board {
		game.Board[c] = cells[c*game.Columns : ((c + 1) * game.Columns)]
	}
	//O(n^2)
	for i, row := range game.Board {
		for j, cell := range row {
			if cell.Mine == true {
				setNumbers(game, i, j)
			}
		}

	}
	return nil
}

//setNumbers pluses by 1 the Number value in each cell surrounding a mine
//when every mine has his neighbours's numbers setted, board is ready
func setNumbers(game *models.Game, i int, j int) {

	for x := i - 1; x < i+2; x++ {
		for y := j - 1; y < j+2; y++ {
			if !((x < 0) || (x > game.Rows-1) || (y < 0) || (y > game.Columns-1)) && !((x == i) && (y == j)) {
				game.Board[x][y].Number++
			}
		}
	}

}
