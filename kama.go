package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

const (
	// how frequently the display should be drawn
	timeTick = time.Second
)

var length = flag.String("l", "25m", "timer length -- default: 25m)'")
var quiet = flag.Bool("q", false, "disable the alert notification on timer completion")
var message = flag.String("m", "timer finished!", "message to display in notification (no effect if silent)")
var width = flag.Int("w", 15, "width of the progress indicator -- default: 15")

var startTime = time.Now()

func main() {
	flag.Parse()
	totalTime, err := time.ParseDuration(*length)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printDisplay(totalTime)

	ticker := time.NewTicker(timeTick)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				printDisplay(totalTime)
			}
		}
	}()

	time.Sleep(totalTime)
	ticker.Stop()
	done <- true

	fmt.Println()
	if !*quiet {
		// TODO add image icon?
		beeep.Notify("kama", *message, "")
	}
}

func printDisplay(totalTime time.Duration) {
	// TODO timer state should be a struct
	currentTime := time.Since(startTime).Truncate(time.Second)
	percent := int(100 * currentTime / totalTime)
	display := ""
	display += ProgressBar(percent, *width)
	display += fmt.Sprintf(" %d%% : ", percent)
	display += TimeDisplay(currentTime, totalTime)
	fmt.Printf("%s\r", display)
}
