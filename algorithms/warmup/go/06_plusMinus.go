/*
 * Nate B. Wangsutthitham
 * <@natebwangsut | nate.bwangsut@gmail.com>
 */

package main

import (
	"fmt"
)

func plusMinus() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	var n, cur int
	var pos, neg, zero float64

	fmt.Scanf("%v", &n)
	//fmt.Printf("%v", n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &cur)

		if cur < 0 {
			neg++
		}

		if cur == 0 {
			zero++
		}

		if cur > 0 {
			pos++
		}
	}

	fmt.Println(pos / float64(n))
	fmt.Println(neg / float64(n))
	fmt.Println(zero / float64(n))
}
