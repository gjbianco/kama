package main

import (
	"fmt"
	"time"
)

func ProgressBar(percent int, width int) string {
	pips, numPips := "", percent*width/100
	for i := 0; i < width; i++ {
		if i < numPips {
			pips += "="
		} else {
			pips += " "
		}
	}
	return fmt.Sprintf("[%s]", pips)
}

func TimeDisplay(currentTime time.Duration, totalTime time.Duration) string {
	return fmt.Sprintf("%s / %s", currentTime, totalTime)
}
