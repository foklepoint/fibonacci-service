package service

import (
	"github.com/go-kit/kit/log"
	"time"
)

type LoggingMiddleware struct {
	Next   *FibonacciService
	Logger log.Logger
}

func (mw LoggingMiddleware) calculate(n uint64) (uint64, error) {
	defer func(begin time.Time) {
		mw.Logger.Log("method", "calculate", "n", n, "time", time.Since(begin))
	}(time.Now())
	return mw.Next.calculate(n)
}
