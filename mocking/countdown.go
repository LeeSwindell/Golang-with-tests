package main

import (
	"fmt"
	"io"
	"os"
	"time"
)


type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type Sleeper interface {
	Sleep()
}

type SpySleeperOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

const sleep = "sleep"
const write = "write"

func (s *SpySleeperOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpySleeperOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return 
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(w, "GO!")
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}