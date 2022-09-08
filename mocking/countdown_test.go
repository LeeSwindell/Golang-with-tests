package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)


func TestCountdown(t *testing.T) {
	
	t.Run("prints correct countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpySleeperOperations{})

		got := buffer.String()
		want := `3
2
1
GO!`

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("sleeps between counting down", func( t *testing.T) {
		spySleeper := &SpySleeperOperations{}
		Countdown(spySleeper, spySleeper)

		got := spySleeper.Calls
		want := []string{write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(got, want) {
			t. Errorf("call order incorrect, got %v, wanted %v", got, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := SpyTime{}
	sleeper := ConfigurableSleeper{duration: sleepTime, sleep: spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("got %v sleep time, wanted %v sleep time", spyTime.durationSlept, sleepTime)
	}
}