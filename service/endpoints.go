package service

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

func MakeCalculateEndpoint(service IFibonacciService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CalculateRequest)
		val, err := service.calculate(req.N)
		var errStr string
		if err != nil {
			errStr = err.Error()
		}
		return CalculateResponse{val, errStr}, nil
	}
}

// Holds params for Fibonacci.calculate as JSON
type CalculateRequest struct {
	N uint64 `json:"nth"`
}

// Holds the return of Fibonacci.calculate for a JSON response
type CalculateResponse struct {
	Val uint64 `json:"val"`
	Err string `json:"error"`
}

func DecodeCalculateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
