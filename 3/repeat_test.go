package iteration

import (
	"testing"
	"fmt"
	"strings"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, expected, repeated string) {
		t.Helper()
		if expected != repeated {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeat 5 times when no repeat number is specified", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("repeat n times when n is specified", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("checking repeat with String func Compare", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"
		if strings.Compare(repeated, expected) != 0 {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("checking repeat with String func Repeat", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := strings.Repeat("a", 3)
		assertCorrectMessage(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 3)
    }
}

func ExampleRepeat() {
    repeated := Repeat("a", 3)
    fmt.Println(repeated)
    // Output: aaa
}