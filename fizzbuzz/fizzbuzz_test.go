package fizzbuzz_test

import (
	"hello/fizzbuzz"
	"testing"
)

func TestGivenOne(t *testing.T) {
	given := 1
	want := "1"

	get := fizzbuzz.Count(given)
	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}
}

func TestGivenTwo(t *testing.T) {
	given := 2
	want := "2"

	get := fizzbuzz.Count(given)
	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}
}

func TestGivenThree(t *testing.T) {
	given := 3
	want := "Fizz"

	get := fizzbuzz.Count(given)
	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}
}

func TestGivenFour(t *testing.T) {
	given := 4
	want := "4"

	get := fizzbuzz.Count(given)
	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}
}

func TestGivenFive(t *testing.T) {
	given := 5
	want := "Buzz"

	get := fizzbuzz.Count(given)
	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}
}
