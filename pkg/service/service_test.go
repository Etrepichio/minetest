package service

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/minesweeper/pkg/db"

	"github.com/minesweeper/pkg/models"
	uuid "github.com/nu7hatch/gouuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewGame(t *testing.T) {
	u, _ := uuid.NewV4()

	newb := models.Game{
		Name:       u.String(),
		Columns:    5,
		Rows:       5,
		Mines:      5,
		Discovered: 0,
		Status:     "start",
	}

	exisb := models.Game{
		Name:       "already",
		Columns:    5,
		Rows:       5,
		Mines:      5,
		Discovered: 0,
		Status:     "start",
	}

	cases := []struct {
		name    string
		board   models.Game
		isValid func(err error) bool
	}{
		{
			name:  "New Game",
			board: newb,
			isValid: func(err error) bool {
				return err == nil
			},
		},
		{
			name:  "Existing Game",
			board: exisb,
			isValid: func(err error) bool {
				return err != nil
			},
		},
	}
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = level.NewFilter(logger, level.AllowAll())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
		logger = log.With(logger, "svc", "Minesweeper Test")
	}
	db := db.New()

	db.InsertGame(&exisb)
	service := minesweeper{
		logger: logger,
		db:     db,
	}
	Convey("Test NewGame", t, func() {
		for _, c := range cases {
			Convey(c.name, func() {
				So(c.isValid(service.NewGame(context.TODO(), &c.board)), ShouldBeTrue)
			})
		}
	})

}

func TestClick(t *testing.T) {

	game := models.Game{
		Name:       "already",
		Columns:    5,
		Rows:       5,
		Mines:      5,
		Discovered: 0,
		Status:     "start",
	}

	u, _ := uuid.NewV4()

	cases := []struct {
		name    string
		click   models.ClickRequest
		isValid func(game *models.Game, err error) bool
	}{
		{
			name: "Correct Click",
			click: models.ClickRequest{
				Column: 2,
				Row:    2,
				Name:   "already",
			},
			isValid: func(game *models.Game, err error) bool {
				return err == nil
			},
		},
		{
			name: "Click out of bounds",
			click: models.ClickRequest{
				Column: 7,
				Row:    8,
				Name:   "already",
			},
			isValid: func(game *models.Game, err error) bool {
				return err != nil
			},
		},
		{
			name: "Inexisting Game",
			click: models.ClickRequest{
				Column: 2,
				Row:    2,
				Name:   u.String(),
			},
			isValid: func(game *models.Game, err error) bool {
				return err != nil
			},
		},
	}
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = level.NewFilter(logger, level.AllowAll())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
		logger = log.With(logger, "svc", "Minesweeper Test")
	}
	db := db.New()

	service := minesweeper{
		logger: logger,
		db:     db,
	}

	service.NewGame(context.TODO(), &game)

	Convey("Test Click", t, func() {
		for _, c := range cases {
			Convey(c.name, func() {
				So(c.isValid(service.Click(context.TODO(), c.click)), ShouldBeTrue)
			})
		}
	})

}
