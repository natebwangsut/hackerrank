/*
 * Nate B. Wangsutthitham
 * <@natebwangsut | nate.bwangsut@gmail.com>
 */

package main

import "fmt"

func birthdayCakeCandles() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n, num, count, height int
	fmt.Scanf("%v", &n)
	height = -1
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &num)
		if num > height {
			height = num
			count = 1
		} else if num == height {
			count++
		}
	}
	fmt.Println(count)
}
