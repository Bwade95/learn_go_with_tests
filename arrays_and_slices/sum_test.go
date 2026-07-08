package arraysandslices

import (
	"testing"
)

func TestSumArray(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		expected := 6
		actual := Sum(numbers)

		AssertEqual(t, expected, actual, numbers)
	})
}

func AssertEqual(t testing.TB, expected, actual, numbers any) {
	if actual != expected {
		t.Errorf("expected %d but got %d, %v", expected, actual, numbers)
	}
}
