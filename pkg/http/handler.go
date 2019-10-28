package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/minesweeper/pkg/endpoints"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoints.Endpoints, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerFinalizer(RequestLogFinalizer(logger)),
	}

	c := chi.NewRouter()
	c.Method(http.MethodGet, "/minesweeper", httptransport.NewServer(
		endpoints.GetMinesweeperEndpoint,
		DecodeGetMinesweeperRequest,
		EncodeGetMinesweeperResponse,
		append(options)...,
	))
	return c
}

// RequestLogFinalizer is called at the end of an http request. Use it to log final
// information regarding a request.
func RequestLogFinalizer(logger log.Logger) httptransport.ServerFinalizerFunc {
	return func(ctx context.Context, code int, r *http.Request) {
		// log a bunch of response values
		level.Info(logger).Log("code", code,
			"ua", r.UserAgent(),
			"remote", r.RemoteAddr,
			"method", r.Method,
			"url", r.URL,
			"path", r.URL.Path,
			"host", r.Host)
	}
}

// DecodeGetMinesweeperRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeGetMinesweeperRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	// GET request takes no parameters so just
	// return empty GetMinesweeperRequest
	return endpoints.GetMinesweeperRequest{}, nil
}

// EncodeGetMinesweeperResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeGetMinesweeperResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	// set response header
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// cast response to known type
	res := response.(endpoints.GetMinesweeperResponse)

	// create json
	err = json.NewEncoder(w).Encode(res.Res)
	return err
}
