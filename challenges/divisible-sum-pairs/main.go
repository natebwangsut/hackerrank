package main

// Complete the divisibleSumPairs function below.
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	count := int32(0)
	for i := 0; int32(i) < n; i++ {
		for j := i + 1; int32(j) < n; j++ {
			if int32(ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return count
}