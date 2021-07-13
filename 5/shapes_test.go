package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangle perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0
		assertCorrectMessage(t, got, want)
	})

}

func TestArea(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("get area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0
		assertCorrectMessage(t, got, want)
	})

	t.Run("get area of circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}
