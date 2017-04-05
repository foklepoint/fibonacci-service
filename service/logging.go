package service

import (
	"fmt"
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
	fmt.Println(len(mw.Next.cache))
	return mw.Next.calculate(n)
}
