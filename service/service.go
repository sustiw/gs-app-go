package service

import (
	"FirstGo/models"
	"fmt"
	"sort"
	"strings"
)

func CreatePackDetails(qty int, packSizes []int) models.ResponseData {

	sortedSizes := sortSizesDescending(packSizes)
	totalCount, remaining := findingCorrectPacks(qty, packSizes)
	totalCount = handlingRemainingCounts(remaining, sortedSizes, totalCount)

	finalResponse := formatPackDetailsString(sortedSizes, totalCount)
	return models.ResponseData{
		Quantity:    qty,
		PackDetails: finalResponse,
	}

}

func sortSizesDescending(packSizes []int) []int {
	sortedPackSizes := make([]int, len(packSizes))
	copy(sortedPackSizes, packSizes)

	sort.Slice(sortedPackSizes, func(i, j int) bool {
		return sortedPackSizes[i] > sortedPackSizes[j]
	})
	return sortedPackSizes
}

func findingCorrectPacks(qty int, packSizes []int) (map[int]int, int) {
	totalCount := make(map[int]int)
	remaining := qty
	sortedSizes := sortSizesDescending(packSizes)
	for _, size := range sortedSizes {
		if remaining >= size {
			totalCount[size] = remaining / size
			remaining = remaining % size
		}
	}
	return totalCount, remaining
}

func handlingRemainingCounts(remaining int, sortedSizes []int, totalCount map[int]int) map[int]int {
	if remaining <= 0 {
		return totalCount
	}

	chosenSize := findClosestFit(remaining, sortedSizes)
	packsInChoiceA := countCurrentBoxes(totalCount) + 1

	totalTargetVolume := remaining
	for size, count := range totalCount {
		totalTargetVolume += size * count
	}
	for i := len(sortedSizes) - 1; i >= 0; i-- {
		nextLargerSize := sortedSizes[i]
		if nextLargerSize >= totalTargetVolume {
			if 1 < packsInChoiceA {
				optimizedCounts := make(map[int]int)
				optimizedCounts[nextLargerSize] = 1
				return optimizedCounts
			}
		}
	}
	totalCount[chosenSize]++
	return totalCount
}

func findClosestFit(remaining int, sortedSizes []int) int {
	chosenSize := sortedSizes[len(sortedSizes)-1]
	for i := len(sortedSizes) - 1; i >= 0; i-- {
		if sortedSizes[i] >= remaining {
			chosenSize = sortedSizes[i]
			break
		}
	}
	return chosenSize
}

func countCurrentBoxes(totalCount map[int]int) int {
	total := 0
	for _, count := range totalCount {
		total += count
	}
	return total
}

func formatPackDetailsString(sortedSizes []int, totalCount map[int]int) string {

	var details []string

	for _, size := range sortedSizes {
		count := totalCount[size]
		if count > 0 {
			itemStr := fmt.Sprintf("%dx%d", count, size)
			details = append(details, itemStr)
		}
	}

	return strings.Join(details, ",")

}
