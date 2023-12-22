package main

// This version has bugs

import (
	"testing"
	"unicode/utf8"
)

func TestR1(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{" ", " "},
		{"!12345@", "@54321!"},
		{"Mastering Go", "oG gniretsaM"},
	}
	for _, tc := range testcases {
		rev := R1(tc.in)
		if string(rev) != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func TestR2(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{" ", " "},
		{"!12345@", "@54321!"},
		{"Mastering Go", "oG gniretsaM"},
	}
	for _, tc := range testcases {
		rev := R1(tc.in)
		if string(rev) != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzR1(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := R1(orig)
		doubleRev := R1(string(rev))
		if orig != string(doubleRev) {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(string(rev)) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

func FuzzR2(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev := R2(orig)
		doubleRev := R2(string(rev))
		if orig != string(doubleRev) {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
