package main

import (
	"testing"
	"time"
)

func TestProgressBar(t *testing.T) {
	var tests = []struct {
		percent int
		width   int
		want    string
	}{
		{1, 10, "[          ]"},
		{10, 10, "[=         ]"},
		{50, 10, "[=====     ]"},
		{80, 10, "[========  ]"},
		{100, 10, "[==========]"},
	}
	for _, test := range tests {
		if got := ProgressBar(test.percent, test.width); got != test.want {
			t.Errorf("ProgressBar(%d, %d) = %v, want %v", test.percent, test.width, got, test.want)
		}
	}
}

func TestTimeDisplay(t *testing.T) {
	var tests = []struct {
		currentTime time.Duration
		totalTime   time.Duration
		want        string
	}{
		{time.Second, 10 * time.Second, "1s / 10s"},
		{30 * time.Minute, 500 * time.Minute, "30m0s / 8h20m0s"},
	}
	for _, test := range tests {
		if got := TimeDisplay(test.currentTime, test.totalTime); got != test.want {
			t.Errorf("TimeDisplay(%d, %d) = %v, want %v", test.currentTime, test.totalTime, got, test.want)
		}
	}
}
