package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(5, 2)
	expected := 3

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}