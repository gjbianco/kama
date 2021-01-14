package main

import (
	"flag"
	"fmt"
	"github.com/gen2brain/beeep"
	"time"
)

const (
	timeTick = 1 * int(time.Second)
)

var length = flag.Float64("l", 25, "timer length in minutes")
var quiet = flag.Bool("q", false, "disable the alert notification on timer completion")
var message = flag.String("m", "timer finished!", "message to display in notification (no effect if silent)")
var width = flag.Int("w", 15, "width of the progress indicator")

func main() {
	flag.Parse()
	millis := int(*length * float64(time.Minute))
	timer(millis, *quiet, *message, *width)
}

func timer(length int, quiet bool, message string, width int) {
	fmt.Printf("starting timer for %.2f minutes\n", float64(length)/float64(time.Minute))
	for i := timeTick; i < length; i += timeTick {
		fmt.Printf("%s %s\r", ProgressBar(100*(i+timeTick)/length, width), TimeDisplay(i/int(time.Second), length/int(time.Second)))
		// fmt.Printf("%s\r", ProgressBar(100*i/length, width))
		time.Sleep(time.Duration(timeTick))
	}
	fmt.Println()
	if !quiet {
		// TODO add image icon?
		beeep.Notify("kama", message, "")
	}
}
