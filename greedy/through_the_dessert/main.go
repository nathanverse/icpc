package main

import (
	"fmt"
	"math"
)

func Execute() {
	var leak, consumeRate, consumption float64
	maxConsumption := float64(-1)
	var lastDistance int
	var isBreak bool
	var temp string // no used field
	for true {
		var d int
		var event string
		_, err := fmt.Scan(&d, &event)
		if err != nil {
			panic(err)
		}

		if d == 0 {
			// Reset problem
			leak = 0
			lastDistance = 0
			maxConsumption = -1
			consumption = 0
		}

		consumption += float64(d-lastDistance) * (consumeRate + leak)

		switch event {
		case "Fuel":
			{
				fmt.Scan(&temp)

				var rate float64
				fmt.Scan(&rate)

				if rate == 0 {
					isBreak = true
					break
				}

				consumeRate = rate / 100
			}
		case "Gas":
			{
				fmt.Scan(&temp)

				maxConsumption = math.Max(consumption, maxConsumption)
				consumption = 0
			}
		case "Mechanic":
			{
				leak = 0
			}
		case "Leak":
			{
				leak += 1
			}
		case "Goal":
			{
				maxConsumption = math.Max(consumption, maxConsumption)
				fmt.Printf("%.3f\n", maxConsumption)
			}
		}

		if isBreak {
			break
		}

		lastDistance = d
	}
}

func main() {
	Execute()
}
