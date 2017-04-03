package fibonacci_service

import (
	"testing"
	"context"
)

func TestMakeCalculateEndpointCreatesEndpoint(t *testing.T) {
	endpoint := makeCalculateEndpoint(fibonacciService{})
	ctx := context.Background()
	req := CalculateRequest{5}
	resp, err := endpoint(ctx, req)
	res := resp.(CalculateResponse)
	var expected_answer_val uint64 = 5
	if err != nil || res.err != "" {
		t.Errorf("Expected err, resp err to be nil, '' but was %v, %s", err, res.err)
	}
	if res.val != expected_answer_val {
		t.Error("Expected response val to be", expected_answer_val, "but was", res.val)
	}
}

func TestMakeCalculateEndpointReturnsError(t *testing.T) {
	endpoint := makeCalculateEndpoint(fibonacciService{})
	ctx := context.Background()
	req := CalculateRequest{0}
	resp, err := endpoint(ctx, req)
	res := resp.(CalculateResponse)
	expected_err_resp := "nth (0) >= 1"
	if err != nil {
		t.Errorf("Expected err, resp err to be nil, '' but was %v, %s", err, res.err)
	}
	if res.err != expected_err_resp {
		t.Errorf("expected error to be '%s' but was '%s'", expected_err_resp, res.err)
	}
}
