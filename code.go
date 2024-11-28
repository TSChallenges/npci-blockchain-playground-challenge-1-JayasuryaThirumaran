package main

import (
	"errors"
	"fmt"
	"sort"
)

// FilterAndSort filters numbers greater than or equal to the threshold and sorts them.
func FilterAndSort(nums []int, threshold int) []int {
	// Filter the numbers greater than or equal to the threshold
	var filtered []int
	for _, num := range nums {
		if num >= threshold {
			filtered = append(filtered, num)
		}
	}
	
	// Sort the filtered slice in ascending order
	sort.Ints(filtered)
	return filtered
}

// FindMostFrequent finds the most frequent word in a slice of strings.
// Returns an error if the slice is empty.
func FindMostFrequent(words []string) (string, error) {
	if len(words) == 0 {
		return "", errors.New("slice is empty")
	}

	// Create a map to count the frequency of each word
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// Find the word with the maximum frequency
	var mostFrequent string
	maxCount := 0
	for word, count := range wordCount {
		if count > maxCount {
			mostFrequent = word
			maxCount = count
		}
	}

	return mostFrequent, nil
}

func main() {
	// Test FilterAndSort
	fmt.Println("Testing FilterAndSort:")
	nums := []int{3, 10, 1, 7, 8, 2}
	threshold := 5
	fmt.Printf("Input: %v, Threshold: %d, Output: %v\n", nums, threshold, FilterAndSort(nums, threshold))

	// Test FindMostFrequent
	fmt.Println("\nTesting FindMostFrequent:")
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	result, err := FindMostFrequent(words)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Input: %v, Most Frequent: %s\n", words, result)
	}

	// Edge case: empty slice
	emptyWords := []string{}
	result, err = FindMostFrequent(emptyWords)
	if err != nil {
		fmt.Printf("Error: %v (empty input case handled)\n", err)
	} else {
		fmt.Printf("Input: %v, Most Frequent: %s\n", emptyWords, result)
	}
}
