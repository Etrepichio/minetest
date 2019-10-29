package http

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/minesweeper/pkg/endpoints"
	"github.com/minesweeper/pkg/models"
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
	c.Method(http.MethodPost, "/minesweeper/games", httptransport.NewServer(
		endpoints.NewGameEndpoint,
		DecodeNewGameRequest,
		EncodeNewGameResponse,
		append(options)...,
	))
	c.Method(http.MethodGet, "/minesweeper/games/{name}", httptransport.NewServer(
		endpoints.LoadGameEndpoint,
		DecodeLoadGameRequest,
		EncodeLoadGameResponse,
		append(options)...,
	))
	c.Method(http.MethodPut, "/minesweeper/games", httptransport.NewServer(
		endpoints.ClickEndpoint,
		DecodeClickRequest,
		EncodeClickResponse,
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

func DecodeNewGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req models.Game
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		if err == io.EOF {
			return nil, errors.New("Missing Body Content")
		} else if err == io.ErrUnexpectedEOF {
			return nil, errors.New("Malformed Body Content")
		} else {
			return nil, err
		}
	}

	return endpoints.NewGameRequest{
		Req: &req,
	}, nil
}

func EncodeNewGameResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {

	// cast response to known type
	res, ok := response.(endpoints.NewGameResponse)
	if !ok {
		return errors.New("Error encoding NewGame response")
	}

	if res.Err != nil {
		return res.Err
	}

	w.WriteHeader(http.StatusCreated)
	return err
}

func DecodeLoadGameRequest(_ context.Context, r *http.Request) (interface{}, error) {

	name := chi.URLParam(r, "name")
	return endpoints.LoadGameRequest{
		Req: name,
	}, nil
}

func EncodeLoadGameResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {

	// cast response to known type
	res, ok := response.(endpoints.LoadGameResponse)
	if !ok {
		return errors.New("Error encoding LoadGame response")
	}

	if res.Err != nil {
		return res.Err
	}

	return json.NewEncoder(w).Encode(res.Res)
}

func DecodeClickRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req models.ClickRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		if err == io.EOF {
			return nil, errors.New("Missing Body Content")
		} else if err == io.ErrUnexpectedEOF {
			return nil, errors.New("Malformed Body Content")
		} else {
			return nil, err
		}
	}
	return endpoints.ClickRequest{
		Req: req,
	}, nil
}

func EncodeClickResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {

	// cast response to known type
	res, ok := response.(endpoints.ClickResponse)
	if !ok {
		return errors.New("Error encoding LoadGame response")
	}

	if res.Err != nil {
		return res.Err
	}

	return json.NewEncoder(w).Encode(res.Res)
}
