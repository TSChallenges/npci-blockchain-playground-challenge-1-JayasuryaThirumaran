# GoLang: Functions Practice


## Overview

This code involves implementing two functions in Go to practice working with slices, sorting, maps, and error handling. You'll also test your implementations using a variety of scenarios.

---

## Tasks

### 1. **FilterAndSort Function**

Write a function `FilterAndSort(nums []int, threshold int) []int` that:
- Filters out numbers less than the given `threshold` from the input slice `nums`.
- Sorts the remaining numbers in ascending order.
- Returns the filtered and sorted slice.

#### Example
```go
nums := []int{3, 10, 1, 7, 8, 2}
threshold := 5
result := FilterAndSort(nums, threshold)
// Output: [7, 8, 10]
```

### 2. **FindMostFrequent Function**

Write a function `FindMostFrequent(words []string) (string, error)` that:
- Identifies and returns the most frequent word in the input slice of strings `words`.
- If there is a tie, it can return any one of the most frequent words.
- Returns an error if the input slice is empty.

#### Example
```go
words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
result, err := FindMostFrequent(words)
// Output: "apple"

words = []string{}
result, err = FindMostFrequent(words)
// Output: error ("slice is empty")
```

### **Instructions**

- Implement the functions in a file named `code.go`.
- Write a `main` function to test the functions with a variety of inputs, including edge cases.
- Run the program to verify the output matches the expected results.

