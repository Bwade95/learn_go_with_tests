package iteration

import (
	"fmt"
	"testing"
)

func TestRepeatConcatenation(t *testing.T) {
	repeated := RepeatConcatenation("a", 7)
	expected := "aaaaaaa"

	AssertEqual(t, expected, repeated)
}

func TestRepeatBuilder(t *testing.T) {
	repeated := RepeatBuilder("a", 10)
	expected := "aaaaaaaaaa"

	AssertEqual(t, expected, repeated)
}

func BenchmarkRepeatConcatenation(b *testing.B) {
	for b.Loop() {
		RepeatConcatenation("a", 10000)
	}
}

func BenchmarkRepeatBuilder(b *testing.B) {
	for b.Loop() {
		RepeatBuilder("a", 10000)
	}
}

func AssertEqual(t testing.TB, expected, repeated string) {
	t.Helper()
	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeatConcatenation() {
	repeated := RepeatConcatenation("a", 7)
	fmt.Println(repeated)
	// Output: aaaaaaa
}

func ExampleRepeatBuilder() {
	repeated := RepeatBuilder("a", 7)
	fmt.Println(repeated)
	// Output: aaaaaaa
}
