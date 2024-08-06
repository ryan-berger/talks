package main

import "strconv"

// ashishbhoi
func IsValidISBN(isbn string) bool {
	// START OMIT
	var isbnNumbers []int
	for i, char := range isbn {
		digit, err := strconv.Atoi(string(char))
		if err == nil {
			isbnNumbers = append([]int{digit}, isbnNumbers...)
		} else if char == 'X' {
			if i == len(isbn)-1 {
				isbnNumbers = append([]int{10}, isbnNumbers...)
			} else {
				return false
			}
		} else if char != '-' {
			return false
		}
	}

	var sum int
	for i, num := range isbnNumbers {
		sum += num * (i + 1)
	}

	return sum%11 == 0 && len(isbnNumbers) == 10
	// END OMIT
}

func main() {}

