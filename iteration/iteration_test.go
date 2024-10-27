package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("Repeat counter: 6", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeat counter: 0", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeat counter: -1", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

func assertCorrectMessage(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
