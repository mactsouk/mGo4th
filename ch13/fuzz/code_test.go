package main

import (
	"testing"
)

func TestAddInt(t *testing.T) {
	testCases := []struct {
		x, y, want int
	}{
		{1, 2, 3},
		{1, 0, 1},
		{100, 10, 110},
	}

	for _, tc := range testCases {
		result := AddInt(tc.x, tc.y)
		if result != tc.want {
			t.Errorf("X: %d, Y: %d, want %d", tc.x, tc.y, tc.want)
		}
	}
}

func FuzzAddInt(f *testing.F) {
	testCases := []struct {
		x, y int
	}{
		{0, 1},
		{0, 100},
	}

	for _, tc := range testCases {
		f.Add(tc.x, tc.y)
	}

	f.Fuzz(func(t *testing.T, x, y int) {
		result := AddInt(x, y)

		if result != x+y {
			t.Errorf("X: %d, Y: %d, Result %d, want %d", x, y, result, x+y)
		}
	})
}
