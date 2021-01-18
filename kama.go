package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

type Verb string

const (
	Break     Verb = "break"
	LongBreak Verb = "longbreak"
	Work      Verb = "work"
)

const (
	// how frequently the display should be drawn
	timeTick = time.Second
)

var (
	startTime = time.Now()

	// CLI flags
	totalTime = flag.Duration("t", time.Duration(0), "timer length")
	quiet     = flag.Bool("q", false, "disable notification")
	message   = flag.String("m", "timer finished!", "notification message")
	size      = flag.Int("s", 15, "width of the progress indicator")

	// flags to force verb/mode
	isWork      = flag.Bool("w", true, "forces a work timer")
	isBreak     = flag.Bool("b", false, "forces a break timer")
	isLongBreak = flag.Bool("l", false, "forces a long break timer")
)

func main() {
	flag.Parse()

	var verb Verb
	switch {
	case *isBreak:
		verb = Break
	case *isLongBreak:
		verb = LongBreak
	case *isWork:
		fallthrough
	default:
		verb = Work
	}

	if *totalTime == time.Duration(0) {
		switch verb {
		case Break:
			*totalTime = 5 * time.Minute
		case LongBreak:
			*totalTime = 15 * time.Minute
		case Work:
			*totalTime = 25 * time.Minute
		}
	}

	fmt.Printf("[[%s]]\n", verb)
	printDisplay(*totalTime)

	ticker := time.NewTicker(timeTick)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				printDisplay(*totalTime)
			}
		}
	}()

	time.Sleep(*totalTime)
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
	display += ProgressBar(percent, *size)
	display += fmt.Sprintf(" %d%% : ", percent)
	display += TimeDisplay(currentTime, totalTime)
	fmt.Printf("%s\r", display)
}
