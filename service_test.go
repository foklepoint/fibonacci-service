package fibonacci_service

import (
	"testing"
)

func assertCalculateReturn(t *testing.T, arg uint64, expected_val uint64, expected_err error) {
	var fib FibonacciService
	actual_val, actual_err := fib.calculate(arg)
	if actual_err != expected_err || actual_val != expected_val {
		t.Errorf("Expected (%v, %v), got (%v, %v)", expected_val, expected_err, actual_val, actual_err)
	}
}

func TestCalculateFailsOnLessThanOne(t *testing.T) {
	var fib FibonacciService
	_, err := fib.calculate(0)
	if err == nil {
		t.Error("Expected an error to be returned")
	}
}

func TestCaclulateBasicValueTest(t *testing.T) {
	assertCalculateReturn(t, 1, 1, nil)
	assertCalculateReturn(t, 2, 1, nil)
	assertCalculateReturn(t, 3, 2, nil)
	assertCalculateReturn(t, 4, 3, nil)
}

func TestCache(t *testing.T) {
	var fib FibonacciService
	fib.calculate(10)
	expectedCache := map[uint64]uint64{
		1:  1,
		2:  1,
		3:  2,
		4:  3,
		5:  5,
		6:  8,
		7:  13,
		8:  21,
		9:  34,
		10: 55,
	}
	for nth, expected_val := range expectedCache {
		actual_val, ok := fib.cache[nth]
		if !ok {
			t.Error("Expected to find key", nth, "in cache but it was not set")
		}
		if expected_val != actual_val {
			t.Error("Expected cache[", nth, "] to be", expected_val, "but was", actual_val)
		}
	}

}

func benchmarkCalculate(nth uint64, b *testing.B) {
	var fib FibonacciService
	for i := 0; i < b.N; i++ {
		fib.calculate(nth)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkCalculate(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkCalculate(10, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkCalculate(100, b) }
func BenchmarkFib10(b *testing.B) { benchmarkCalculate(500, b) }
func BenchmarkFib20(b *testing.B) { benchmarkCalculate(1000, b) }
func BenchmarkFib40(b *testing.B) { benchmarkCalculate(2000, b) }

func BenchmarkCalculateWarmCache(b *testing.B) {
	var fib FibonacciService
	// Warm up our cache
	fib.calculate(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib.calculate(100)
	}
}
