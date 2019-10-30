package db

import (
	"errors"

	"github.com/minesweeper/pkg/models"
)

//MineDBManager is the interface that express the operations needed to manage the storage of the minesweeper. Its isolated from its implementation
type MineDBManager interface {
	InsertGame(game *models.Game) error
	UpdateGame(game *models.Game) error
	GetGame(name string) (*models.Game, error)
}

//MineStorage implements MineDBManager. For now it just contains a map of strings and games.
//In a future version, it could have a real db client and implement the interface around that
type MineStorage struct {
	data map[string]*models.Game
}

//New creates a new MineStorage and instantiates the data parameter of it
func New() *MineStorage {
	ms := MineStorage{
		data: make(map[string]*models.Game),
	}
	return &ms
}

//InsertGame checks that there's no other game stored with the same name as the new game. Afterwards, its saved in the map
func (ms MineStorage) InsertGame(game *models.Game) error {

	if _, ok := ms.data[game.Name]; ok {
		return errors.New("Name already used")
	}
	ms.data[game.Name] = game

	return nil
}

//UpdateGame ensures a game with the update exists already and after that updates it
func (ms MineStorage) UpdateGame(game *models.Game) error {

	if _, ok := ms.data[game.Name]; !ok {
		return errors.New("Game not found")
	}
	ms.data[game.Name] = game

	return nil
}

//GetGame obtains a game from the map according to its name
func (ms MineStorage) GetGame(name string) (*models.Game, error) {

	if _, ok := ms.data[name]; !ok {
		return &models.Game{}, errors.New("Game not found")
	}
	resp := ms.data[name]
	return resp, nil

}
