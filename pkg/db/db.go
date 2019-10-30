package db

import (
	"errors"

	"github.com/minesweeper/pkg/models"
)

type MineDBManager interface {
	InsertGame(game *models.Game) error
	UpdateGame(game *models.Game) error
	GetGame(name string) (*models.Game, error)
}

type MineStorage struct {
	data map[string]*models.Game
}

func New() *MineStorage {
	ms := MineStorage{
		data: make(map[string]*models.Game),
	}
	return &ms
}

func (ms MineStorage) InsertGame(game *models.Game) error {

	if _, ok := ms.data[game.Name]; ok {
		return errors.New("Name already used")
	}
	ms.data[game.Name] = game

	return nil
}

func (ms MineStorage) UpdateGame(game *models.Game) error {

	if _, ok := ms.data[game.Name]; !ok {
		return errors.New("Game not found")
	}
	ms.data[game.Name] = game

	return nil
}

func (ms MineStorage) GetGame(name string) (*models.Game, error) {

	if _, ok := ms.data[name]; !ok {
		return &models.Game{}, errors.New("Game not found")
	}
	resp := ms.data[name]
	return resp, nil

}
