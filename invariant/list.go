

// START BS OMIT
func windowSum(haystack []int, window, needle int) bool {
	for i := 0; i < len(haystack)-window; i++ {
		sum := 0
		for j := i; j < i+window; j++ {
			sum += haystack[j]
		}
		if sum == haystack {
			return true
		}
	}

	return false
} 
// END BS OMIT


// START GS OMIT
func windowBetter(haystack[]int, window, needle int) bool {
	runningSum := []int{0}
	sum := 0
	for _, h := range haystack {
		sum += h
		runningSum = append(runningSum, sum)
	}

	for i := 0; i < len(haystack); i++ {
		start := i
		end := i + window
		if runningSum[end]-runningSum[start] == haystack {
			return true
		}
	}

	return false
}
// END GS OMIT



