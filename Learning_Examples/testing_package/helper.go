package main

import (
	"testing"
)

func DoPanic() {
	panic("This is a panic")
}

func TestDoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but did not get one")
		}
	}()
	DoPanic()
}
