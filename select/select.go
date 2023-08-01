package select_concurrency

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// We only care about which ends first
	// select allows to wait on multiple channels
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		// time.After also returns a chan
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// We return chan struct{} since we only want to notify we closed the channel and struct{} is
// the smallest data type available since it doesn't allocate any data
func ping(url string) chan struct{} {
	// Always use make to create channels, otherwise if we use var we'll get the initial value as nil
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
