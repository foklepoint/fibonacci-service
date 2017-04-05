package main

import (
	"log"
	"net/http"
	"os"
	"net"

	ktlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/foklepoint/fibonacci-service/service"
)

func main() {
	logger := getLogger()
	svc := service.LoggingMiddleware{
		&service.FibonacciService{},
		ktlog.With(logger, "type", "endpoint"),
	}
	endpoint := service.MakeCalculateEndpoint(svc)
	service.MakeCalculateEndpoint(svc)
	var server http.Handler = *httptransport.NewServer(
		endpoint,
		service.DecodeCalculateRequest,
		service.EncodeResponse,
	)
	http.Handle(
		"/calculate",
		server,
	)
	listener, err := net.Listen("tcp", ":8080")
	if err == nil {
		logger.Log("status", "ready")
	}
	log.Fatal(http.Serve(listener, nil))

}

func getLogger() ktlog.Logger {
	logger := ktlog.NewLogfmtLogger(ktlog.NewSyncWriter(os.Stdout))
	logger = ktlog.With(logger, "component", "fibonacci-service")
	return logger
}
