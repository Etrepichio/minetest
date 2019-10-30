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
//This allow us to wrap our service's final functions with layers of logging, decoding, encoding, etc, abstracting the core functionality of the service
//from the rest
func New(svc service.Minesweepersvc, logger log.Logger) (ep Endpoints) {
	// create the GetMinesweeper endpoint

	ep.GetMinesweeperEndpoint = MakeGetMinesweeperEndpoint(svc)
	ep.GetMinesweeperEndpoint = LoggingMiddleware(log.With(logger, "method", "GetMinesweeper"))(ep.GetMinesweeperEndpoint)

	//create the NewGame endpoint
	ep.NewGameEndpoint = MakeNewGameEndpoint(svc)
	ep.NewGameEndpoint = LoggingMiddleware(log.With(logger, "method", "NewGame"))(ep.NewGameEndpoint)

	//create the LoadGame endpoint
	ep.LoadGameEndpoint = MakeLoadGameEndpoint(svc)
	ep.LoadGameEndpoint = LoggingMiddleware(log.With(logger, "method", "LoadGame"))(ep.LoadGameEndpoint)

	//create the SaveGame endpoint
	ep.SaveGameEndpoint = MakeSaveGameEndpoint(svc)
	ep.SaveGameEndpoint = LoggingMiddleware(log.With(logger, "method", "SaveGame"))(ep.SaveGameEndpoint)

	//create the Click endpoint
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

// MakeNewGameEndpoint returns an endpoint that invokes NewGame on the service.
func MakeNewGameEndpoint(svc service.Minesweepersvc) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(NewGameRequest)
		err := svc.NewGame(ctx, req.Req)

		// wrap service response with endpoint response
		return NewGameResponse{Err: err}, nil
	}
}

// MakeLoadGameEndpoint returns an endpoint that invokes LoadGame on the service.
func MakeLoadGameEndpoint(svc service.Minesweepersvc) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoadGameRequest)
		res, err := svc.LoadGame(ctx, req.Req)

		// wrap service response with endpoint response
		return LoadGameResponse{Res: res, Err: err}, nil
	}
}

// MakeSaveGameEndpoint returns an endpoint that invokes SaveGame on the service.
func MakeSaveGameEndpoint(svc service.Minesweepersvc) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SaveGameRequest)
		err := svc.SaveGame(ctx, req.Req)

		// wrap service response with endpoint response
		return SaveGameResponse{Err: err}, nil
	}
}

// MakeClickEndpoint returns an endpoint that invokes Click on the service.
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

// NewGameRequest contains the data required for the endpoint
type NewGameRequest struct {
	Req *models.Game
}

//NewGameResponse tracks errors from the service
type NewGameResponse struct {
	Err error
}

// LoadGameRequest is an empty request object
// because no parameters are required to make this request
type LoadGameRequest struct {
	Req string
}

// LoadGameResponse is the endpoint response
// that wraps the service response object and tracks
// errors from Minesweeper Service.
type LoadGameResponse struct {
	Res *models.Game
	Err error
}

// SaveGameRequest contains the data required for the endpoint
type SaveGameRequest struct {
	Req *models.Game
}

// SaveGameResponse tracks errors from the service
type SaveGameResponse struct {
	Err error
}

// ClickRequest contains the data required for the endpoint
type ClickRequest struct {
	Req models.ClickRequest
}

// ClickResponse is the endpoint response
// that wraps the service response object and tracks
// errors from Minesweeper Service
type ClickResponse struct {
	Res *models.Game
	Err error
}
