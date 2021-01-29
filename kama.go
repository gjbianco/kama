package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"gopkg.in/ini.v1"
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
	isBreak     = flag.Bool("b", false, "forces a break timer")
	isLongBreak = flag.Bool("l", false, "forces a long break timer")
	isWork      = flag.Bool("w", true, "forces a work timer")

	breakSeconds     = 5 * time.Minute
	longBreakSeconds = 15 * time.Minute
	workSeconds      = 25 * time.Minute
)

func main() {
	flag.Parse()

	userHome, _ := os.UserHomeDir()
	cfg, err := ini.Load(userHome + "/.kamarc")

	// retrieve time duration configs
	if err == nil {
		breakSeconds = time.Duration(cfg.Section("").Key("break_time").MustInt()) * time.Second
		longBreakSeconds = time.Duration(cfg.Section("").Key("long_break_time").MustInt()) * time.Second
		workSeconds = time.Duration(cfg.Section("").Key("work_time").MustInt()) * time.Second

		*size = cfg.Section("").Key("width").MustInt(15)
	}

	verb := Work
	switch {
	case *isBreak:
		verb = Break
	case *isLongBreak:
		verb = LongBreak
	case *isWork:
		verb = Work
	}

	if *totalTime == time.Duration(0) {
		switch verb {
		case Break:
			*totalTime = time.Duration(breakSeconds)
		case LongBreak:
			*totalTime = time.Duration(longBreakSeconds)
		case Work:
			*totalTime = time.Duration(workSeconds)
		}
	}

	fmt.Printf("%s\n", strings.ToUpper(string(verb)))
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
		hour, minute, _ := time.Now().Clock()
		fmt.Printf("finished at %d:%d\n", hour%12, minute)
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
