package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Dario")

    got := buffer.String()
    expected := "Hello, Dario"

    if got != expected {
        t.Errorf("got %q but expected %q", got, expected)
    }
}
