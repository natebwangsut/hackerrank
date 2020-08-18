/*
 * Nate B. Wangsutthitham
 * <@natebwangsut | nate.bwangsut@gmail.com>
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var num [5]int64

	var sum, min, max int64

	min = math.MaxInt64
	max = math.MinInt64

	for i := 0; i < 5; i++ {
		fmt.Scanf("%v", &num[i])
	}

	for i := 0; i < 5; i++ {

		sum = 0
		for j := 0; j < 5; j++ {
			if j != i {
				sum += num[j]
			}
		}

		if sum < min {
			min = sum
		}

		if sum > max {
			max = sum
		}
	}

	fmt.Printf("%v %v", min, max)
}
