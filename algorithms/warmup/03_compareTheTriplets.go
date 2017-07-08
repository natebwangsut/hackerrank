/*
 * Nate B. Wangsutthitham
 * <@natebwangsut | nate.bwangsut@gmail.com>
 */

package main

import "fmt"

func compareTheTriplets() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var a0, a1, a2, b0, b1, b2 int
	fmt.Scanf("%v %v %v", &a0, &a1, &a2)
	fmt.Scanf("%v %v %v", &b0, &b1, &b2)

	aScore := 0
	bScore := 0

	if a0 > b0 {
		aScore++
	} else if a0 < b0 {
		bScore++
	}

	if a1 > b1 {
		aScore++
	} else if a1 < b1 {
		bScore++
	}

	if a2 > b2 {
		aScore++
	} else if a2 < b2 {
		bScore++
	}

	fmt.Printf("%v %v", aScore, bScore)
}
