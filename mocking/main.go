package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

type Sleeper interface {
	Sleep()
}

// More generic sleeper with arbitrary long countdowns
type ConfigurableSleeper struct {
	sleep    func(time.Duration)
	duration time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Default sleeper simply sleeps for 1 second
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{time.Sleep, 1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
