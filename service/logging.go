package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func RequestLoggingMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			logger.Log("ctx", ctx)
			return next(ctx, request)
		}
	}
}
