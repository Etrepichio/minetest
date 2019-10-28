package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/minesweeper/pkg/service"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetMinesweeperEndpoint endpoint.Endpoint
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

// New will create an Endpoints struct with initialized endpoint(s) and
// middleware(s).
func New(svc service.Minesweepersvc, logger log.Logger) (ep Endpoints) {
	// create the GetMinesweeper endpoint
	ep.GetMinesweeperEndpoint = MakeGetMinesweeperEndpoint(svc)

	// create logging middleware for endpoint
	ep.GetMinesweeperEndpoint = LoggingMiddleware(log.With(logger, "method", "GetMinesweeper"))(ep.GetMinesweeperEndpoint)

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
