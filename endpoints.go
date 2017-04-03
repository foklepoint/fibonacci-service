package fibonacci_service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


func makeCalculateEndpoint(service fibonacciService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CalculateRequest)
		val, err := service.calculate(req.nth)
		var errStr string
		if err != nil {
			errStr = err.Error()
		}
		return CalculateResponse{val, errStr}, nil
	}
}