package main

import (
	"fmt"
	"math"
)

func calculateRisk(baseRisk, fragments int) int {
	return int(math.Pow(float64(baseRisk), float64(fragments)))
}

func isAchievable(dataCenters []int, fragments, maxRisk int) bool {
	fragmentsLeft := fragments

	for _, risk := range dataCenters {
		usedFragments := 0
		for usedFragments < fragmentsLeft {
			if calculateRisk(risk, usedFragments+1) > maxRisk {
				break
			}
			usedFragments++
		}
		fragmentsLeft -= usedFragments
		if fragmentsLeft <= 0 {
			return true
		}
	}

	return false
}

func distributeFragments(dataCenters []int, fragments int) int {
	left, right := 0, int(1e9) // Start with a high upper limit for binary search
	result := right

	for left <= right {
		mid := (left + right) / 2
		if isAchievable(dataCenters, fragments, mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func main() {
	dataCenters := []int{10, 20, 30}
	fragments := 5
	minimizedMaxRisk := distributeFragments(dataCenters, fragments)
	fmt.Printf("Minimized maximum risk: %d\n", minimizedMaxRisk)
}
