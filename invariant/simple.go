package main

import (
	"strconv"
	"strings"
)

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


// START IB OMIT
func IsValidISBNBetter(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}

	list := strings.Split(isbn, "")
	if list[9] == "X" {
		list[9] = "10"
	}

	sum := 0
	for i, c := range list {
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			return false
		}
		sum += digit * (len(list)-i)
	}

	return sum % 11 == 0
}
// END IB OMIT


func main() {}

