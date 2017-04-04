// Defines the fibonacci microservice
package service

import "fmt"

// Fibonacci service interface
type IFibonacciService interface {
	calculate(num uint64) (uint64, error)
}

// Fibonacci service implements the interface
// Keeps a cache of calculated values
type FibonacciService struct {
	cache map[uint64]uint64
}

// Calculates nth fibonacci number, nth >= 1
// Memoizes past calculated values in service cache
func (service *FibonacciService) calculate(nth uint64) (ret uint64, err error) {
	if !(nth >= 1) {
		err = fmt.Errorf("nth (%d) >= 1", nth)
		return
	}
	if service.cache == nil {
		service.cache = make(map[uint64]uint64)
	}
	if nth == 1 || nth == 2 {
		ret = 1
	} else if val, ok := service.cache[nth]; ok {
		ret = val
	} else {
		n_2, _ := service.calculate(nth - 2)
		n_1, _ := service.calculate(nth - 1)
		ret = n_1 + n_2
	}
	service.cache[nth] = ret
	return
}
