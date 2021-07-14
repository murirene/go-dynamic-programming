package pattern3_test

import (
	"github.com/murirene/go-dynamic-programming/pattern3"
	"testing"
)

func TestFib(t *testing.T) {
	if pattern3.Fib(5) != 5 {
		t.Fatal("Failed 3-1")
	}

	if pattern3.Fib(6) != 8 {
		t.Fatal("Failed 3-1")
	}

	if pattern3.Fib(7) != 13 {
		t.Fatal("Failed 3-1")
	}
}
