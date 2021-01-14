package main

import (
	"fmt"
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

func TimeDisplay(currentTime int, totalTime int) string {
	return fmt.Sprintf("%ds / %ds", currentTime, totalTime)
}
