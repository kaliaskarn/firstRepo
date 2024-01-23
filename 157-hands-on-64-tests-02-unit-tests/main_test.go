package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("sum was incoreect, got: %d want %d", total, 10)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(10, 5)
	if total != 5 {
		t.Errorf("sum was incoreect, got: %d want %d", total, 5)
	}
}

func TestDoMathAdd(t *testing.T) {
	result := doMath(5, 3, add)
	expected := 8
	if result != expected {
		t.Errorf("doMath with add was incorrect, got: %d, want: %d", result, expected)
	}
}
