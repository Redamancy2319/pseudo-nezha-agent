package pseudostat

import (
	"math/rand"
	"strconv"
	"strings"
)

func ProcessValue(input string) float64 {
	var randValue float64
	switch {
	case strings.Contains(input, "-"):
		randValue = processRange(input)
	case strings.Contains(input, ","):
		randValue = processDistribution(input)
	default:
		randValue = processValue(input)
	}

	return randValue
}

func processValue(input string) float64 {
	inputValue, err := strconv.ParseFloat(input, 64)
	var randomValue float64

	switch {
	case err != nil:
		println("Failed to parse integer: ", err)
	case inputValue < 0 || inputValue > 100:
		println("Invalid input:", inputValue)
	default:
		randomValue = inputValue
	}

	return randomValue
}

func processRange(input string) float64 {
	inputRange := strings.Split(input, "-")
	var randomValue float64
	if len(inputRange) == 2 {
		lower, err1 := strconv.ParseFloat(inputRange[0], 64)
		upper, err2 := strconv.ParseFloat(inputRange[1], 64)

		switch {
		case err1 != nil || err2 != nil:
			println("Failed to parse: ", err1, err2)
		case lower < 0 || upper > 100 || upper < lower:
			println("Input values out of range:", lower, upper)
		default:
			randomValue = rand.Float64()*(upper-lower) + lower
		}
	} else {
		println("Invalid input:", inputRange)
	}

	return randomValue
}

func processDistribution(input string) float64 {
	inputNd := strings.Split(input, ",")
	var randomValue float64
	if len(inputNd) == 2 {
		expectedValue, err1 := strconv.ParseFloat(inputNd[0], 64)
		stdValue, err2 := strconv.ParseFloat(inputNd[1], 64)

		switch {
		case err1 != nil || err2 != nil:
			println("Failed to parse: ", err1, err2)
		case expectedValue < 0 || expectedValue > 100 || stdValue < 0:
			println("Input values out of range:", expectedValue, stdValue)
		default:
			randomValue = rand.NormFloat64()*stdValue + expectedValue
			switch {
			case randomValue < 0:
				randomValue = 0
			case randomValue > 100:
				randomValue = 100
			}
		}
	} else {
		println("Invalid input:", inputNd)
	}

	return randomValue
}
