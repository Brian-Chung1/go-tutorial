package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want int, numbers [5]int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	}

	t.Run("sums numbers in array", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assertCorrectMessage(t, got, want, numbers)
	})

}