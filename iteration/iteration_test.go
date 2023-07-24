package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	expected := "aaaaaa"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func ExampleRepeat() {
    repeated := Repeat("t", 5)
    fmt.Println(repeated)
    // Output: ttttt
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
