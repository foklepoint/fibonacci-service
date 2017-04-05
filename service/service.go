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
func (service *FibonacciService) calculate(n uint64) (ret uint64, err error) {
	if !(n >= 1) {
		err = fmt.Errorf("n (%d) >= 1", n)
		return
	}
	if service.cache == nil {
		service.cache = make(map[uint64]uint64)
	}
	if n == 1 || n == 2 {
		ret = 1
	} else if val, ok := service.cache[n]; ok {
		ret = val
	} else {
		n_2, _ := service.calculate(n - 2)
		n_1, _ := service.calculate(n - 1)
		ret = n_1 + n_2
	}
	service.cache[n] = ret
	return
}
