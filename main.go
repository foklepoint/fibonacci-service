package main

import (
	"net/http"
	"log"
	"./service"

	httptransport "github.com/go-kit/kit/transport/http"
)


func main()  {
	svc := service.FibonacciService{}
	http.Handle(
		"/calculate",
		httptransport.NewServer(
			service.MakeCalculateEndpoint(svc),
			service.DecodeCalculateRequest,
			service.EncodeResponse,
		),
	)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
