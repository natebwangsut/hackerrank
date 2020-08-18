/*
 * Nate B. Wangsutthitham
 * <@natebwangsut | nate.bwangsut@gmail.com>
 */

package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	n := 0
	fmt.Scanf("%v", &n)
	//fmt.Printf("%v\n", n)

	var sum, num int64
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &num)
		sum += num
		//fmt.Printf("%v: ", num)
	}
	fmt.Printf("%v", sum)
}
