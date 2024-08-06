
package main


// START BB OMIT
func binarySearch(haystack []int, needle int) int {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		mid := (low + high) / 2
		elem := haystack[mid]
		switch  {
		case elem < needle:
			low = mid + 1
		case elem > needle:
			high = mid - 1
		default:
			return mid
		}
	}
	return -(low + 1)
}
// END BB OMIT


// START BG OMIT
func binarySearchGood(haystack []int, needle int) int {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		mid := int((uint(low) + uint(high)) / 2)
		elem := haystack[mid]
		switch  {
		case elem < needle:
			low = mid + 1
		case elem > needle:
			high = mid - 1
		default:
			return mid
		}
	}
	return -(low + 1)
}
// END BG OMIT

func main() {}

