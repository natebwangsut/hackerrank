/*
 * Nate B. Wangsutthitham
 * <@natebwangsut | nate.bwangsut@gmail.com>
 */

package main

import (
	"fmt"
	"math"
)

func diagonalDifference() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var primary, secondary, cur, n int
	fmt.Scanf("%v", &n)
	//fmt.Println(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%v", &cur)

			if i == j {
				primary += cur
				//fmt.Printf("primary: %v %v %v\n", i , j, cur)
			}

			if j+i == n-1 {
				secondary += cur
				//fmt.Printf("secondary: %v %v %v\n", i , j, cur)
			}

		}
	}
	fmt.Println(math.Abs(float64(primary - secondary)))
}
