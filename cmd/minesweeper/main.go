package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
	"github.com/minesweeper/pkg/endpoints"
	httpmine "github.com/minesweeper/pkg/http"
	"github.com/minesweeper/pkg/service"
	"github.com/oklog/run"
)

func main() {
	var (
		addr = flag.String("addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	// Prepare logging
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
		logger = log.With(logger, "svc", "minesweeper")
	}

	// Build service layers from inside out
	var handler http.Handler
	{
		svc := service.New(logger)
		eps := endpoints.New(svc, logger)
		handler = httpmine.NewHTTPHandler(eps, logger)
	}

	var g run.Group
	{
		// Set up service's http listener
		httpListener, err := net.Listen("tcp", *addr)
		bailOnError(logger, err)

		g.Add(func() error {
			logger.Log("transport", "http", "addr", *addr)
			return http.Serve(httpListener, handler)
		}, func(err error) {
			logger.Log("transport", "http", "err", err)
			httpListener.Close()
		})
	}
	{
		// Set-up our signal handler.
		var (
			cancelInterrupt = make(chan struct{})
			c               = make(chan os.Signal, 2)
		)
		defer close(c)

		g.Add(func() error {
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}

func bailOnError(logger log.Logger, err error) {
	if err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}
}
