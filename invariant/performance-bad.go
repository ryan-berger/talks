package main

import "fmt"

// START OMIT
func firstDifference(x, y []int) int {
	for i := 0; i < min(len(y), len(x)); i++ {
		if x[i] != y[i] {
			return i
		}
	}
	return len(x)
}
// END OMIT

func main() {
	example := make([]int, 20)
	fmt.Println(firstDifference(example, example))
}
