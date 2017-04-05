package main

import (
	"log"
	"net/http"
	"os"

	"github.com/foklepoint/fibonacci-service/service"
	ktlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := getLogger()
	svc := service.LoggingMiddleware{&service.FibonacciService{}, logger}
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getLogger() ktlog.Logger {
	logger := ktlog.NewLogfmtLogger(ktlog.NewSyncWriter(os.Stdout))
	logger = ktlog.With(logger, "component", "fibonacci-service")
	return logger
}
