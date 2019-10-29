package endpoints

import (
	"context"

	"github.com/minesweeper/pkg/models"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/minesweeper/pkg/service"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetMinesweeperEndpoint endpoint.Endpoint
	NewGameEndpoint        endpoint.Endpoint
	LoadGameEndpoint       endpoint.Endpoint
	SaveGameEndpoint       endpoint.Endpoint
	ClickEndpoint          endpoint.Endpoint
}

// New will create an Endpoints struct with initialized endpoint(s) and
// middleware(s).
func New(svc service.Minesweepersvc, logger log.Logger) (ep Endpoints) {
	// create the GetMinesweeper endpoint

	ep.GetMinesweeperEndpoint = MakeGetMinesweeperEndpoint(svc)
	ep.GetMinesweeperEndpoint = LoggingMiddleware(log.With(logger, "method", "GetMinesweeper"))(ep.GetMinesweeperEndpoint)

	ep.NewGameEndpoint = MakeNewGameEndpoint(svc)
	ep.NewGameEndpoint = LoggingMiddleware(log.With(logger, "method", "NewGame"))(ep.NewGameEndpoint)

	ep.LoadGameEndpoint = MakeLoadGameEndpoint(svc)
	ep.LoadGameEndpoint = LoggingMiddleware(log.With(logger, "method", "LoadGame"))(ep.LoadGameEndpoint)

	ep.SaveGameEndpoint = MakeSaveGameEndpoint(svc)
	ep.SaveGameEndpoint = LoggingMiddleware(log.With(logger, "method", "SaveGame"))(ep.SaveGameEndpoint)

	ep.ClickEndpoint = MakeClickEndpoint(svc)
	ep.ClickEndpoint = LoggingMiddleware(log.With(logger, "method", "Click"))(ep.ClickEndpoint)
	return ep
}

// MakeGetMinesweeperEndpoint returns an endpoint that invokes GetMinesweeper on the service.
func MakeGetMinesweeperEndpoint(svc service.Minesweepersvc) (ep endpoint.Endpoint) {

	// interface parameter is ignored because request does not
	// require input parameters.
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		// invoke eysvcService method
		res, err := svc.GetMinesweeper(ctx)

		// wrap service response with endpoint response
		return GetMinesweeperResponse{Res: res, Err: err}, nil
	}
}

func MakeNewGameEndpoint(svc service.Minesweepersvc) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(NewGameRequest)
		err := svc.NewGame(ctx, req.Req)

		// wrap service response with endpoint response
		return NewGameResponse{Err: err}, nil
	}
}

func MakeLoadGameEndpoint(svc service.Minesweepersvc) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoadGameRequest)
		res, err := svc.LoadGame(ctx, req.Req)

		// wrap service response with endpoint response
		return LoadGameResponse{Res: res, Err: err}, nil
	}
}

func MakeSaveGameEndpoint(svc service.Minesweepersvc) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SaveGameRequest)
		err := svc.SaveGame(ctx, req.Req)

		// wrap service response with endpoint response
		return SaveGameResponse{Err: err}, nil
	}
}

func MakeClickEndpoint(svc service.Minesweepersvc) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ClickRequest)
		res, err := svc.Click(ctx, req.Req)

		// wrap service response with endpoint response
		return ClickResponse{Res: res, Err: err}, nil
	}
}

// GetMinesweeperRequest is an empty request object
// because no parameters are required to make this request
type GetMinesweeperRequest struct{}

// GetMinesweeperResponse is the endpoint response
// that wraps the service response object and tracks
// errors from Minesweeper Service.
type GetMinesweeperResponse struct {
	Res service.MinesweeperResponse // Res is the Minesweepersvc Response
	Err error                       // Err is an error encountered in Minesweepersvc
}

type NewGameRequest struct {
	Req *models.Game
}

type NewGameResponse struct {
	Err error
}

type LoadGameRequest struct {
	Req string
}

type LoadGameResponse struct {
	Res *models.Game
	Err error
}

type SaveGameRequest struct {
	Req *models.Game
}

type SaveGameResponse struct {
	Err error
}

type ClickRequest struct {
	Req models.ClickRequest
}

type ClickResponse struct {
	Res *models.Game
	Err error
}
